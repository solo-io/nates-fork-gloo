package services

import (
	"net"
	"time"

	"github.com/solo-io/solo-projects/projects/gloo/pkg/setup"

	gatewaysyncer "github.com/solo-io/gloo/projects/gateway/pkg/syncer"

	"context"
	"sync/atomic"

	"github.com/solo-io/go-utils/contextutils"
	"github.com/solo-io/solo-kit/pkg/api/external/kubernetes/service"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients/factory"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients/kube/cache"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients/memory"

	gatewayv1 "github.com/solo-io/gloo/projects/gateway/pkg/api/v1"
	gloov1 "github.com/solo-io/gloo/projects/gloo/pkg/api/v1"
	v1 "github.com/solo-io/gloo/projects/gloo/pkg/api/v1"
	"github.com/solo-io/gloo/projects/gloo/pkg/bootstrap"
	skkube "github.com/solo-io/solo-kit/pkg/api/v1/resources/common/kubernetes"
	"google.golang.org/grpc"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	"go.uber.org/zap"

	. "github.com/onsi/gomega"
	fds_syncer "github.com/solo-io/gloo/projects/discovery/pkg/fds/syncer"
	uds_syncer "github.com/solo-io/gloo/projects/discovery/pkg/uds/syncer"
	"github.com/solo-io/gloo/projects/gloo/pkg/defaults"
	"github.com/solo-io/gloo/projects/gloo/pkg/syncer"

	"k8s.io/client-go/kubernetes"
)

type TestClients struct {
	GatewayClient        gatewayv1.GatewayClient
	VirtualServiceClient gatewayv1.VirtualServiceClient
	ProxyClient          gloov1.ProxyClient
	UpstreamClient       gloov1.UpstreamClient
	SecretClient         gloov1.SecretClient
	ServiceClient        skkube.ServiceClient
	GlooPort             int
}

var glooPort int32 = 8100

func RunGateway(ctx context.Context, justgloo bool) TestClients {
	return RunGatewayWithNamespaceAndKubeClient(ctx, justgloo, defaults.GlooSystem, nil)
}

func RunGatewayWithNamespaceAndKubeClient(ctx context.Context, justgloo bool, ns string, kubeclient kubernetes.Interface) TestClients {
	return RunGatewayWithKubeClientAndSettings(ctx, justgloo, ns, kubeclient, nil)
}

func RunGatewayWithSettings(ctx context.Context, justgloo bool, extensions *v1.Extensions) TestClients {
	return RunGatewayWithKubeClientAndSettings(ctx, justgloo, defaults.GlooSystem, nil, extensions)
}

func RunGatewayWithKubeClientAndSettings(ctx context.Context, justgloo bool, ns string, kubeclient kubernetes.Interface, extensions *v1.Extensions) TestClients {
	cache := memory.NewInMemoryResourceCache()

	testclients := GetTestClients(cache)
	testclients.GlooPort = RunGlooGatewayUdsFds(ctx, cache, What{DisableGateway: justgloo}, ns, kubeclient, extensions)
	return testclients
}

func GetTestClients(cache memory.InMemoryResourceCache) TestClients {

	// construct our own resources:
	factory := &factory.MemoryResourceClientFactory{
		Cache: cache,
	}

	gatewayClient, err := gatewayv1.NewGatewayClient(factory)
	Expect(err).NotTo(HaveOccurred())
	virtualServiceClient, err := gatewayv1.NewVirtualServiceClient(factory)
	Expect(err).NotTo(HaveOccurred())
	upstreamClient, err := gloov1.NewUpstreamClient(factory)
	Expect(err).NotTo(HaveOccurred())
	secretClient, err := gloov1.NewSecretClient(factory)
	Expect(err).NotTo(HaveOccurred())
	proxyClient, err := gloov1.NewProxyClient(factory)
	Expect(err).NotTo(HaveOccurred())

	return TestClients{
		GatewayClient:        gatewayClient,
		VirtualServiceClient: virtualServiceClient,
		UpstreamClient:       upstreamClient,
		SecretClient:         secretClient,
		ProxyClient:          proxyClient,
	}
}

type What struct {
	DisableGateway bool
	DisableUds     bool
	DisableFds     bool
}

