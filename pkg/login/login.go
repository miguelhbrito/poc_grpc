package login

import (
	"errors"

	"github.com/golang-jwt/jwt"
	"github.com/poc_grpc/pkg/api/entity"
	"github.com/poc_grpc/pkg/mcontext"
)

var (
	JwtKey                 = []byte("my_secret_key")
	errUserOrPassIncorrect = errors.New("Username or Password is incorrect")
)

type Claims struct {
	Cpf string `json:"username"`
	jwt.StandardClaims
}

type Login interface {
	LoginIntoSystem(mctx mcontext.Context, l entity.Login) (entity.Login, error)
}