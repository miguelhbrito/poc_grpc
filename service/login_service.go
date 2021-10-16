package service

import (
	"context"
	"fmt"

	"github.com/poc_grpc/manager"
	"github.com/poc_grpc/mcontext"
	"github.com/poc_grpc/mlog"
	"github.com/poc_grpc/models/entity"
	proto "github.com/poc_grpc/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type LoginService struct {
	proto.UnimplementedLoginServer
}

func (l LoginService) CreateLogin(ctx context.Context, req *proto.CreateLoginRequest) (*proto.CreateLoginResponse, error) {
	mctx := mcontext.NewFrom(ctx)
	mlog.Info(mctx).Msg("Received request to login")
	loginEntity := entity.GrpcLgToEntity(req)
	loginResponse, err := manager.Login(mctx, loginEntity)
	if err != nil {
		return nil, status.Error(codes.PermissionDenied, fmt.Sprintf("Fail to login into the system, error: %v", err))
	}
	return &proto.CreateLoginResponse{
		Token: loginResponse.Token,
		Ttl:   loginResponse.TTL,
	}, nil
}
