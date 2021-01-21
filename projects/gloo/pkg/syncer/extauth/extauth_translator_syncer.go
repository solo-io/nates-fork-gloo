package extauth

import (
	"context"
	"fmt"

	"github.com/solo-io/solo-projects/projects/extauth/pkg/runner"
	"go.opencensus.io/stats"
	"go.opencensus.io/stats/view"
	"go.opencensus.io/tag"

	"github.com/golang/protobuf/proto"
	"github.com/mitchellh/hashstructure"
	gloov1 "github.com/solo-io/gloo/projects/gloo/pkg/api/v1"
	extauth "github.com/solo-io/gloo/projects/gloo/pkg/api/v1/enterprise/options/extauth/v1"
	"github.com/solo-io/gloo/projects/gloo/pkg/syncer"
	glooutils "github.com/solo-io/gloo/projects/gloo/pkg/utils"
	"github.com/solo-io/go-utils/contextutils"
	"github.com/solo-io/go-utils/hashutils"
	envoycache "github.com/solo-io/solo-kit/pkg/api/v1/control-plane/cache"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources"
	"github.com/solo-io/solo-kit/pkg/api/v2/reporter"
	extAuthPlugin "github.com/solo-io/solo-projects/projects/gloo/pkg/plugins/extauth"
	"go.uber.org/zap"
)

// Compile-time assertion
var (
	_ syncer.TranslatorSyncerExtension            = new(TranslatorSyncerExtension)
	_ syncer.UpgradeableTranslatorSyncerExtension = new(TranslatorSyncerExtension)
)

var (
	extauthConnectedStateDescription = "zero indicates gloo detected an error with the auth config and did not update its XDS snapshot, check the gloo logs for errors"
	extauthConnectedState            = stats.Int64("glooe.extauth/connected_state", extauthConnectedStateDescription, "1")

	extauthConnectedStateView = &view.View{
		Name:        "glooe.extauth/connected_state",
		Measure:     extauthConnectedState,
		Description: extauthConnectedStateDescription,
		Aggregation: view.LastValue(),
		TagKeys:     []tag.Key{},
	}
)

const (
	Name = "extauth"
)

func init() {
	_ = view.Register(extauthConnectedStateView)
}

type TranslatorSyncerExtension struct {
}

func NewTranslatorSyncerExtension(params syncer.TranslatorSyncerExtensionParams) *TranslatorSyncerExtension {
	return &TranslatorSyncerExtension{}
}

func (s *TranslatorSyncerExtension) Sync(
	ctx context.Context,
	snap *gloov1.ApiSnapshot,
	xdsCache envoycache.SnapshotCache,
	reports reporter.ResourceReports,
) (string, error) {
	ctx = contextutils.WithLogger(ctx, "extAuthTranslatorSyncer")
	logger := contextutils.LoggerFrom(ctx)
	snapHash := hashutils.MustHash(snap)
	logger.Infof("begin auth sync %v (%v proxies, %v upstreams, %v endpoints, %v secrets, %v artifacts, %v auth configs)", snapHash,
		len(snap.Proxies), len(snap.Upstreams), len(snap.Endpoints), len(snap.Secrets), len(snap.Artifacts), len(snap.AuthConfigs))
	defer logger.Infof("end auth sync %v", snapHash)

	return runner.ExtAuthServerRole, s.SyncAndSet(ctx, snap, xdsCache, reports)
}

func (s *TranslatorSyncerExtension) ExtensionName() string {
	return Name
}

func (s *TranslatorSyncerExtension) IsUpgrade() bool {
	return true
}

type SnapshotSetter interface {
	SetSnapshot(node string, snapshot envoycache.Snapshot) error
}

