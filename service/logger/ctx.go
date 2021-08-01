package logger

import (
	"context"

	"github.com/sirupsen/logrus"
)

// Get gets logger from a context
func Get(ctx context.Context) *logrus.Logger {
	lg, ok := ctx.Value(ctxKey).(*logrus.Logger)
	if !ok {
		return logrus.StandardLogger()
	}
	return lg
}

// WithLogger returns context with a logger
func WithLogger(ctx context.Context, l *logrus.Logger) context.Context {
	return context.WithValue(ctx, ctxKey, l)
}

type ctxKeyType struct{}

var ctxKey ctxKeyType
