package middleware

import (
	"context"
	"encoding/base64"
	"fmt"
	"strings"

	"github.com/poc_grpc/api"
	"github.com/poc_grpc/mcontext"
	"github.com/poc_grpc/mlog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func Interceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (reply interface{}, err error) {
		md := ExtractIncoming(ctx)
		tags := make(map[string]string)
		for k, v := range md {
			tags[k] = strings.Join(v, "")
		}

		mctx := mcontext.NewFrom(ctx)
		username := tags[string(api.UsernameCtxKey)]
		//token := tags[string(api.AuthorizationCtxKey)]

		mctx = mcontext.WithValue(mctx, api.UsernameCtxKey, api.Username(username))

		mlog.Info(mctx).Msgf("GrpcServerHandler tags: %v", tags)
		mlog.Info(mctx).Msgf(fmt.Sprintf("fullmethod : %s", info.FullMethod))

		grpcBasicAuth, ok := tags[string(api.AuthorizationCtxKey)]
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

		return handler(mctx, req)
	}
}
