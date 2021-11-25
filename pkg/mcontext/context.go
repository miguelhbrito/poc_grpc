package mcontext

import (
	"context"

	"github.com/poc_grpc/pkg/api"
)

type Context interface {
	context.Context
	Username() api.Username
}

type myContext struct {
	context.Context
}

func NewContext() Context {
	return myContext{Context: context.Background()}
}

func NewFrom(ctx context.Context) Context {
	return myContext{ctx}
}

func WithValue(ctx Context, key interface{}, val interface{}) Context {
	return NewFrom(context.WithValue(ctx, key, val))
}

func (ctx myContext) Username() api.Username {
	user, ok := ctx.Value(api.UsernameCtxKey).(api.Username)
	if !ok && user.String() == "" {
		return ""
	}
	return api.Username(user)
}
