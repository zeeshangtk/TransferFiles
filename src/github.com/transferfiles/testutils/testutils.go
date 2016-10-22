package testutils

import (
	"bytes"
	"github.com/urfave/negroni"
	"io"
	"net/http"
	"net/http/httptest"

	"github.com/transferfiles/router"
)

type TestRequest struct {
	Method string
	Url    string
	Body   io.Reader
}

type TestResponse struct {
	Code int           // the HTTP response code from WriteHeader
	Body *bytes.Buffer // if non-nil, the bytes.Buffer to append written data to
}

func NewTestRequest(method, url string) *TestRequest {
	return &TestRequest{
		Method: method,
		Url:    url,
	}
}

func (tr *TestRequest) Run() *TestResponse {
	req, _ := http.NewRequest(tr.Method, tr.Url, tr.Body)
	w := httptest.NewRecorder()

	n := negroni.New()
	n.UseHandler(router.GetRouter())
	n.ServeHTTP(w, req)

	return &TestResponse{
		Code: w.Code,
		Body: w.Body,
	}
}
