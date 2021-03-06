package observability

import (
	"fmt"
	"io"
	"time"

	"github.com/uber/jaeger-client-go/zipkin"

	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go/config"
	jaegerlog "github.com/uber/jaeger-client-go/log"
)

func InitJaeger(service string) io.Closer {
	zipkinPropagator := zipkin.NewZipkinB3HTTPHeaderPropagator()
	cfg := config.Configuration{
		ServiceName: service,
		Sampler: &config.SamplerConfig{
			//already in default
			SamplingServerURL: "http://127.0.0.1:5778/sampling",
			Type:              "const",
			Param:             1.0,
		},
		Reporter: &config.ReporterConfig{
			LogSpans:            true,
			BufferFlushInterval: 1 * time.Millisecond,
			LocalAgentHostPort:  "127.0.0.1:6831",
		},
	}

	jLogger := jaegerlog.StdLogger

	tracer, closer, err := cfg.NewTracer(
		config.Logger(jLogger),
		config.ZipkinSharedRPCSpan(true),
		config.Injector(opentracing.HTTPHeaders, zipkinPropagator),
		config.Extractor(opentracing.HTTPHeaders, zipkinPropagator))
	if err != nil {
		panic(fmt.Sprintf("ERROR: cannot init Jaeger: %v\n", err))
	}

	opentracing.SetGlobalTracer(tracer)

	return closer
}
