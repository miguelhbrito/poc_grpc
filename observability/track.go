package observability

import (
	"context"

	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"github.com/poc_grpc/mcontext"
	"github.com/poc_grpc/middleware"
	"github.com/poc_grpc/mlog"
)

func SpanByGprc(ctx context.Context, ts Span) mcontext.Context {
	md := middleware.ExtractIncoming(ctx).Clone()
	parentSpanContext, err := opentracing.GlobalTracer().Extract(opentracing.HTTPHeaders, middleware.MetadataTextMap(md))
	if err != nil && err != opentracing.ErrSpanContextNotFound {
		mlog.Info(mcontext.NewFrom(ctx)).Err(err).Msg("failed parsing trace information")
	}
	serverSpan := opentracing.GlobalTracer().StartSpan(
		ts.OperationName,
		ext.RPCServerOption(parentSpanContext),
	)
	injectOpentracingIdsToTags("tracking-server", serverSpan, Extract(ctx))

	openCtx := opentracing.ContextWithSpan(ctx, serverSpan)
	finalCtx := dcontext.NewFrom(openCtx)
	finalCtx = dcontext.WithValue(finalCtx, api.TraceIDCtxKey, t.GetTraceID(finalCtx))
	finalCtx = dcontext.WithValue(finalCtx, api.ServiceNameCtxKey, t.serviceName)
	return finalCtx
}
