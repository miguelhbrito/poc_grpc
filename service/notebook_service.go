package service

import (
	"context"
	"fmt"

	"github.com/poc_grpc/mcontext"
	"github.com/poc_grpc/mlog"
	"github.com/poc_grpc/models/entity"
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
		return nil, status.Error(codes.Internal, fmt.Sprintf("Fail to save notebook, error: %v", err))
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

func (n NotebookService) GetNotebook(ctx context.Context, req *proto.GetNotebookRequest) (*proto.GetNotebookResponse, error) {
	mctx := mcontext.NewFrom(ctx)
	mlog.Info(mctx).Msg("Received request to get a notebook by id")
	nb, err := storage.GetByIdNotebook(mctx, req.Id)
	if err != nil {
		return nil, status.Error(codes.Internal, fmt.Sprintf("Fail to get notebook, error: %v", err))
	}
	return &proto.GetNotebookResponse{
		Notebook: &proto.Notebook{
			Id:          nb.ID,
			Name:        nb.Name,
			Marca:       nb.Marca,
			Modelo:      nb.Modelo,
			NumeroSerie: nb.NumeroSerie,
		},
	}, nil
}

func (n NotebookService) ListNotebooks(req *proto.ListNotebooksRequest, stream proto.NotebookService_ListNotebooksServer) error {
	mctx := mcontext.NewFrom(stream.Context())
	mlog.Info(mctx).Msg("Received request to list all notebooks stream")
	nbs, err := storage.ListNotebooks(mctx)
	if err != nil {
		return status.Error(codes.Internal, fmt.Sprintf("Fail to get notebook, error: %v", err))
	}
	for _, nb := range nbs {
		result := &proto.ListNotebooksResponse{
			Notebook: &proto.Notebook{
				Id:          nb.ID,
				Name:        nb.Name,
				Marca:       nb.Marca,
				Modelo:      nb.Modelo,
				NumeroSerie: nb.NumeroSerie,
			},
		}
		err := stream.Send(result)
		if err != nil {
			mlog.Error(mctx).Err(err).Msg("error on send stream notebook list")
			return err
		}
	}
	return nil
}

func (n NotebookService) DeleteNotebook(ctx context.Context, req *proto.DeleteNotebookRequest) (*proto.DeleteNotebookResponse, error) {
	mctx := mcontext.NewFrom(ctx)
	mlog.Info(mctx).Msgf("Received request to delete a notebook by id {%s}", req.GetId())
	err := storage.DeleteNotebook(mctx, req.Id)
	if err != nil {
		return nil, status.Error(codes.Internal, fmt.Sprintf("Fail to delete notebook, error: %v", err))
	}
	return &proto.DeleteNotebookResponse{}, nil
}

func (n NotebookService) UpdateNotebook(ctx context.Context, req *proto.UpdateNotebookRequest) (*proto.UpdateNotebookResponse, error) {
	mctx := mcontext.NewFrom(ctx)
	nbEntity := &entity.Notebook{
		ID:          req.Notebook.Id,
		Name:        req.Notebook.Name,
		Marca:       req.Notebook.Marca,
		Modelo:      req.Notebook.Modelo,
		NumeroSerie: req.Notebook.NumeroSerie,
	}
	mlog.Info(mctx).Msgf("Received request to update a notebook by id {%s}")
	err := storage.UpdateNotebook(mctx, *nbEntity)
	if err != nil {
		return nil, status.Error(codes.Internal, fmt.Sprintf("Fail to update notebook, error: %v", err))
	}
	nbResponse, err := storage.GetByIdNotebook(mctx, nbEntity.ID)
	if err != nil {
		return nil, status.Error(codes.Internal, fmt.Sprintf("Fail to get notebook by id, error: %v", err))
	}
	return &proto.UpdateNotebookResponse{
		Notebook: &proto.Notebook{
			Id:          nbResponse.ID,
			Name:        nbResponse.Name,
			Marca:       nbResponse.Marca,
			Modelo:      nbResponse.Modelo,
			NumeroSerie: nbEntity.NumeroSerie,
		},
	}, nil
}