func (s *TranslatorSyncerExtension) SyncAndSet(
	ctx context.Context,
	snap *gloov1.ApiSnapshot,
	xdsCache SnapshotSetter,
	reports reporter.ResourceReports,
) error {
	helper := newHelper()
	reports.Accept(snap.AuthConfigs.AsInputResources()...)
	reports.Accept(snap.Proxies.AsInputResources()...)

	for _, proxy := range snap.Proxies {
		for _, listener := range proxy.Listeners {
			httpListener, ok := listener.ListenerType.(*gloov1.Listener_HttpListener)
			if !ok {
				// not an http listener - skip it as currently ext auth is only supported for http
				continue
			}

			virtualHosts := httpListener.HttpListener.VirtualHosts

			for _, virtualHost := range virtualHosts {
				virtualHost = proto.Clone(virtualHost).(*gloov1.VirtualHost)
				virtualHost.Name = glooutils.SanitizeForEnvoy(ctx, virtualHost.Name, "virtual host")

				if err := helper.processAuthExtension(ctx, snap, virtualHost.GetOptions().GetExtauth(), reports, proxy); err != nil {
					// Continue to next virtualHost, error has been added to the report.
					continue
				}

				for _, route := range virtualHost.Routes {
					if err := helper.processAuthExtension(ctx, snap, route.GetOptions().GetExtauth(), reports, proxy); err != nil {
						// Continue to next route, error has been added to the report.
						continue
					}

					for _, weightedDestination := range route.GetRouteAction().GetMulti().GetDestinations() {
						if err := helper.processAuthExtension(ctx, snap, weightedDestination.GetOptions().GetExtauth(),
							reports, proxy); err != nil {
							// Continue to next weighted destination, error has been added to the report.
							continue
						}
					}
				}
			}
		}
	}

	var resources []envoycache.Resource
	for _, cfg := range helper.translatedConfigs {
		resource := extauth.NewExtAuthConfigXdsResourceWrapper(cfg)
		resources = append(resources, resource)
	}
	h, err := hashstructure.Hash(resources, nil)
	if err != nil {
		contextutils.LoggerFrom(ctx).With(zap.Error(err)).DPanic("error hashing ext auth")
		return syncerError(ctx, err)
	}
	extAuthSnapshot := envoycache.NewEasyGenericSnapshot(fmt.Sprintf("%d", h), resources)
	err = xdsCache.SetSnapshot(runner.ExtAuthServerRole, extAuthSnapshot)
	if err != nil {
		return syncerError(ctx, err)
	}

	stats.Record(ctx, extauthConnectedState.M(int64(1)))

	return nil
}

func syncerError(ctx context.Context, err error) error {
	stats.Record(ctx, extauthConnectedState.M(int64(0)))
	return err
}

// This translation helper contains a map where each key is the unique identifier of an AuthConfig and the corresponding
// value is the translated config. We use it avoid translating the same configuration multiple times.
type helper struct {
	translatedConfigs map[string]*extauth.ExtAuthConfig
}

func newHelper() *helper {
	return &helper{
		translatedConfigs: make(map[string]*extauth.ExtAuthConfig),
	}
}

func (h *helper) processAuthExtension(ctx context.Context, snap *gloov1.ApiSnapshot, config *extauth.ExtAuthExtension,
	reports reporter.ResourceReports, parentProxy resources.InputResource) error {
	configRef := config.GetConfigRef()
	if configRef == nil {
		// Just return if there is nothing to translate
		return nil
	}

	// Don't perform duplicate work if we already have translated this resource
	if _, ok := h.translatedConfigs[configRef.Key()]; ok {
		return nil
	}

	cfg, err := snap.AuthConfigs.Find(configRef.GetNamespace(), configRef.GetName())
	if err != nil {
		contextutils.LoggerFrom(ctx).Warnf("Unable to find referenced auth config %v in snapshot", configRef)
		reports.AddError(parentProxy, err)
		return err

	}

	// do validation
	extAuthPlugin.ValidateAuthConfig(cfg, reports)

	translatedConfig, err := extAuthPlugin.TranslateExtAuthConfig(ctx, snap, configRef)
	if err != nil {
		reports.AddError(cfg, err)
		return err
	} else if translatedConfig == nil {
		// Do nothing if config is empty or consists only of custom auth configs
		return nil
	}

	h.translatedConfigs[configRef.Key()] = translatedConfig
	return nil
}
