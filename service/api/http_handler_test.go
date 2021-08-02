package api

import (
	"fmt"
	"net/http"
	"net/http/httptest"

	"url_shortener/service/domain"

	"github.com/gojuno/minimock/v3"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("httpHandler Test", func() {
	var (
		mc             *minimock.Controller
		urlsDomainMock *domain.ShortURLsMock
		h              httpHandler
		cfg            Config
		respRec        *httptest.ResponseRecorder
		req            *http.Request
	)
	BeforeEach(func() {
		mc = minimock.NewController(GinkgoT())
		urlsDomainMock = domain.NewShortURLsMock(mc)
		cfg.DocsPath = "/api/swagger/"
		h = NewHandler(cfg, urlsDomainMock)
		respRec = httptest.NewRecorder()
	})
	AfterEach(func() {
		mc.Finish()
	})

	Context("Redirection", func() {
		BeforeEach(func() {
			req = httptest.NewRequest(http.MethodPost, "/qwe", nil)
		})
		It("should GetOriginURLByShortPath", func() {
			urlsDomainMock.GetOriginURLByShortPathMock.Expect(req.Context(), "/qwe").Return("", fmt.Errorf("test"))
			h.ServeHTTP(respRec, req)
		})
		When("domain fails", func() {
			BeforeEach(func() {
				urlsDomainMock.GetOriginURLByShortPathMock.Return("", fmt.Errorf("test"))
			})
			It("should fail with internal status", func() {
				h.ServeHTTP(respRec, req)
				Ω(respRec.Code).Should(Equal(http.StatusInternalServerError))
			})
		})
		When("domain succeeds", func() {
			const originURLURL = "http://example.com/jkasjdkajsdksjd"
			BeforeEach(func() {
				urlsDomainMock.GetOriginURLByShortPathMock.Return(originURLURL, nil)
			})
			It("should return 301 with the orogin URL in a Location heder", func() {
				h.ServeHTTP(respRec, req)
				Ω(respRec.Code).Should(Equal(http.StatusMovedPermanently))
				Ω(respRec.Header().Get("Location")).Should(Equal(originURLURL))
			})
		})
	})

	Context("Create short URL", func() {
		BeforeEach(func() {
			req = httptest.NewRequest(http.MethodPost, "/api/short_url?origin_url=qwe", nil)
		})
		It("should create URL", func() {
			urlsDomainMock.CreateMock.Expect(req.Context(), "qwe").Return("", fmt.Errorf("test"))
			h.ServeHTTP(respRec, req)
		})
		When("domain fails", func() {
			BeforeEach(func() {
				urlsDomainMock.CreateMock.Return("", fmt.Errorf("test"))
			})
			It("should fail with internal status", func() {
				h.ServeHTTP(respRec, req)
				Ω(respRec.Code).Should(Equal(http.StatusInternalServerError))
			})
		})
		When("domain succeeds", func() {
			const shortURL = "http://example.com/qwe"
			BeforeEach(func() {
				urlsDomainMock.CreateMock.Return(shortURL, nil)
			})
			It("should return OK with the short URL in a body", func() {
				h.ServeHTTP(respRec, req)
				Ω(respRec.Code).Should(Equal(http.StatusOK))
				Ω(respRec.Body.String()).Should(Equal(shortURL))
			})
		})
	})

	// TODO: cover other endpoints
})
