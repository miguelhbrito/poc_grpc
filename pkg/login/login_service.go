package login

import (
	"context"
	"fmt"

	proto "github.com/poc_grpc/pb"
	"github.com/poc_grpc/pkg/api/entity"
	"github.com/poc_grpc/pkg/mcontext"
	"github.com/poc_grpc/pkg/mlog"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type LoginService struct {
	proto.UnimplementedLoginServer
	Manager Login
}

func (l LoginService) CreateLogin(ctx context.Context, req *proto.CreateLoginRequest) (*proto.CreateLoginResponse, error) {
	mctx := mcontext.NewFrom(ctx)
	mlog.Info(mctx).Msg("Received request to create an new login")
	lrEntity := entity.Login{}.GenerateEntity(req)
	err := l.Manager.Create(mctx, lrEntity)
	if err != nil {
		return nil, status.Error(codes.Internal, fmt.Sprintf("Fail to save an new user, error: %v", err))
	}
	return &proto.CreateLoginResponse{}, nil
}

func (l LoginService) LoginSystem(ctx context.Context, req *proto.CreateLoginRequest) (*proto.CreateLoginSystemResponse, error) {
	mctx := mcontext.NewFrom(ctx)
	mlog.Info(mctx).Msg("Received request to login into system")
	lrEntity := entity.Login{}.GenerateEntity(req)
	token, err := l.Manager.Login(mctx, lrEntity)
	if err != nil {
		return nil, status.Error(codes.Internal, fmt.Sprintf("Error to get token, error: %v", err))
	}
	return &proto.CreateLoginSystemResponse{
		Token: token.Token,
		Ttl:   token.ExpTime,
	}, nil
}
