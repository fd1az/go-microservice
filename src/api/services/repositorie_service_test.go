package services

import (
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/fdiaz7/go-microservice/src/api/clients/restclient"
	"github.com/fdiaz7/go-microservice/src/api/domain/repositories"
)

func TestMain(m *testing.M) {
	restclient.StartMockups()
	os.Exit(m.Run())
}

func TestCreateRepo(t *testing.T) {
	//initilization
	request := repositories.CreateRepoRequest{}

	//execution
	result, err := RepositoryService.CreateRepo(request)

	//validation
	assert.Nil(t, result)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusBadRequest, err.Status())
	assert.EqualValues(t, "invalid repository name", err.Message())
}

func TestCreateRepoErrorFromGithub(t *testing.T) {
	restclient.FlushMockups()
	restclient.AddMockup(restclient.Mock{
		Url:        "https://api.github.com/user/repos",
		HttpMethod: http.MethodPost,
		Response: &http.Response{
			StatusCode: http.StatusUnauthorized,
			Body:       ioutil.NopCloser(strings.NewReader(`{"message":"Requieres authentication", "documentation_url":"https://api.github.com/user/repos"}`)),
		},
	})
	request := repositories.CreateRepoRequest{Name: "testing"}
	//execution
	result, err := RepositoryService.CreateRepo(request)

	//validation
	assert.Nil(t, result)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusUnauthorized, err.Status())
	assert.EqualValues(t, "Requieres authentication", err.Message())
}

func TestCreateRepoNoError(t *testing.T) {
	restclient.FlushMockups()
	restclient.AddMockup(restclient.Mock{
		Url:        "https://api.github.com/user/repos",
		HttpMethod: http.MethodPost,
		Response: &http.Response{
			StatusCode: http.StatusCreated,
			Body:       ioutil.NopCloser(strings.NewReader(`{"id":123}`)),
		},
	})
	request := repositories.CreateRepoRequest{Name: "testing"}
	//execution
	result, err := RepositoryService.CreateRepo(request)

	//validation
	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.EqualValues(t, 123, result.Id)
	assert.EqualValues(t, "", result.Name)
	assert.EqualValues(t, "", result.Owner)
}
