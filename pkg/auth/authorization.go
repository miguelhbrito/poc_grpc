package auth

import (
	"golang.org/x/crypto/bcrypt"
)

type Authorization struct{}

func NewManager() Auth {
	return Authorization{}
}

func (a Authorization) CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func (a Authorization) GenerateHashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}