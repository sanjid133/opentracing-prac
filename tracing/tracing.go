package tracing

import (
	"io"
	"github.com/uber/jaeger-client-go/config"
	"github.com/uber/jaeger-client-go"
	"fmt"
	"github.com/opentracing/opentracing-go"
)

func Init(service string)(opentracing.Tracer, io.Closer) {
	cfg := config.Configuration{
		Sampler: &config.SamplerConfig{
			Type: "const",
			Param: 1,
		},
		Reporter: &config.ReporterConfig{
			LogSpans: true,
		},
	}

	tracer, closer, err := cfg.New(service, config.Logger(jaeger.StdLogger))
	if err != nil {
		panic(fmt.Sprintf("Failed to init jaeger, err= %v", err))
	}
	return tracer, closer
}
