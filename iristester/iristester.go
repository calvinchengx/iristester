package iristester

import (
	"net/http"
	"testing"

	"github.com/gavv/httpexpect"
	"github.com/valyala/fasthttp"
)

func IrisTester(t *testing.T, handler fasthttp.RequestHandler, baseUrl string) *httpexpect.Expect {

	return httpexpect.WithConfig(httpexpect.Config{
		BaseURL: baseUrl,
		Client: &http.Client{
			Transport: httpexpect.NewFastBinder(handler),
			Jar:       httpexpect.NewJar(),
		},
		Reporter: httpexpect.NewAssertReporter(t),
		Printers: []httpexpect.Printer{
			httpexpect.NewDebugPrinter(t, true),
		},
	})
}
