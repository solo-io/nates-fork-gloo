// Code generated by solo-kit. DO NOT EDIT.

package v1

import (
	"context"
	"fmt"
	"time"

	"github.com/hashicorp/go-multierror"
	"go.opencensus.io/stats"
	"go.opencensus.io/stats/view"
	"go.opencensus.io/tag"
	"go.opencensus.io/trace"

	"github.com/solo-io/go-utils/contextutils"
	"github.com/solo-io/go-utils/errutils"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients"
	"github.com/solo-io/solo-kit/pkg/api/v1/eventloop"
	"github.com/solo-io/solo-kit/pkg/errors"
	skstats "github.com/solo-io/solo-kit/pkg/stats"
)

var (
	mEdsSnapshotTimeSec     = stats.Float64("eds.gloo.solo.io/sync/time_sec", "The time taken for a given sync", "1")
	mEdsSnapshotTimeSecView = &view.View{
		Name:        "eds.gloo.solo.io/sync/time_sec",
		Description: "The time taken for a given sync",
		TagKeys:     []tag.Key{tag.MustNewKey("syncer_name")},
		Measure:     mEdsSnapshotTimeSec,
		Aggregation: view.Distribution(0.01, 0.05, 0.1, 0.25, 0.5, 1, 5, 10, 60),
	}
)

func init() {
	view.Register(
		mEdsSnapshotTimeSecView,
	)
}

type EdsSyncer interface {
	Sync(context.Context, *EdsSnapshot) error
}

type EdsSyncers []EdsSyncer

func (s EdsSyncers) Sync(ctx context.Context, snapshot *EdsSnapshot) error {
	var multiErr *multierror.Error
	for _, syncer := range s {
		if err := syncer.Sync(ctx, snapshot); err != nil {
			multiErr = multierror.Append(multiErr, err)
		}
	}
	return multiErr.ErrorOrNil()
}

type edsEventLoop struct {
	emitter EdsSnapshotEmitter
	syncer  EdsSyncer
	ready   chan struct{}
}

func NewEdsEventLoop(emitter EdsSnapshotEmitter, syncer EdsSyncer) eventloop.EventLoop {
	return &edsEventLoop{
		emitter: emitter,
		syncer:  syncer,
		ready:   make(chan struct{}),
	}
}

func (el *edsEventLoop) Ready() <-chan struct{} {
	return el.ready
}

func (el *edsEventLoop) Run(namespaces []string, opts clients.WatchOpts) (<-chan error, error) {
	opts = opts.WithDefaults()
	opts.Ctx = contextutils.WithLogger(opts.Ctx, "v1.event_loop")
	logger := contextutils.LoggerFrom(opts.Ctx)
	logger.Infof("event loop started")

	errs := make(chan error)

	watch, emitterErrs, err := el.emitter.Snapshots(namespaces, opts)
	if err != nil {
		return nil, errors.Wrapf(err, "starting snapshot watch")
	}
	go errutils.AggregateErrs(opts.Ctx, errs, emitterErrs, "v1.emitter errors")
	go func() {
		var channelClosed bool

		// create a new context for each loop, cancel it before each loop
		var cancel context.CancelFunc = func() {}

		// use closure to allow cancel function to be updated as context changes
		defer func() { cancel() }()

		// cache the previous snapshot for comparison
		var previousSnapshot *EdsSnapshot

		for {
			select {
			case snapshot, ok := <-watch:
				if !ok {
					return
				}

				if syncDecider, isDecider := el.syncer.(EdsSyncDecider); isDecider {
					if shouldSync := syncDecider.ShouldSync(previousSnapshot, snapshot); !shouldSync {
						continue // skip syncing this syncer
					}
				} else if syncDeciderWithContext, isDecider := el.syncer.(EdsSyncDeciderWithContext); isDecider {
					if shouldSync := syncDeciderWithContext.ShouldSync(opts.Ctx, previousSnapshot, snapshot); !shouldSync {
						continue // skip syncing this syncer
					}
				}

				// cancel any open watches from previous loop
				cancel()

				startTime := time.Now()
				ctx, span := trace.StartSpan(opts.Ctx, "eds.gloo.solo.io.EventLoopSync")
				ctx, canc := context.WithCancel(ctx)
				cancel = canc
				err := el.syncer.Sync(ctx, snapshot)
				stats.RecordWithTags(
					ctx,
					[]tag.Mutator{
						tag.Insert(skstats.SyncerNameKey, fmt.Sprintf("%T", el.syncer)),
					},
					mEdsSnapshotTimeSec.M(time.Now().Sub(startTime).Seconds()),
				)
				span.End()

				if err != nil {
					select {
					case errs <- err:
					default:
						logger.Errorf("write error channel is full! could not propagate err: %v", err)
					}
				} else if !channelClosed {
					channelClosed = true
					close(el.ready)
				}
			case <-opts.Ctx.Done():
				return
			}
		}
	}()
	return errs, nil
}
