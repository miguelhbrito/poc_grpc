package entity

import proto "github.com/poc_grpc/pb"

type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginToken struct {
	Token   string `json:"token"`
	ExpTime int64  `json:"expTime"`
}

func (l Login) GenerateEntity(lr *proto.CreateLoginRequest) Login {
	return Login{
		Username: lr.Username,
		Password: lr.Password,
	}
}
