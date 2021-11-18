package entity

import proto "github.com/poc_grpc/pb"

type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (l Login) GenerateEntity(lr *proto.CreateLoginRequest) Login {
	return Login{
		Username: lr.Username,
		Password: lr.Password,
	}
}
