package middleware

import (
	"context"
	"encoding/base64"
	"fmt"
	"strings"

	"github.com/poc_grpc/mcontext"
	"github.com/poc_grpc/mlog"
	"github.com/poc_grpc/models"
	"github.com/poc_grpc/observability"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func Interceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (reply interface{}, err error) {
		md := observability.ExtractIncoming(ctx)
		tags := make(map[string]string)
		for k, v := range md {
			tags[k] = strings.Join(v, "")
		}

		mCtx := mcontext.NewFrom(ctx)
		mctx := observability.SpanByGprc(mCtx, observability.MySpan{
			ServiceName: info.FullMethod,
			Infos:       tags,
		})

		authCheck, ok := tags[string(models.AuthorizationCtxKey)]
		if authCheck != "" && ok {
			authBytes, err := base64.StdEncoding.DecodeString(authCheck)
			if err != nil {
				return "", status.Error(codes.Internal, "Error to decode base64")
			}
			if string(authBytes) != "gandalf:mithrandir" {
				return "", status.Error(codes.PermissionDenied, "User not allowed")
			}
			username := strings.Split(string(authBytes), ":")
			mctx = mcontext.WithValue(mctx, models.UsernameCtxKey, models.Username(username[0]))
		}

		mlog.Info(mctx).Msgf("Grpc-Server tags: %v", tags)
		mlog.Info(mctx).Msgf(fmt.Sprintf("fullmethod : %s", info.FullMethod))

		return handler(mctx, req)
	}
}
