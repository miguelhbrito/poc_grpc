package api

const (
	AuthorizationCtxKey Context = "authorization"
	UsernameCtxKey      Context = "username"
	TrackingIdCtxKey    Context = "trackingId"
)

type (
	Context    string
	Username   string
	TrackingId string
)

func (u Username) String() string {
	return string(u)
}

func (t TrackingId) String() string {
	return string(t)
}
