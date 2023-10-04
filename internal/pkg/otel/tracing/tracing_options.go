package tracing

import (
	"github.com/mehdihadeli/go-ecommerce-microservices/internal/pkg/config"
	"github.com/mehdihadeli/go-ecommerce-microservices/internal/pkg/config/environemnt"
	typeMapper "github.com/mehdihadeli/go-ecommerce-microservices/internal/pkg/reflection/type_mappper"

	"github.com/iancoleman/strcase"
)

type OTLPProvider struct {
	Name         string            `mapstructure:"name"`
	Enabled      bool              `mapstructure:"enabled"`
	OTLPEndpoint string            `mapstructure:"otlpEndpoint"`
	OTLPHeaders  map[string]string `mapstructure:"otlpHeaders"`
}

type TracingOptions struct {
	Enabled                   bool                   `mapstructure:"enabled"`
	ServiceName               string                 `mapstructure:"serviceName"`
	Version                   string                 `mapstructure:"version"`
	InstrumentationName       string                 `mapstructure:"instrumentationName"`
	Id                        int64                  `mapstructure:"id"`
	AlwaysOnSampler           bool                   `mapstructure:"alwaysOnSampler"`
	ZipkinExporterOptions     *ZipkinExporterOptions `mapstructure:"zipkinExporterOptions"`
	JaegerExporterOptions     *OTLPProvider          `mapstructure:"jaegerExporterOptions"`
	ElasticApmExporterOptions *OTLPProvider          `mapstructure:"elasticApmExporterOptions"`
	UptraceExporterOptions    *OTLPProvider          `mapstructure:"uptraceExporterOptions"`
	SignozExporterOptions     *OTLPProvider          `mapstructure:"signozExporterOptions"`
	TempoExporterOptions      *OTLPProvider          `mapstructure:"tempoExporterOptions"`
	UseStdout                 bool                   `mapstructure:"useStdout"`
	UseOTLP                   bool                   `mapstructure:"useOTLP"`
	OTLPProviders             []OTLPProvider         `mapstructure:"otlpProviders"`
}

type ZipkinExporterOptions struct {
	Url string `mapstructure:"url"`
}

func ProvideTracingConfig(
	environment environemnt.Environment,
) (*TracingOptions, error) {
	optionName := strcase.ToLowerCamel(
		typeMapper.GetTypeNameByT[TracingOptions](),
	)

	return config.BindConfigKey[*TracingOptions](optionName, environment)
}
