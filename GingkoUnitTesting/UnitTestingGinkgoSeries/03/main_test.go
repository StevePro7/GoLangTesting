package main

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/ghttp"
	"net/http"
	"os"
	"testing"
)

func TestUnitTestingGinkgoSeries(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "TestingDemo Suite")
}

var _ = Describe("Client", func() {

	var (
		server     *ghttp.Server
		statusCode int
		body       []byte
		path       string
		addr       string
	)

	BeforeEach(func() {
		// start a test http server
		server = ghttp.NewServer()
	})
	AfterEach(func() {
		server.Close()
	})

	Context("When given empty url", func() {
		BeforeEach(func() {
			addr = ""
		})
		It("Returns the empty path", func() {
			_, err := getResponse(addr)
			Expect(err).Should(HaveOccurred())
		})
	})

	Context("When give unsupported protocol scheme", func() {
		BeforeEach(func() {
			addr = "tcp://loclahoset"
		})
		It("Returns the empty path", func() {
			_, err := getResponse(addr)
			Expect(err).Should(HaveOccurred())
		})
	})

	Context("When get request is sent to empty path", func() {
		BeforeEach(func() {
			statusCode = http.StatusOK
			path = "/"
			body = []byte("Hi there, the end point is :!")
			addr = "http://" + server.Addr() + path
			server.AppendHandlers(
				ghttp.CombineHandlers(
					ghttp.VerifyRequest("GET", path),
					ghttp.RespondWithPtr(&statusCode, &body),
				))
		})
		It("Returns the hello path", func() {
			bdy, err := getResponse(addr)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(bdy).To(Equal(body))
		})
	})

	Context("When get request is sent to read path but therre is no file", func() {
		BeforeEach(func() {
			statusCode = http.StatusInternalServerError
			path = "/read"
			body = []byte("open data.txt: no such fileor directory\r\n")
			addr = "https://" + server.Addr() + path
			server.AppendHandlers(
				ghttp.CombineHandlers(
					ghttp.VerifyRequest("GET", path),
					ghttp.RespondWithPtr(&statusCode, &body),
				))
		})
		It("Returns internal server error", func() {
			_, err := getResponse(addr)
			Expect(err).Should(HaveOccurred())
		})
	})

	Context("When get request is sent to read path but file exists", func() {
		BeforeEach(func() {
			file, err := os.Create("data.txt")
			Expect(err).NotTo(HaveOccurred())
			body = []byte("Hi there!")
			_, err = file.Write(body)
			if err != nil {
				return
			}
			statusCode = http.StatusOK
			path = "/read"
			addr = "http://" + server.Addr() + path
			server.AppendHandlers(
				ghttp.CombineHandlers(
					ghttp.VerifyRequest("GET", path),
					ghttp.RespondWithPtr(&statusCode, &body),
				))
		})

		AfterEach(func() {
			err := os.Remove("data.txt")
			Expect(err).NotTo(HaveOccurred())
		})
		It("Reads data from file successfully", func() {
			bdy, err := getResponse(addr)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(bdy).To(Equal(body))
		})

	})
})
