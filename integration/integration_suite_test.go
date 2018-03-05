package integration_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gexec"
)

func TestRecurse(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Recurse Suite")
}

var serverBinaryPath string

var _ = BeforeSuite(func() {
	var err error
	serverBinaryPath, err = gexec.Build("github.com/henryaj/recurse")
	Expect(err).NotTo(HaveOccurred())
})

var _ = AfterSuite(func() {
	gexec.CleanupBuildArtifacts()
})
