package router_test

import (
	"net/http"

	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/transferfiles/router"
	"github.com/urfave/negroni"
)

func TestIndexHandler(t *testing.T) {
	url := "/healthcheck"
	req, _ := http.NewRequest("GET", url, nil)
	w := httptest.NewRecorder()

	n := negroni.New()
	n.UseHandler(router.GetRouter())
	n.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "Transfer file is up and running", w.Body.String())
}
