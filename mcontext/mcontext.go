package mcontext

import (
	"context"

	"github.com/poc_grpc/models"
)

type Context interface {
	context.Context
	Username() models.Username
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

func (ctx myContext) Username() models.Username {
	user, ok := ctx.Value(models.UsernameCtxKey).(models.Username)
	if !ok && user.String() == "" {
		return ""
	}
	return models.Username(user)
}
