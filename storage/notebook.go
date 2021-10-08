package storage

import (
	"database/sql"

	"github.com/poc_grpc/api/entity"
	dbconnect "github.com/poc_grpc/db_connect"
	"github.com/poc_grpc/mcontext"
	"github.com/poc_grpc/mlog"
)

func SaveNotebook(mctx mcontext.Context, nb entity.Notebook) error {
	db := dbconnect.InitDB()
	defer db.Close()
	sqlStatement := `INSERT INTO notebook VALUES ($1, $2, $3, $4, $5)`
	err := db.QueryRow(sqlStatement, nb.ID, nb.Name, nb.Marca, nb.Modelo, nb.NumeroSerie)
	if err != nil {
		mlog.Error(mctx).Err(err.Err()).Msgf("Error to insert notebook into db %v", err)
	}
	return nil
}

func GetByIdNotebook(mctx mcontext.Context, id string) (entity.Notebook, error) {
	db := dbconnect.InitDB()
	defer db.Close()
	var nb entity.Notebook
	sqlStatement := `SELECT id, name, marca, modelo, numero_serie FROM notebook WHERE id = $1`
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

func ListNotebooks(mctx mcontext.Context) ([]entity.Notebook, error) {
	db := dbconnect.InitDB()
	defer db.Close()
	var nbs []entity.Notebook
	sqlStatement := `SELECT id, name, marca, modelo, numero_serie FROM notebook`
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
