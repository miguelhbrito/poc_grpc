package api

const (
	AuthorizationCtxKey Context = "x-authorization"
	UsernameCtxKey      Context = "username"
)

type (
	Context  string
	Username string
)

func (u Username) String() string {
	return string(u)
}
