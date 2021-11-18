package notebook

import (
	"github.com/poc_grpc/pkg/api/entity"
	"github.com/poc_grpc/pkg/mcontext"
)

type Notebook interface {
	Create(mctx mcontext.Context, ac entity.Notebook) (entity.Notebook, error)
	GetById(mctx mcontext.Context, id string) (entity.Notebook, error)
	List(mctx mcontext.Context) ([]entity.Notebook, error)
	Delete(mctx mcontext.Context, id string) error
	Update(mctx mcontext.Context, nb entity.Notebook) error
}
