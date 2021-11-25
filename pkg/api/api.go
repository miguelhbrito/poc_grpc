package api

const (
	AuthorizationCtxKey Context = "authorization"
	UsernameCtxKey      Context = "username"
	TrackingIdCtxKey    Context = "trackingId"
)

type (
	Context  string
	Username string
)

func (u Username) String() string {
	return string(u)
}
