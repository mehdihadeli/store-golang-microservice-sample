package otel

import (
	"context"

	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/fx"

	"github.com/mehdihadeli/go-ecommerce-microservices/internal/pkg/otel/config"
	"github.com/mehdihadeli/go-ecommerce-microservices/internal/pkg/otel/metrics"
	"github.com/mehdihadeli/go-ecommerce-microservices/internal/pkg/otel/tracing"
)

// Module provided to fxlog
// https://uber-go.github.io/fx/modules.html
var Module = fx.Module(
	"otelfx",
	fx.Provide(
		config.ProvideOtelConfig,
		metrics.NewOtelMetrics,
		tracing.NewOtelTracing,
		fx.Annotate(
			provideMeter,
			fx.As(new(metric.Meter))),
		fx.Annotate(
			provideTracer,
			fx.As(new(tracing.AppTracer)),
			fx.As(new(trace.Tracer)),
		),
	),
	fx.Invoke(registerHooks),
)

func provideMeter(otelMetrics *metrics.OtelMetrics) metric.Meter {
	return otelMetrics.Meter
}

func provideTracer(tracingOtel *tracing.TracingOpenTelemetry) tracing.AppTracer {
	return tracingOtel.AppTracer
}

// we don't want to register any dependencies here, its func body should execute always even we don't request for that, so we should use `invoke`
func registerHooks(
	lc fx.Lifecycle,
	metrics *metrics.OtelMetrics,
	tracingOtel *tracing.TracingOpenTelemetry,
) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go func() {
				if err := metrics.Run(); err != nil {
					// do a fatal for running OnStop hook
					metrics.Logger.Fatalf(
						"(metrics.RunHttpServer) error in running metrics server: {%v}",
						err,
					)
				}
			}()
			metrics.Logger.Infof(
				"Metrics server %s is listening on Host:{%s} Http PORT: {%s}",
				metrics.Config.OTelMetricsOptions.Name,
				metrics.Config.OTelMetricsOptions.Host,
				metrics.Config.OTelMetricsOptions.Port,
			)

			return nil
		},
		OnStop: func(ctx context.Context) error {
			if err := tracingOtel.TracerProvider.Shutdown(ctx); err != nil {
				metrics.Logger.Errorf("error in shutting down trace provider: %v", err)
			} else {
				metrics.Logger.Info("trace provider shutdown gracefully")
			}

			if err := metrics.GracefulShutdown(ctx); err != nil {
				metrics.Logger.Errorf("error shutting down metrics server: %v", err)
			} else {
				metrics.Logger.Info("metrics server shutdown gracefully")
			}
			return nil
		},
	})
}
