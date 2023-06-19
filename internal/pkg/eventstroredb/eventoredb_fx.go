package eventstroredb

import (
	"context"

	"github.com/EventStore/EventStore-Client-Go/esdb"
	"go.uber.org/fx"

	"github.com/mehdihadeli/go-ecommerce-microservices/internal/pkg/eventstroredb/config"
	"github.com/mehdihadeli/go-ecommerce-microservices/internal/pkg/logger"
)

// Module provided to fxlog
// https://uber-go.github.io/fx/modules.html
var Module = func(projectionBuilderConstructor interface{}) fx.Option {
	return fx.Module(
		"eventstoredbfx",
		// - order is not important in provide
		// - provide can have parameter and will resolve if registered
		// - execute its func only if it requested
		fx.Provide(
			config.ProvideConfig,
			NewEsdbSerializer,
			NewEventStoreDB,
			NewEventStoreDbEventStore,
			NewEsdbSubscriptionCheckpointRepository,
			NewEsdbSubscriptionAllWorker,
		),
		fx.Provide(projectionBuilderConstructor),
		// - execute after registering all of our provided
		// - they execute by their orders
		// - invokes always execute its func compare to provides that only run when we request for them.
		// - return value will be discarded and can not be provided
		fx.Invoke(registerHooks),
	)
}

// we don't want to register any dependencies here, its func body should execute always even we don't request for that, so we should use `invoke`
func registerHooks(
	lc fx.Lifecycle,
	worker EsdbSubscriptionAllWorker,
	logger logger.Logger,
	cfg *config.EventStoreDbOptions,
) {
	lifeTimeCtx := context.Background()

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			// https://github.com/uber-go/fx/blob/v1.20.0/app.go#L573
			// this ctx is just for startup dependencies setup and OnStart callbacks, and it has short timeout 15s, and it is not alive in whole lifetime app
			// if we need an app context which is alive until the app context done we should create it manually here
			go func() {
				option := &EventStoreDBSubscriptionToAllOptions{
					FilterOptions: &esdb.SubscriptionFilter{
						Type:     esdb.StreamFilterType,
						Prefixes: cfg.Subscription.Prefix,
					},
					SubscriptionId: cfg.Subscription.SubscriptionId,
				}

				if err := worker.SubscribeAll(lifeTimeCtx, option); err != nil {
					logger.Fatalf(
						"(EsdbSubscriptionAllWorker.Start) error in running esdb subscription worker: {%v}",
						err,
					)
				}
			}()
			logger.Info("esdb subscription worker is listening.")

			return nil
		},
		OnStop: func(ctx context.Context) error {
			// https://github.com/uber-go/fx/blob/v1.20.0/app.go#L573
			// this ctx is just for stopping callbacks or OnStop callbacks, and it has short timeout 15s, and it is not alive in whole lifetime app
			lifeTimeCtx.Done()

			return nil
		},
	})
}
