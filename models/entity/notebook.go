package entity

import (
	"github.com/google/uuid"
	proto "github.com/poc_grpc/pb"
)

type Notebook struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Marca       string `json:"marca"`
	Modelo      string `json:"modelo"`
	NumeroSerie int64  `json:"numeroSerie"`
}

func GrpcToEntity(nb *proto.CreateNotebookRequest) Notebook {
	return Notebook{
		ID:          uuid.New().String(),
		Name:        nb.Name,
		Marca:       nb.Marca,
		Modelo:      nb.Marca,
		NumeroSerie: nb.NumeroSerie,
	}
}
