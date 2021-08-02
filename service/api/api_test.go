package api

import (
	"testing"

	. "github.com/onsi/ginkgo"
	"github.com/onsi/ginkgo/reporters"
	. "github.com/onsi/gomega"
)

func TestAPI(t *testing.T) {
	RegisterFailHandler(Fail)
	junitReporter := reporters.NewJUnitReporter("test.xml")
	RunSpecsWithDefaultAndCustomReporters(t, "API tests", []Reporter{junitReporter})
}
