package middleware

import (
	"context"
	"fmt"
	"strings"

	"github.com/golang-jwt/jwt"
	"github.com/poc_grpc/pkg/api"
	"github.com/poc_grpc/pkg/login"
	"github.com/poc_grpc/pkg/mcontext"
	"github.com/poc_grpc/pkg/mlog"
	"github.com/poc_grpc/pkg/observability"

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

		if info.FullMethod != "/Login/CreateLogin" && info.FullMethod != "/Login/LoginSystem" {
			mlog.Debug(mctx).Msgf("Authorization middleware checking token auth")

			tokenAuth, ok := tags[string(api.AuthorizationCtxKey)]
			if !ok {
				return "", status.Error(codes.FailedPrecondition, "Token was not passed into request")
			}
			token, err := jwt.Parse(tokenAuth, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
				}
				return login.JwtKey, nil
			})
			if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
				username := claims["username"]
				usernameString := fmt.Sprintf("%s", username)
				mctx = mcontext.WithValue(mctx, "props", claims)
				mctx = mcontext.WithValue(mctx, api.UsernameCtxKey, api.Username(usernameString))
				mctx = mcontext.WithValue(mctx, api.AuthorizationCtxKey, tokenAuth)
				return handler(mctx, req)
			} else {

				mlog.Error(mctx).Msgf("Error on decode token, err: %v", err)
				return "", status.Error(codes.PermissionDenied, "User not allowed")
			}

		}

		mlog.Info(mctx).Msgf("Auth middleware")
		mlog.Info(mctx).Msgf("Grpc-Server tags: %v", tags)
		mlog.Info(mctx).Msgf(fmt.Sprintf("fullmethod : %s", info.FullMethod))

		return handler(mctx, req)
	}
}
