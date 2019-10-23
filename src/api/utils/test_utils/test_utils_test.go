package test_utils

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetMockedContext(t *testing.T) {
	request, err := http.NewRequest(http.MethodGet, "http://localhost:123/something", nil)
	assert.Nil(t, err)

	response := httptest.NewRecorder()
	request.Header = http.Header{"X-Mock": {"true"}}
	c := GetMockContext(request, response)

	assert.EqualValues(t, http.MethodGet, c.Request.Method)
	assert.EqualValues(t, "123", c.Request.URL.Port())
	assert.EqualValues(t, "/something", c.Request.URL.Path)
	assert.EqualValues(t, "http", c.Request.URL.Scheme)
}
