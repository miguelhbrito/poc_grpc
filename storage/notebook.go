package storage

import (
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
		mlog.Error(mctx).Err(err.Err()).Msgf("Error to insert notebook into db")
	}
	return nil
}
