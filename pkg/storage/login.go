package storage

import (
	"database/sql"

	dbconnect "github.com/poc_grpc/db_connect"
	"github.com/poc_grpc/pkg/api/entity"
	"github.com/poc_grpc/pkg/mcontext"
	"github.com/poc_grpc/pkg/mlog"
	"github.com/poc_grpc/pkg/observability"
)

type LoginPostgres struct{}

func NewLoginPostgres() Login {
	return LoginPostgres{}
}

func (l LoginPostgres) SaveLogin(mctx mcontext.Context, lr entity.Login) error {
	db := dbconnect.InitDB()
	defer db.Close()
	sqlStatement := `INSERT INTO login VALUES ($1, $2)`
	mctx = observability.CreateSpan(mctx, observability.MySpan{
		ServiceName: "communication database",
		Infos: map[string]string{
			"SQL": sqlStatement,
		},
	})
	defer observability.Finish(mctx)
	_, err := db.Exec(sqlStatement, lr.Username, lr.Password)
	if err != nil {
		mlog.Error(mctx).Err(err).Msgf("Error to insert notebook into db %v", err)
		return err
	}
	return nil
}

func (l LoginPostgres) GetByIdLogin(mctx mcontext.Context, id string) (entity.Login, error) {
	db := dbconnect.InitDB()
	defer db.Close()
	var lr entity.Login
	sqlStatement := `SELECT username, password FROM login WHERE username = $1`
	mctx = observability.CreateSpan(mctx, observability.MySpan{
		ServiceName: "communication database",
		Infos: map[string]string{
			"SQL": sqlStatement,
		},
	})
	defer observability.Finish(mctx)
	result := db.QueryRow(sqlStatement, id)
	err := result.Scan(&lr.Username, &lr.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			mlog.Error(mctx).Err(err).Msgf("Not found notebook with id %s", id)
			return entity.Login{}, err
		}
		mlog.Error(mctx).Err(err).Msgf("Error to get notebook from db, with id %s", id)
		return entity.Login{}, err
	}
	return lr, nil
}
