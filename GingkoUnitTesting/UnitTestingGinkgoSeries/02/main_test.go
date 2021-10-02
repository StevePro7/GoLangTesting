package main_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/ghttp"
	"io"
	"io/ioutil"
	"net/http"
	"testing"

	. "unittesting"
)

func TestUnitTestingGinkgoSeries(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "TestingDemo Suite")
}

var _ = Describe("Server", func() {

	var server *ghttp.Server
	msg := "Hi there, the end point is :"

	BeforeEach(func() {
		// start a test http server
		server = ghttp.NewServer()
	})

	AfterEach(func() {
		server.Close()
	})

	Context("When get request is sent to empty path", func() {

		BeforeEach(func() {
			// Add your handler which has to be called for a given path
			// If there are multiple redirects then append all the handlers
			server.AppendHandlers(Handler)
		})

		It("Returns the empty path", func() {
			resp, err := http.Get(server.URL() + "/")
			Expect(err).ShouldNot(HaveOccurred())
			Expect(resp.StatusCode).Should(Equal(http.StatusOK))

			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return
			}

			defer func(Body io.ReadCloser) {
				_ = Body.Close()
			}(resp.Body)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(string(body)).To(Equal(msg + "!"))
		})
	})

})
