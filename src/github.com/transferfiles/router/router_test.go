package router_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/transferfiles/testutils"
	"net/http"
	"testing"
)

func TestHealthCheckHandler(t *testing.T) {
	url := "/healthcheck"
	tr := testutils.NewTestRequest("GET", url)
	response := tr.Run()
	assert.Equal(t, http.StatusOK, response.Code)
	assert.Equal(t, "Transfer file is up and running", response.Body.String())
}
