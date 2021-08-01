package database

import (
	"context"

	"url_shortener/service/logger"

	"github.com/jackc/pgx/v4"
	"github.com/sirupsen/logrus"
)

// pgxLogAdapter - adapts logrus logger for usage with pgx
type pgxLogAdapter struct{}

func (a *pgxLogAdapter) Log(ctx context.Context, level pgx.LogLevel, msg string, data map[string]interface{}) {
	logger.Get(ctx).WithContext(ctx).WithFields(logrus.Fields(data)).Log(a.mapLevel(level), msg)
}

func (a *pgxLogAdapter) mapLevel(level pgx.LogLevel) logrus.Level {
	switch level {
	case pgx.LogLevelTrace:
		return logrus.TraceLevel
	case pgx.LogLevelDebug, pgx.LogLevelInfo:
		return logrus.DebugLevel
	case pgx.LogLevelWarn:
		return logrus.WarnLevel
	case pgx.LogLevelError:
		return logrus.ErrorLevel
	default:
		return logrus.DebugLevel
	}
}