func RunGlooGatewayUdsFds(ctx context.Context, cache memory.InMemoryResourceCache, what What, ns string, kubeclient kubernetes.Interface, extensions *v1.Extensions) int {
	port := AllocateGlooPort()
	RunGlooGatewayUdsFdsOnPort(ctx, cache, port, what, ns, kubeclient, extensions)
	return int(port)
}

func AllocateGlooPort() int32 {
	return atomic.AddInt32(&glooPort, 1)

}

func RunGlooGatewayUdsFdsOnPort(ctx context.Context, cache memory.InMemoryResourceCache, localglooPort int32, what What, ns string, kubeclient kubernetes.Interface, extensions *v1.Extensions) {

	glooopts := DefaultGlooOpts(ctx, cache, ns, kubeclient)
	glooopts.BindAddr.(*net.TCPAddr).Port = int(localglooPort)
	// no gateway for now
	if !what.DisableGateway {
		opts := DefaultTestConstructOpts(ctx, cache, ns)
		go gatewaysyncer.RunGateway(opts)
	}
	settings := v1.Settings{
		Extensions: extensions,
	}
	glooopts.Settings = &settings
	glooopts.ControlPlane.StartGrpcServer = true
	go syncer.RunGlooWithExtensions(glooopts, setup.GetGlooEeExtensions())
	if !what.DisableFds {
		go fds_syncer.RunFDS(glooopts)
	}
	if !what.DisableUds {
		go uds_syncer.RunUDS(glooopts)
	}

}

func DefaultTestConstructOpts(ctx context.Context, cache memory.InMemoryResourceCache, ns string) gatewaysyncer.Opts {
	ctx = contextutils.WithLogger(ctx, "gateway")
	ctx = contextutils.SilenceLogger(ctx)
	f := &factory.MemoryResourceClientFactory{
		Cache: cache,
	}

	return gatewaysyncer.Opts{
		WriteNamespace:  ns,
		WatchNamespaces: []string{"default", ns},
		Gateways:        f,
		VirtualServices: f,
		Proxies:         f,
		WatchOpts: clients.WatchOpts{
			Ctx:         ctx,
			RefreshRate: time.Minute,
		},
		DevMode: false,
	}
}

func DefaultGlooOpts(ctx context.Context, cache memory.InMemoryResourceCache, ns string, kubeclient kubernetes.Interface) bootstrap.Opts {
	ctx = contextutils.WithLogger(ctx, "gloo")
	logger := contextutils.LoggerFrom(ctx)
	grpcServer := grpc.NewServer(grpc.StreamInterceptor(
		grpc_middleware.ChainStreamServer(
			grpc_ctxtags.StreamServerInterceptor(),
			grpc_zap.StreamServerInterceptor(zap.NewNop()),
			func(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
				logger.Infof("gRPC call: %v", info.FullMethod)
				return handler(srv, ss)
			},
		)),
	)
	f := &factory.MemoryResourceClientFactory{
		Cache: cache,
	}
	return bootstrap.Opts{
		WriteNamespace: ns,
		Upstreams:      f,
		UpstreamGroups: f,
		Proxies:        f,
		Secrets:        f,
		Services:       newServiceClient(ctx, f, kubeclient),

		Artifacts:       f,
		WatchNamespaces: []string{"default", ns},
		WatchOpts: clients.WatchOpts{
			Ctx:         ctx,
			RefreshRate: time.Second / 10,
		},
		ControlPlane: syncer.NewControlPlane(ctx, grpcServer, nil, true),
		BindAddr: &net.TCPAddr{
			IP:   net.ParseIP("0.0.0.0"),
			Port: 8081,
		},
		KubeClient: kubeclient,
		DevMode:    true,
	}
}

func newServiceClient(ctx context.Context, memFactory *factory.MemoryResourceClientFactory, kubeclient kubernetes.Interface) skkube.ServiceClient {

	// If the KubeClient option is set, the kubernetes discovery plugin will be activated and we must provide a
	// kubernetes service client in order for service-derived upstreams to be included in the snapshot
	if kubeclient != nil {
		kubeCache, err := cache.NewKubeCoreCache(ctx, kubeclient)
		if err != nil {
			panic(err)
		}
		return service.NewServiceClient(kubeclient, kubeCache)
	}

	// Else return in-memory client
	client, err := skkube.NewServiceClient(memFactory)
	if err != nil {
		panic(err)
	}
	return client
}
