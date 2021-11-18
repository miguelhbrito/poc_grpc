package auth

type Auth interface {
	CheckPasswordHash(password, hash string) bool
	GenerateHashPassword(password string) (string, error)
}
