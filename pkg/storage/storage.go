package storage

import (
	"github.com/poc_grpc/pkg/api/entity"
	"github.com/poc_grpc/pkg/mcontext"
)

type Notebook interface {
	SaveNotebook(mctx mcontext.Context, nb entity.Notebook) error
	GetByIdNotebook(mctx mcontext.Context, id string) (entity.Notebook, error)
	ListNotebooks(mctx mcontext.Context) ([]entity.Notebook, error)
	DeleteNotebook(mctx mcontext.Context, id string) error
	UpdateNotebook(mctx mcontext.Context, nb entity.Notebook) error
}
