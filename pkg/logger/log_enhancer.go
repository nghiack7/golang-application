package logger

import (
	"context"

	"go.uber.org/zap"

	"github.com/nghiack7/golang-application/pkg/xcontext"
)

// Enhance log with additional infos
// Usage: add freely to any existing logger call
// ie: logger.EnhanceWith(ctx).Debugw("a log")
func EnhanceWith(ctx context.Context) LoggerInterface {
	if sg == nil {
		return zl
	}
	return enhanceLogger(sg, ctx, false)
}

// Returns a LoggerInterface with caller info: FuncName@FileName:Line
// This will increase approximately a little bit time each call for logging
func EnhanceWithCallerInfo(ctx context.Context) LoggerInterface {
	if sg == nil {
		return zl
	}
	return enhanceLogger(sg, ctx, true)
}

func enhanceLogger(logger *zap.SugaredLogger, ctx context.Context, withCaller bool) LoggerInterface {
	if ctx == nil {
		return logger
	}

	for _, i := range xcontext.AllKeys {
		val := ctx.Value(i.String())
		if val == nil {
			continue
		}
		if value, ok := val.(string); ok {
			logger = logger.With(zap.String(i.String(), value))
		}
	}

	return &LogInstance{logger, withCaller}
}
