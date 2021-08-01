package domain

import (
	"testing"

	. "github.com/onsi/ginkgo"
	"github.com/onsi/ginkgo/reporters"
	. "github.com/onsi/gomega"
)

func TestDomain(t *testing.T) {
	RegisterFailHandler(Fail)
	junitReporter := reporters.NewJUnitReporter("test.xml")
	RunSpecsWithDefaultAndCustomReporters(t, "domain tests", []Reporter{junitReporter})
}
