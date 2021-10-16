package entity

import (
	proto "github.com/poc_grpc/pb"
)

type LoginReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResp struct {
	Token string `json:"token"`
	TTL   int64  `json:"ttl"`
}

func GrpcLgToEntity(login *proto.CreateLoginRequest) LoginReq {
	return LoginReq{
		Username: login.Username,
		Password: login.Password,
	}
}
