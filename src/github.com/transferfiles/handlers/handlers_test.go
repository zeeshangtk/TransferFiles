package handlers_test

import (
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"regexp"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/transferfiles/handlers"
)

const (
	fileaContents = "This is a test file."
	filebContents = "Another test file."
	textaValue    = "foo"
	textbValue    = "bar"
	boundary      = `MyBoundary`
)

const message = `
--MyBoundary
Content-Disposition: form-data; name="filea"; filename="filea.txt"
Content-Type: text/plain

` + fileaContents + `
--MyBoundary
Content-Disposition: form-data; name="fileb"; filename="fileb.txt"
Content-Type: text/plain

` + filebContents + `
--MyBoundary
Content-Disposition: form-data; name="texta"

` + textaValue + `
--MyBoundary
Content-Disposition: form-data; name="textb"

` + textbValue + `
--MyBoundary--
`

func TestUploadHandler(t *testing.T) {
	testBody := regexp.MustCompile("\n").ReplaceAllString(message, "\r\n")
	b := strings.NewReader(testBody)
	r := multipart.NewReader(b, boundary)
	f, _ := r.ReadForm(25)
	w := httptest.NewRecorder()
	request := &http.Request{
		Method:        "POST",
		MultipartForm: f,
	}
	handlers.UploadHandler(w, request)
	assert.Equal(t, "Files uploaded successfully:", w.Body.String())
}
