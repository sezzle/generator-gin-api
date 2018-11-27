package gin_test

import (
	"bytes"
	"encoding/json"
	"flag"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/glog"

	"github.com/gin-gonic/gin"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func init() {
	if testing.Verbose() {
		flag.Set("alsologtostderr", "true")
		flag.Set("v", "5")
	}
}

var (
	jsonErr error
)

var _ = Describe("SignupHandler", func() {
	RegisterFailHandler(Fail)

	BeforeEach(func() {
		endpointHeaders = make(http.Header)
		endpointHeaders.Add("X-Real-IP", "74.37.200.161") //Setting a fake IP address for login security tests.
		form = gin.H{}
	})

	JustBeforeEach(func() {
		glog.Error("Fire request")
		response = httptest.NewRecorder()

		if (endpointMethod == "POST") || (endpointMethod == "PUT") || endpointMethod == "PATCH" {
			jsonString, _ := json.Marshal(form)
			var err error
			request, err = http.NewRequest(endpointMethod, endpointURL, bytes.NewReader(jsonString))
			if err != nil {
				glog.Error(err)
			}
			request.Header = endpointHeaders
			request.Header.Add("Content-Type", "application/json")

		} else {
			request, _ = http.NewRequest(endpointMethod, endpointURL, nil)
			request.Header = endpointHeaders
		}

		s.ServeHTTP(response, request)
	})

	//g.GET("/debug/test",)
	Describe("GET /debug/test", func() {
		var responseKeyValue keyValueResp
		BeforeEach(func() {
			endpointMethod = "GET"
			endpointURL = "http://localhost:8000/debug/test"
		})

		Context("On testing the debut route", func() {
			BeforeEach(func() {
				form = gin.H{}
			})

			JustBeforeEach(func() {
				jsonErr = DecodeTestJson(response, &responseKeyValue)
			})

			It("should return an error", func() {
				Ω(response.Code).Should(Equal(http.StatusOK))
				// Ω(jsonErr).ShouldNot(HaveOccurred())
				// Ω(responseKeyValue).ShouldNot(BeEmpty())
			})
		})

		Context("On next context", func() {
			BeforeEach(func() {
				form = gin.H{}
				endpointMethod = "POST"
				endpointURL = "http://localhost:8000/v1/user"
			})

			JustBeforeEach(func() {
				jsonErr = DecodeTestJson(response, &responseKeyValue)
			})

			It("Return status ok", func() {
				Ω(response.Code).Should(Equal(http.StatusOK))
				// Ω(jsonErr).ShouldNot(HaveOccurred())
				// Ω(responseKeyValue).ShouldNot(BeEmpty())
				// Ω(responseKeyValue).Should(HaveKeyWithValue("name", "Test Name"))
			})
		})
	})

})
