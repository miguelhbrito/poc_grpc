package service

import (
	"context"
	"fmt"

	"github.com/poc_grpc/api/entity"
	"github.com/poc_grpc/mcontext"
	"github.com/poc_grpc/mlog"
	proto "github.com/poc_grpc/pb"
	"github.com/poc_grpc/storage"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type NotebookService struct {
	proto.UnimplementedNotebookServiceServer
}

func (n NotebookService) CreateNotebook(ctx context.Context, req *proto.CreateNotebookRequest) (*proto.CreateNotebookResponse, error) {
	mctx := mcontext.NewFrom(ctx)
	mlog.Info(mctx).Msg("Received request to create a notebook")
	nbEntity := entity.GrpcToEntity(req)
	err := storage.SaveNotebook(mctx, nbEntity)
	if err != nil {
		return nil, status.Error(codes.Internal, fmt.Sprint("Fail to save notebook, error: %v", err))
	}
	return &proto.CreateNotebookResponse{
		Notebook: &proto.Notebook{
			Id:          nbEntity.ID,
			Name:        nbEntity.Name,
			Marca:       nbEntity.Marca,
			Modelo:      nbEntity.Modelo,
			NumeroSerie: nbEntity.NumeroSerie,
		},
	}, nil
}
