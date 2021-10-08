package mlog

import (
	"context"
	"os"

	"github.com/poc_grpc/models"
	"github.com/rs/zerolog"
)

const (
	debug = "debug"
	info  = "info"
	error = "error"
	warn  = "warn"
	panic = "panic"
)

type (
	logInfo struct {
		trackingId string
		username   string
	}
)

func buildLog(ctx context.Context, level string) *zerolog.Event {
	logger := zerolog.New(os.Stdout)

	if ctx != nil {
		li := loadLogInfo(ctx)

		if li.username != "" {
			logger = logger.With().
				Str(string(models.UsernameCtxKey), li.username).Logger()
		}

		if li.trackingId != "" {
			logger = logger.With().
				Str(string(models.TrackingIdCtxKey), li.trackingId).Logger()
		}
	}

	switch level {
	case info:
		return logger.Info()
	case debug:
		return logger.Debug()
	case error:
		return logger.Error()
	case warn:
		return logger.Warn()
	case panic:
		return logger.Panic()
	default:
		return logger.Info()
	}
}

func loadLogInfo(ctx context.Context) logInfo {
	var li logInfo

	if v := ctx.Value(models.UsernameCtxKey); v != nil {
		li.username = v.(models.Username).String()
	}
	if v := ctx.Value(models.TrackingIdCtxKey); v != nil {
		li.trackingId = v.(models.TrackingId).String()
	}
	return li
}

func Info(ctx context.Context) *zerolog.Event {
	return buildLog(ctx, info)
}

func Error(ctx context.Context) *zerolog.Event {
	return buildLog(ctx, error)
}

func Debug(ctx context.Context) *zerolog.Event {
	return buildLog(ctx, debug)
}

func Warn(ctx context.Context) *zerolog.Event {
	return buildLog(ctx, warn)
}

func Panic(ctx context.Context) *zerolog.Event {
	return buildLog(ctx, panic)
}
