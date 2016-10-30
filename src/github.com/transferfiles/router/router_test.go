package router_test

import (
	"net/http"
	"testing"

	"os"

	"github.com/stretchr/testify/assert"
	"github.com/transferfiles/testutils"
)

func TestHealthCheckHandler(t *testing.T) {
	url := "/healthcheck"
	tr := testutils.NewTestRequest("GET", url)
	response := tr.Run()
	assert.Equal(t, http.StatusOK, response.Code)
	assert.Equal(t, "Transfer file is up and running", response.Body.String())
}

func TestSenderAndReceiver(t *testing.T) {
	os.RemoveAll("/tmp/out.txt")
	url := "/sendFiles"
	tr := testutils.NewTestRequest("POST", url)
	response := tr.Run()
	assert.Equal(t, http.StatusOK, response.Code)
	assert.Equal(t, "Files sent successfully:", response.Body.String())
	assert.True(t, testutils.CheckIfFileExists("/tmp/out.txt"))
}
