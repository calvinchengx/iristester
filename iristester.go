package iristester

import (
	"net/http"
	"testing"

	"github.com/gavv/httpexpect"
	"github.com/valyala/fasthttp"
)

// IrisTester is a wrapper to help us run end-to-end tests easily
// testVerbosity 2, 1, others (0)
func IrisTester(t *testing.T, handler fasthttp.RequestHandler, baseURL string, testVerbosity int) *httpexpect.Expect {

	var printers []httpexpect.Printer
	if testVerbosity == 2 {
		printers = []httpexpect.Printer{
			httpexpect.NewDebugPrinter(t, true),
		}
	} else if testVerbosity == 1 {
		printers = []httpexpect.Printer{
			httpexpect.NewDebugPrinter(t, false),
		}
	} else {
		printers = nil
	}

	return httpexpect.WithConfig(httpexpect.Config{
		BaseURL: baseURL,
		Client: &http.Client{
			Transport: httpexpect.NewFastBinder(handler),
			Jar:       httpexpect.NewJar(),
		},
		Reporter: httpexpect.NewAssertReporter(t),
		Printers: printers,
	})
}
