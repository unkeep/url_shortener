package logger

import (
	logrushbgcp "github.com/TV4/logrus-stackdriver-formatter"
	"github.com/sirupsen/logrus"
)

// Setup setups logger
func Setup(c Config) *logrus.Logger {
	l := logrus.StandardLogger()

	lvl, err := logrus.ParseLevel(c.Level)
	if err != nil {
		panic(err)
	}
	l.SetLevel(lvl)

	if c.Format == "json" {
		l.SetFormatter(logrushbgcp.NewFormatter())
	} else {
		l.SetFormatter(&logrus.TextFormatter{})
	}

	return l
}
