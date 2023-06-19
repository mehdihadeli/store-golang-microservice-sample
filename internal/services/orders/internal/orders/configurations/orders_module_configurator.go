package configurations

import (
	"github.com/go-playground/validator"
	googleGrpc "google.golang.org/grpc"

	"github.com/mehdihadeli/go-ecommerce-microservices/internal/pkg/es/contracts/store"
	"github.com/mehdihadeli/go-ecommerce-microservices/internal/pkg/fxapp"
	grpcServer "github.com/mehdihadeli/go-ecommerce-microservices/internal/pkg/grpc"
	customEcho "github.com/mehdihadeli/go-ecommerce-microservices/internal/pkg/http/custom_echo"
	"github.com/mehdihadeli/go-ecommerce-microservices/internal/pkg/logger"
	"github.com/mehdihadeli/go-ecommerce-microservices/internal/pkg/otel/tracing"
	"github.com/mehdihadeli/go-ecommerce-microservices/internal/pkg/web/route"
	"github.com/mehdihadeli/go-ecommerce-microservices/internal/services/orders/internal/orders/configurations/mappings"
	"github.com/mehdihadeli/go-ecommerce-microservices/internal/services/orders/internal/orders/configurations/mediatr"
	ordersservice "github.com/mehdihadeli/go-ecommerce-microservices/internal/services/orders/internal/orders/contracts/proto/service_clients"
	"github.com/mehdihadeli/go-ecommerce-microservices/internal/services/orders/internal/orders/contracts/repositories"
	"github.com/mehdihadeli/go-ecommerce-microservices/internal/services/orders/internal/orders/grpc"
	"github.com/mehdihadeli/go-ecommerce-microservices/internal/services/orders/internal/orders/models/orders/aggregate"
	"github.com/mehdihadeli/go-ecommerce-microservices/internal/services/orders/internal/shared/contracts"
)

type OrdersModuleConfigurator struct {
	*fxapp.Application
}

func NewOrdersModuleConfigurator(
	fxapp *fxapp.Application,
) *OrdersModuleConfigurator {
	return &OrdersModuleConfigurator{
		Application: fxapp,
	}
}

func (c *OrdersModuleConfigurator) ConfigureOrdersModule() {
	c.ResolveFunc(
		func(logger logger.Logger,
			server customEcho.EchoHttpServer,
			orderRepository repositories.OrderMongoRepository,
			orderAggregateStore store.AggregateStore[*aggregate.Order],
			tracer tracing.AppTracer,
		) error {
			// Config Orders Mappings
			err := mappings.ConfigureOrdersMappings()
			if err != nil {
				return err
			}

			// Config Orders Mediators
			err = mediatr.ConfigOrdersMediator(logger, orderRepository, orderAggregateStore, tracer)
			if err != nil {
				return err
			}

			return nil
		},
	)
}

func (c *OrdersModuleConfigurator) MapOrdersEndpoints() {
	// Config Orders Http Endpoints
	c.ResolveFuncWithParamTag(func(endpoints []route.Endpoint) {
		for _, endpoint := range endpoints {
			endpoint.MapEndpoint()
		}
	}, `group:"order-routes"`,
	)

	// Config Orders Grpc Endpoints
	c.ResolveFunc(
		func(ordersGrpcServer grpcServer.GrpcServer, ordersMetrics *contracts.OrdersMetrics, logger logger.Logger, validator *validator.Validate) error {
			orderGrpcService := grpc.NewOrderGrpcService(logger, validator, ordersMetrics)
			ordersGrpcServer.GrpcServiceBuilder().RegisterRoutes(func(server *googleGrpc.Server) {
				ordersservice.RegisterOrdersServiceServer(server, orderGrpcService)
			})
			return nil
		},
	)
}
