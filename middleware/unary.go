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

		mctx := observability.SpanByGprc(ctx, observability.MySpan{
			FullMethod: info.FullMethod,
			Infos:      tags,
		})

		username := tags[string(models.UsernameCtxKey)]
		mctx = mcontext.WithValue(mctx, models.UsernameCtxKey, models.Username(username))

		grpcBasicAuth, ok := tags[string(models.AuthorizationCtxKey)]
		if grpcBasicAuth != "" && ok {
			xAuthBytes, err := base64.StdEncoding.DecodeString(grpcBasicAuth)
			if err != nil {
				return "", status.Error(codes.Internal, "Error to decode base64")
			}
			if string(xAuthBytes) != "user:password" {
				return "", status.Error(codes.PermissionDenied, "User not allowed")
			}
			return handler(mctx, req)
		}

		mlog.Info(mctx).Msgf("GrpcServerHandler tags: %v", tags)
		mlog.Info(mctx).Msgf(fmt.Sprintf("fullmethod : %s", info.FullMethod))

		return handler(mctx, req)
	}
}
