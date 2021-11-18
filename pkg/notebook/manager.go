package notebook

import (
	"github.com/poc_grpc/pkg/api/entity"
	"github.com/poc_grpc/pkg/mcontext"
	"github.com/poc_grpc/pkg/storage"
)

type manager struct {
	notebookStorage storage.Notebook
}

func NewManager(notebookStorage storage.Notebook) Notebook {
	return manager{
		notebookStorage: notebookStorage,
	}
}

func (m manager) Create(mctx mcontext.Context, nb entity.Notebook) (entity.Notebook, error) {
	return nb, m.notebookStorage.SaveNotebook(mctx, nb)
}

func (m manager) GetById(mctx mcontext.Context, id string) (entity.Notebook, error) {
	return m.notebookStorage.GetByIdNotebook(mctx, id)
}

func (m manager) List(mctx mcontext.Context) ([]entity.Notebook, error) {
	return m.notebookStorage.ListNotebooks(mctx)
}

func (m manager) Delete(mctx mcontext.Context, id string) error {
	return m.notebookStorage.DeleteNotebook(mctx, id)
}

func (m manager) Update(mctx mcontext.Context, nb entity.Notebook) error {
	return m.notebookStorage.UpdateNotebook(mctx, nb)
}
