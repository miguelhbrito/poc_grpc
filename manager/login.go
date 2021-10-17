package manager

import (
	"github.com/Nerzal/gocloak/v9"
	"github.com/poc_grpc/mcontext"
	"github.com/poc_grpc/models/entity"
)

const (
	HostName     = "http://localhost:8080"
	Realm        = "Valhalla"
	ClientID     = "valhalla-login"
	ClientSecret = "b0aab250-074a-4dae-b3cd-a43fe730c415"
	ConfigURL    = "http://localhost:8080/auth/realms/Valhalla"
)

func Login(mctx mcontext.Context, loginReq entity.LoginReq) (entity.LoginResp, error) {
	client := gocloak.NewClient(HostName)
	token, err := client.Login(mctx, ClientID, ClientSecret, Realm, loginReq.Username, loginReq.Password)
	if err != nil {
		return entity.LoginResp{}, err
	}
	return entity.LoginResp{
		Token: token.AccessToken,
		TTL:   1000000,
	}, nil
}
