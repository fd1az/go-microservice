package github_provider

import (
	"errors"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"testing"

	"github.com/fdiaz7/go-mricroservice/src/clients/restclient"
	"github.com/fdiaz7/go-mricroservice/src/domain/github"

	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	restclient.StoptMockups()
	os.Exit(m.Run())
}

func TestConstans(t *testing.T) {
	assert.EqualValues(t, "Authorization", headerAuthorization)
	assert.EqualValues(t, "token %s", headerAuthorizationFormat)
	assert.EqualValues(t, "https://api.github.com/user/repos", urlCreateRepo)
}

func TestGetAuthorizationHeader(t *testing.T) {
	header := getAuthorizationHeader("abc123")
	assert.EqualValues(t, "token abc123", header)
}

func TestCreateRepoErrorRestclient(t *testing.T) {
	restclient.StartMockups()
	restclient.FlushMockups()
	restclient.AddMockup(restclient.Mock{
		Url:        "https://api.github.com/user/repos",
		HttpMethod: http.MethodPost,
		Err:        errors.New("invalid restclient response2"),
	})
	response, err := CreateRepo("", github.CreateRepoRequest{})
	assert.Nil(t, response)
	assert.NotNil(t, err)
	assert.EqualValues(t, "invalid restclient response2", err.Message)
	restclient.StoptMockups()
}

func TestCreateRepoErrorInvalidResponseBody(t *testing.T) {
	restclient.StartMockups()
	restclient.FlushMockups()
	invalidCloser, _ := os.Open("-sdf3")
	restclient.AddMockup(restclient.Mock{
		Url:        "https://api.github.com/user/repos",
		HttpMethod: http.MethodPost,
		Response: &http.Response{
			StatusCode: http.StatusCreated,
			Body:       invalidCloser,
		},
	})
	response, err := CreateRepo("", github.CreateRepoRequest{})
	assert.Nil(t, response)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.StatusCode)
	assert.EqualValues(t, "Invalid response Body", err.Message)
	restclient.StoptMockups()
}

func TestCreateRepoInvalidErrorInterface(t *testing.T) {
	restclient.StartMockups()
	restclient.FlushMockups()
	restclient.AddMockup(restclient.Mock{
		Url:        "https://api.github.com/user/repos",
		HttpMethod: http.MethodPost,
		Response: &http.Response{
			StatusCode: http.StatusInternalServerError,
			Body:       ioutil.NopCloser(strings.NewReader(`{"message":1}`)),
		},
	})
	response, err := CreateRepo("", github.CreateRepoRequest{})
	assert.Nil(t, response)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.StatusCode)
	assert.EqualValues(t, "Invalid JSON response", err.Message)
	restclient.StoptMockups()
}

func TestCreateRepoUnauthorized(t *testing.T) {
	restclient.StartMockups()
	restclient.FlushMockups()
	restclient.AddMockup(restclient.Mock{
		Url:        "https://api.github.com/user/repos",
		HttpMethod: http.MethodPost,
		Response: &http.Response{
			StatusCode: http.StatusUnauthorized,
			Body:       ioutil.NopCloser(strings.NewReader(`{"message":"Requires authentication","documentaion_url":"https://developer.github.com/v3/repos/#create"}`)),
		},
	})
	response, err := CreateRepo("", github.CreateRepoRequest{})
	assert.Nil(t, response)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusUnauthorized, err.StatusCode)
	assert.EqualValues(t, "Requires authentication", err.Message)
	restclient.StoptMockups()
}

func TestCreateRepoInvalidSuccessResponse(t *testing.T) {
	restclient.StartMockups()
	restclient.FlushMockups()
	restclient.AddMockup(restclient.Mock{
		Url:        "https://api.github.com/user/repos",
		HttpMethod: http.MethodPost,
		Response: &http.Response{
			StatusCode: http.StatusCreated,
			Body:       ioutil.NopCloser(strings.NewReader(`{"id":"123"}`)),
		},
	})
	response, err := CreateRepo("", github.CreateRepoRequest{})
	assert.Nil(t, response)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.StatusCode)
	assert.EqualValues(t, "Error unmarshaling github create repo", err.Message)
	restclient.StoptMockups()
}

func TestCreateRepoNoError(t *testing.T) {
	restclient.StartMockups()
	restclient.FlushMockups()
	restclient.AddMockup(restclient.Mock{
		Url:        "https://api.github.com/user/repos",
		HttpMethod: http.MethodPost,
		Response: &http.Response{
			StatusCode: http.StatusCreated,
			Body: ioutil.NopCloser(strings.NewReader(`{
				"id": 123,
				"name": "golang-tutorial",
				"full_name": "fdiaz7/golang-tutorial"
			  }`)),
		},
	})
	response, err := CreateRepo("", github.CreateRepoRequest{})
	assert.Nil(t, err)
	assert.NotNil(t, response)
	assert.EqualValues(t, 123, response.Id)
	assert.EqualValues(t, "golang-tutorial", response.Name)
	assert.EqualValues(t, "fdiaz7/golang-tutorial", response.FullName)
	restclient.StoptMockups()
}
