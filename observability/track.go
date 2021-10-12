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
	ServiceName string
	Infos       map[string]string
}

func SpanByGprc(ctx context.Context, sp MySpan) mcontext.Context {
	md := ExtractIncoming(ctx).Clone()
	parentSpanContext, err := opentracing.GlobalTracer().Extract(opentracing.HTTPHeaders, MetadataTextMap(md))
	if err != nil && err != opentracing.ErrSpanContextNotFound {
		mlog.Info(mcontext.NewFrom(ctx)).Err(err).Msg("failed")
	}

	serverSpan := opentracing.GlobalTracer().StartSpan(
		sp.ServiceName,
		ext.RPCServerOption(parentSpanContext),
	)

	openCtx := opentracing.ContextWithSpan(ctx, serverSpan)
	ctxResult := mcontext.NewFrom(openCtx)
	ctxResult = mcontext.WithValue(ctxResult, models.TrackingIdCtxKey, models.TrackingId(getTrackingID(ctxResult)))

	for k, v := range sp.Infos {
		serverSpan.SetTag(k, v)
	}

	return ctxResult
}

func CreateSpan(ctx mcontext.Context, mySpan MySpan) mcontext.Context {
	trackingId := getTrackingID(ctx)
	var span opentracing.Span
	var openCtx context.Context

	if trackingId == "" {
		span, openCtx = opentracing.StartSpanFromContext(ctx, mySpan.ServiceName)
		spanCtx, ok := span.Context().(jaeger.SpanContext)
		if !ok {
			putLog(ctx, "error", "erroro to create span !")
			return mcontext.NewContext()
		}
		trackingId = spanCtx.TraceID().String()
	} else {
		span, openCtx = opentracing.StartSpanFromContext(ctx, mySpan.ServiceName)
	}

	mctx := mcontext.NewFrom(openCtx)
	mctx = mcontext.WithValue(mctx, models.TrackingIdCtxKey, trackingId)

	for k, v := range mySpan.Infos {
		span.SetTag(k, v)
	}
	return mctx
}

func Finish(ctx mcontext.Context) {
	span := opentracing.SpanFromContext(ctx)
	if span == nil {
		mlog.Info(ctx).Msg("didnt not found span !")
		return
	}
	span.Finish()
}

func putLog(ctx mcontext.Context, alternatingKeyValue ...interface{}) {
	span := opentracing.SpanFromContext(ctx)
	if span == nil {
		return
	}
	span.LogKV(alternatingKeyValue...)
}

func getTrackingID(ctx mcontext.Context) string {
	spanCtx := opentracing.SpanFromContext(ctx)
	if spanCtx == nil {
		return ""
	}
	if spanCtxResult, ok := spanCtx.Context().(jaeger.SpanContext); ok {
		return spanCtxResult.TraceID().String()
	}
	return ""
}
