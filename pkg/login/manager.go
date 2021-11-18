package login

import (
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/poc_grpc/pkg/api/entity"
	"github.com/poc_grpc/pkg/auth"
	"github.com/poc_grpc/pkg/mcontext"
	"github.com/poc_grpc/pkg/storage"
)

type Manager struct {
	loginStorage storage.Login
	auth         auth.Auth
}

func NewManager(loginStorage storage.Login, auth auth.Auth) Login {
	return Manager{
		loginStorage: loginStorage,
		auth:         auth,
	}
}

func (m Manager) Create(mctx mcontext.Context, l entity.Login) error {
	newPassword, err := m.auth.GenerateHashPassword(l.Password)
	if err != nil {
		return errPasswordHash
	}
	l.Password = newPassword
	err = m.loginStorage.SaveLogin(mctx, l)
	if err != nil {
		return err
	}
	return nil
}

func (m Manager) Login(mctx mcontext.Context, l entity.Login) (entity.LoginToken, error) {
	lr, err := m.loginStorage.GetByIdLogin(mctx, l.Username)
	if err != nil {
		return entity.LoginToken{}, err
	}

	//Checking input secretHash with secretHash from database
	check := m.auth.CheckPasswordHash(l.Password, lr.Password)
	if !check {
		return entity.LoginToken{}, errUserOrPassIncorrect
	}

	//Generation jwt token
	expirationTime := time.Now().Add(1 * time.Hour)
	claims := &Claims{
		Username: l.Password,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	//Signing jwt token with our key
	tokenString, err := token.SignedString(JwtKey)
	if err != nil {
		return entity.LoginToken{}, err
	}

	tokenResponse := entity.LoginToken{
		Token:   tokenString,
		ExpTime: expirationTime.Unix(),
	}

	return tokenResponse, err
}
