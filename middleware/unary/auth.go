package middleware

import (
	"context"
	"fmt"
	"strings"

	"github.com/coreos/go-oidc/v3/oidc"
	"github.com/poc_grpc/manager"
	"github.com/poc_grpc/mcontext"
	"github.com/poc_grpc/mlog"
	"github.com/poc_grpc/models"
	"github.com/poc_grpc/observability"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func Authorization() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (reply interface{}, err error) {
		md := observability.ExtractIncoming(ctx)
		tags := make(map[string]string)
		for k, v := range md {
			tags[k] = strings.Join(v, "")
		}
		mctx := mcontext.NewFrom(ctx)

		provider, err := oidc.NewProvider(mctx, manager.ConfigURL)
		if err != nil {
			panic(err)
		}

		oidcConfig := &oidc.Config{
			ClientID: manager.ClientID,
		}
		verifier := provider.Verifier(oidcConfig)

		if info.FullMethod != "/Login/CreateLogin" {
			authCheck, ok := tags[string(models.AuthorizationCtxKey)]
			if authCheck != "" && ok {
				_, err = verifier.Verify(ctx, authCheck)
				if err != nil {
					return nil, status.Error(codes.PermissionDenied,
						fmt.Sprintf("Fail to login into the system, error on verify, error: %v", err))
				}
			}

		}

		mlog.Info(mctx).Msgf("Auth middleware")
		mlog.Info(mctx).Msgf("Grpc-Server tags: %v", tags)
		mlog.Info(mctx).Msgf(fmt.Sprintf("fullmethod : %s", info.FullMethod))

		return handler(mctx, req)
	}
}
