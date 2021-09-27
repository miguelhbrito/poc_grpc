package service

import (
	"context"

	proto "github.com/poc_grpc/pb"
)

type loginService struct {
	proto.UnimplementedLoginServer
}

func (l loginService) CreateLogin(ctx context.Context, req *proto.CreateLoginRequest) (*proto.CreateLoginResponse, error) {
	return &proto.CreateLoginResponse{
		Token: "test",
		Ttl:   123,
	}, nil
}
