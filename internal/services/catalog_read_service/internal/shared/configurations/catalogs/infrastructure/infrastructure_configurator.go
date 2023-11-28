package infrastructure

import (
	"github.com/mehdihadeli/go-ecommerce-microservices/internal/pkg/fxapp/contracts"
	"github.com/mehdihadeli/go-ecommerce-microservices/internal/pkg/logger"
	loggingpipelines "github.com/mehdihadeli/go-ecommerce-microservices/internal/pkg/logger/pipelines"
	"github.com/mehdihadeli/go-ecommerce-microservices/internal/pkg/otel/metrics"
	metricspipelines "github.com/mehdihadeli/go-ecommerce-microservices/internal/pkg/otel/metrics/mediatr/pipelines"
	"github.com/mehdihadeli/go-ecommerce-microservices/internal/pkg/otel/tracing"
	tracingpipelines "github.com/mehdihadeli/go-ecommerce-microservices/internal/pkg/otel/tracing/mediatr/pipelines"
	postgrespipelines "github.com/mehdihadeli/go-ecommerce-microservices/internal/pkg/postgresGorm/mediatr/pipelines"

	"github.com/mehdihadeli/go-mediatr"
	"gorm.io/gorm"
)

type InfrastructureConfigurator struct {
	contracts.Application
}

func NewInfrastructureConfigurator(
	app contracts.Application,
) *InfrastructureConfigurator {
	return &InfrastructureConfigurator{
		Application: app,
	}
}

func (ic *InfrastructureConfigurator) ConfigInfrastructures() {
	ic.ResolveFunc(
		func(l logger.Logger, tracer tracing.AppTracer, metrics metrics.AppMetrics, db *gorm.DB) error {
			err := mediatr.RegisterRequestPipelineBehaviors(
				loggingpipelines.NewMediatorLoggingPipeline(l),
				tracingpipelines.NewMediatorTracingPipeline(
					tracer,
					tracingpipelines.WithLogger(l),
				),
				metricspipelines.NewMediatorMetricsPipeline(
					metrics,
					metricspipelines.WithLogger(l),
				),
				postgrespipelines.NewMediatorTransactionPipeline(l, db),
			)

			return err
		},
	)
}
