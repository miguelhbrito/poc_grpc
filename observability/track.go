package observability

import (
	"context"

	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"github.com/poc_grpc/mcontext"
	"github.com/poc_grpc/mlog"
	"github.com/poc_grpc/models"
	"github.com/uber/jaeger-client-go"
)

type MySpan struct {
	FullMethod string
	Infos      map[string]string
}

func SpanByGprc(ctx context.Context, sp MySpan) mcontext.Context {
	md := ExtractIncoming(ctx).Clone()
	parentSpanContext, err := opentracing.GlobalTracer().Extract(opentracing.HTTPHeaders, MetadataTextMap(md))
	if err != nil && err != opentracing.ErrSpanContextNotFound {
		mlog.Info(mcontext.NewFrom(ctx)).Err(err).Msg("failed")
	}
	serverSpan := opentracing.GlobalTracer().StartSpan(
		sp.FullMethod,
		ext.RPCServerOption(parentSpanContext),
	)

	var tid string
	if sc, ok := serverSpan.Context().(jaeger.SpanContext); ok {
		tid = sc.TraceID().String()
	}

	openCtx := opentracing.ContextWithSpan(ctx, serverSpan)
	ctxResult := mcontext.NewFrom(openCtx)
	ctxResult = mcontext.WithValue(ctxResult, models.TrackingIdCtxKey, models.TrackingId(tid))

	for k, v := range sp.Infos {
		serverSpan.SetTag(k, v)
	}

	return nil
}
