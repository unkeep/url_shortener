package database

import (
	"context"
	"sync"
	"testing"

	"github.com/jackc/pgx/v4"
	"github.com/kelseyhightower/envconfig"
	. "github.com/onsi/ginkgo"
	"github.com/onsi/ginkgo/reporters"
	. "github.com/onsi/gomega"
)

func TestDatabase(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping database tests")
		return
	}

	RegisterFailHandler(Fail)
	junitReporter := reporters.NewJUnitReporter("test.xml")
	RunSpecsWithDefaultAndCustomReporters(t, "database tests", []Reporter{junitReporter})
}

var (
	once sync.Once
	conn *pgx.Conn
)

func getDBConn() *pgx.Conn {
	once.Do(func() {
		var cfg Config
		if err := envconfig.Process("", &cfg); err != nil {
			panic(err)
		}

		var err error
		conn, err = Open(context.Background(), cfg)
		if err != nil {
			panic(err)
		}
	})

	return conn
}
