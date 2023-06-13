package rabbitmq

import (
	"github.com/go-playground/validator"

	"github.com/mehdihadeli/go-ecommerce-microservices/internal/pkg/logger"
	"github.com/mehdihadeli/go-ecommerce-microservices/internal/pkg/messaging/consumer"
	rabbitmqConfigurations "github.com/mehdihadeli/go-ecommerce-microservices/internal/pkg/rabbitmq/configurations"
	"github.com/mehdihadeli/go-ecommerce-microservices/internal/pkg/rabbitmq/consumer/configurations"

	createProductExternalEventV1 "github.com/mehdihadeli/go-ecommerce-microservices/internal/services/catalogs/read_service/internal/products/features/creating_product/v1/events/integration_events/external_events"
	deleteProductExternalEventV1 "github.com/mehdihadeli/go-ecommerce-microservices/internal/services/catalogs/read_service/internal/products/features/deleting_products/v1/events/integration_events/external_events"
	updateProductExternalEventsV1 "github.com/mehdihadeli/go-ecommerce-microservices/internal/services/catalogs/read_service/internal/products/features/updating_products/v1/events/integration_events/external_events"
)

func ConfigProductsRabbitMQ(
	builder rabbitmqConfigurations.RabbitMQConfigurationBuilder,
	logger logger.Logger,
	validator *validator.Validate,
) {
	// add custom message type mappings
	// utils.RegisterCustomMessageTypesToRegistrty(map[string]types.IMessage{"productCreatedV1": &creatingProductIntegration.ProductCreatedV1{}})

	builder.
		AddConsumer(
			createProductExternalEventV1.ProductCreatedV1{},
			func(builder configurations.RabbitMQConsumerConfigurationBuilder) {
				builder.WithHandlers(
					func(handlersBuilder consumer.ConsumerHandlerConfigurationBuilder) {
						handlersBuilder.AddHandler(
							createProductExternalEventV1.NewProductCreatedConsumer(
								logger,
								validator,
							),
						)
					},
				)
			}).
		AddConsumer(
			deleteProductExternalEventV1.ProductDeletedV1{},
			func(builder configurations.RabbitMQConsumerConfigurationBuilder) {
				builder.WithHandlers(
					func(handlersBuilder consumer.ConsumerHandlerConfigurationBuilder) {
						handlersBuilder.AddHandler(
							deleteProductExternalEventV1.NewProductDeletedConsumer(
								logger,
								validator,
							),
						)
						deleteProductExternalEventV1.NewProductDeletedConsumer(logger, validator)
					},
				)
			}).
		AddConsumer(
			updateProductExternalEventsV1.ProductUpdatedV1{},
			func(builder configurations.RabbitMQConsumerConfigurationBuilder) {
				builder.WithHandlers(
					func(handlersBuilder consumer.ConsumerHandlerConfigurationBuilder) {
						handlersBuilder.AddHandler(
							updateProductExternalEventsV1.NewProductUpdatedConsumer(
								logger,
								validator,
							),
						)
						updateProductExternalEventsV1.NewProductUpdatedConsumer(logger, validator)
					},
				)
			})
}
