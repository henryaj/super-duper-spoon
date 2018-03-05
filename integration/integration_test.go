package integration_test

import (
	"io/ioutil"
	"net/http"
	"os/exec"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gexec"
)

var session *gexec.Session

var _ = Describe("my tiny server", func() {
	BeforeEach(func() {
		var err error

		command := exec.Command(serverBinaryPath)
		session, err = gexec.Start(command, GinkgoWriter, GinkgoWriter)
		Expect(err).NotTo(HaveOccurred())
		time.Sleep(3 * time.Second)
	})

	AfterEach(func() {
		session.Kill()
	})

	It("lets a user set and get a key", func() {
		makeAPIRequest("set?foo=bar")
		makeAPIRequest("set?baz=bam")

		Expect(makeAPIRequest("get?key=foo")).To(Equal("bar"))
		Expect(makeAPIRequest("get?key=baz")).To(Equal("bam"))
	})
})

func makeAPIRequest(path string) string {
	resp, err := http.Get("http://localhost:4000/" + path)
	Expect(err).NotTo(HaveOccurred())

	result, err := ioutil.ReadAll(resp.Body)
	Expect(err).NotTo(HaveOccurred())

	return string(result)
}
