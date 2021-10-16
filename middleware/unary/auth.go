package middleware

import (
	"context"
	"fmt"
	"strings"

	"github.com/Nerzal/gocloak/v9"
	"github.com/poc_grpc/manager"
	"github.com/poc_grpc/mcontext"
	"github.com/poc_grpc/mlog"
	"github.com/poc_grpc/models"
	"github.com/poc_grpc/observability"
	"google.golang.org/grpc"
)

func Authorization() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (reply interface{}, err error) {
		md := observability.ExtractIncoming(ctx)
		tags := make(map[string]string)
		for k, v := range md {
			tags[k] = strings.Join(v, "")
		}
		mctx := mcontext.NewFrom(ctx)

		if info.FullMethod != "/Login/CreateLogin" {
			authCheck, ok := tags[string(models.AuthorizationCtxKey)]
			if authCheck != "" && ok {
				client := gocloak.NewClient(manager.HostName)
				token, err := client.LoginClient(mctx, manager.ClientID, manager.ClientSecret, manager.Realm)
				if err != nil {
					panic("Login failed:" + err.Error())
				}
				client.LoginClientSignedJWT()
				token.IDToken
				username := strings.Split(string(authBytes), ":")
				mctx = mcontext.WithValue(mctx, models.UsernameCtxKey, models.Username(username[0]))
			}

		}

		mlog.Info(mctx).Msgf("Grpc-Server tags: %v", tags)
		mlog.Info(mctx).Msgf(fmt.Sprintf("fullmethod : %s", info.FullMethod))

		return handler(mctx, req)
	}
}
