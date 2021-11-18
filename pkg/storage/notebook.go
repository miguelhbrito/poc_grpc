package storage

import (
	"database/sql"

	dbconnect "github.com/poc_grpc/db_connect"
	"github.com/poc_grpc/pkg/api/entity"
	"github.com/poc_grpc/pkg/mcontext"
	"github.com/poc_grpc/pkg/mlog"
	"github.com/poc_grpc/pkg/observability"
)

type NotebookPostgres struct{}

func NewNotebookPostgres() Notebook {
	return NotebookPostgres{}
}

func (n NotebookPostgres) SaveNotebook(mctx mcontext.Context, nb entity.Notebook) error {
	db := dbconnect.InitDB()
	defer db.Close()
	sqlStatement := `INSERT INTO notebook VALUES ($1, $2, $3, $4, $5)`
	mctx = observability.CreateSpan(mctx, observability.MySpan{
		ServiceName: "communication database",
		Infos: map[string]string{
			"SQL": sqlStatement,
		},
	})
	defer observability.Finish(mctx)
	_, err := db.Exec(sqlStatement, nb.ID, nb.Name, nb.Marca, nb.Modelo, nb.NumeroSerie)
	if err != nil {
		mlog.Error(mctx).Err(err).Msgf("Error to insert notebook into db %v", err)
		return err
	}
	return nil
}

func (n NotebookPostgres) GetByIdNotebook(mctx mcontext.Context, id string) (entity.Notebook, error) {
	db := dbconnect.InitDB()
	defer db.Close()
	var nb entity.Notebook
	sqlStatement := `SELECT id, name, marca, modelo, numero_serie FROM notebook WHERE id = $1`
	mctx = observability.CreateSpan(mctx, observability.MySpan{
		ServiceName: "communication database",
		Infos: map[string]string{
			"SQL": sqlStatement,
		},
	})
	defer observability.Finish(mctx)
	result := db.QueryRow(sqlStatement, id)
	err := result.Scan(&nb.ID, &nb.Name, &nb.Marca, &nb.Modelo, &nb.NumeroSerie)
	if err != nil {
		if err == sql.ErrNoRows {
			mlog.Error(mctx).Err(err).Msgf("Not found notebook with id %s", id)
			return entity.Notebook{}, err
		}
		mlog.Error(mctx).Err(err).Msgf("Error to get notebook from db, with id %s", id)
		return entity.Notebook{}, err
	}
	return nb, nil
}

func (n NotebookPostgres) ListNotebooks(mctx mcontext.Context) ([]entity.Notebook, error) {
	db := dbconnect.InitDB()
	defer db.Close()
	var nbs []entity.Notebook
	sqlStatement := `SELECT id, name, marca, modelo, numero_serie FROM notebook`
	mctx = observability.CreateSpan(mctx, observability.MySpan{
		ServiceName: "communication database",
		Infos: map[string]string{
			"SQL": sqlStatement,
		},
	})
	defer observability.Finish(mctx)
	rows, err := db.Query(sqlStatement)
	if err != nil {
		mlog.Error(mctx).Err(err).Msg("Error to get all notebooks from db")
		return nil, err
	}
	for rows.Next() {
		var nb entity.Notebook
		err := rows.Scan(&nb.ID, &nb.Name, &nb.Marca, &nb.Modelo, &nb.NumeroSerie)
		if err != nil {
			mlog.Error(mctx).Err(err).Msgf("Error to extract result from row, err: %s", err)
		}
		nbs = append(nbs, nb)
	}
	return nbs, nil
}

func (n NotebookPostgres) DeleteNotebook(mctx mcontext.Context, id string) error {
	db := dbconnect.InitDB()
	defer db.Close()
	sqlStatement := `DELETE FROM notebook WHERE id=$1`
	mctx = observability.CreateSpan(mctx, observability.MySpan{
		ServiceName: "communication database",
		Infos: map[string]string{
			"SQL": sqlStatement,
		},
	})
	defer observability.Finish(mctx)
	_, err := db.Exec(sqlStatement, id)
	if err != nil {
		mlog.Error(mctx).Err(err).Msgf("Error to delete notebook from db %v", err)
		return err
	}
	return nil
}

func (n NotebookPostgres) UpdateNotebook(mctx mcontext.Context, nb entity.Notebook) error {
	db := dbconnect.InitDB()
	defer db.Close()
	sqlStatement := `UPDATE notebook SET name=$2, marca=$3, modelo=$4, numero_serie=$5 WHERE id=$1`
	mctx = observability.CreateSpan(mctx, observability.MySpan{
		ServiceName: "communication database",
		Infos: map[string]string{
			"SQL": sqlStatement,
		},
	})
	defer observability.Finish(mctx)
	_, err := db.Exec(sqlStatement, nb.ID, nb.Name, nb.Marca, nb.Modelo, nb.NumeroSerie)
	if err != nil {
		mlog.Error(mctx).Err(err).Msgf("Error to update notebook from db %v", err)
		return err
	}
	return nil
}
