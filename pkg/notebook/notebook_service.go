package notebook

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

type NotebookService struct {
	proto.UnimplementedNotebookServiceServer
	Manager Notebook
}

func (n NotebookService) CreateNotebook(ctx context.Context, req *proto.CreateNotebookRequest) (*proto.CreateNotebookResponse, error) {
	mctx := mcontext.NewFrom(ctx)
	mlog.Info(mctx).Msg("Received request to create a notebook")
	nbEntity := entity.Notebook{}.GenerateEntity(req)
	nbResponse, err := n.Manager.Create(mctx, nbEntity)
	if err != nil {
		return nil, status.Error(codes.Internal, fmt.Sprintf("Fail to save notebook, error: %v", err))
	}
	return &proto.CreateNotebookResponse{
		Notebook: &proto.Notebook{
			Id:          nbResponse.ID,
			Name:        nbResponse.Name,
			Marca:       nbResponse.Marca,
			Modelo:      nbResponse.Modelo,
			NumeroSerie: nbResponse.NumeroSerie,
		},
	}, nil
}

func (n NotebookService) GetNotebook(ctx context.Context, req *proto.GetNotebookRequest) (*proto.GetNotebookResponse, error) {
	mctx := mcontext.NewFrom(ctx)
	mlog.Info(mctx).Msg("Received request to get a notebook by id")
	nb, err := n.Manager.GetById(mctx, req.Id)
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
	nbs, err := n.Manager.List(mctx)
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
	err := n.Manager.Delete(mctx, req.Id)
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
	err := n.Manager.Update(mctx, *nbEntity)
	if err != nil {
		return nil, status.Error(codes.Internal, fmt.Sprintf("Fail to update notebook, error: %v", err))
	}
	nbResponse, err := n.Manager.GetById(mctx, nbEntity.ID)
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
