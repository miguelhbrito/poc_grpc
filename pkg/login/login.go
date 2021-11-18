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
	errPasswordHash        = errors.New("Error to generate password hash")
)

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

type Login interface {
	Create(mctx mcontext.Context, l entity.Login) error
	Login(mctx mcontext.Context, l entity.Login) (entity.LoginToken, error)
}
