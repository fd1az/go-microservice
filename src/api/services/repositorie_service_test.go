package services

import (
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/fdiaz7/go-microservice/src/api/clients/restclient"
	"github.com/fdiaz7/go-microservice/src/api/domain/repositories"
	"github.com/fdiaz7/go-microservice/src/api/utils/errors"
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
func TestCreateRepoConcurrentInvalidRequest(t *testing.T) {
	request := repositories.CreateRepoRequest{}
	output := make(chan repositories.CreateRepositoriesResult)

	service := reposService{}
	go service.createRepoConcurrent(request, output)
	result := <-output
	assert.NotNil(t, result)
	assert.Nil(t, result.Response)
	assert.NotNil(t, result.Error)
	assert.EqualValues(t, http.StatusBadRequest, result.Error.Status())
	assert.EqualValues(t, "invalid repository name", result.Error.Message())
}

func TestCreateRepoConcurrentErrorFromGithub(t *testing.T) {
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
	output := make(chan repositories.CreateRepositoriesResult)

	service := reposService{}
	go service.createRepoConcurrent(request, output)
	result := <-output
	assert.NotNil(t, result)
	assert.Nil(t, result.Response)
	assert.NotNil(t, result.Error)
	assert.EqualValues(t, http.StatusUnauthorized, result.Error.Status())
	assert.EqualValues(t, "Requieres authentication", result.Error.Message())
}

func TestCreateRepoConcurrentNotErrorFromGithub(t *testing.T) {
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
	output := make(chan repositories.CreateRepositoriesResult)

	service := reposService{}
	go service.createRepoConcurrent(request, output)
	result := <-output
	assert.Nil(t, result.Error)
	assert.NotNil(t, result)
	assert.EqualValues(t, 123, result.Response.Id)
	assert.EqualValues(t, "", result.Response.Name)
	assert.EqualValues(t, "", result.Response.Owner)
}

func TestHandleRepoResults(t *testing.T) {
	input := make(chan repositories.CreateRepositoriesResult)
	output := make(chan repositories.CreateReposResponse)
	var wg sync.WaitGroup

	service := reposService{}
	go service.handleRepoResults(&wg, input, output)
	wg.Add(1)
	go func() {
		input <- repositories.CreateRepositoriesResult{
			Error: errors.NewBadRequestError("invalid repository name"),
		}
	}()

	wg.Wait()
	close(input)
	result := <-output
	assert.NotNil(t, result)
	assert.NotNil(t, result.Results[0].Error)
	assert.EqualValues(t, 0, result.StatusCode)
	assert.EqualValues(t, 1, len(result.Results))
	assert.EqualValues(t, http.StatusBadRequest, result.Results[0].Error.Status())
	assert.EqualValues(t, "invalid repository name", result.Results[0].Error.Message())
}

func TestCreatesReposInvalidRequest(t *testing.T) {
	requests := []repositories.CreateRepoRequest{
		{},
		{Name: "  "},
	}
	result, err := RepositoryService.CreateRepos(requests)
	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.EqualValues(t, 2, len(result.Results))
	assert.EqualValues(t, http.StatusBadRequest, result.StatusCode)

	assert.Nil(t, result.Results[0].Response)
	assert.EqualValues(t, http.StatusBadRequest, result.Results[0].Error.Status())
	assert.EqualValues(t, "invalid repository name", result.Results[0].Error.Message())

	assert.Nil(t, result.Results[0].Response)
	assert.EqualValues(t, http.StatusBadRequest, result.Results[1].Error.Status())
	assert.EqualValues(t, "invalid repository name", result.Results[1].Error.Message())
}

func TestCreatesReposOneSuccesOneFaild(t *testing.T) {
	restclient.FlushMockups()
	restclient.AddMockup(restclient.Mock{
		Url:        "https://api.github.com/user/repos",
		HttpMethod: http.MethodPost,
		Response: &http.Response{
			StatusCode: http.StatusCreated,
			Body:       ioutil.NopCloser(strings.NewReader(`{"id":123}`)),
		},
	})

	requests := []repositories.CreateRepoRequest{
		{},
		{Name: "testing"},
	}
	result, err := RepositoryService.CreateRepos(requests)
	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.EqualValues(t, http.StatusPartialContent, result.StatusCode)
	assert.EqualValues(t, 2, len(result.Results))

	for _, result := range result.Results {
		if result.Error != nil {
			assert.EqualValues(t, http.StatusBadRequest, result.Error.Status())
			assert.EqualValues(t, "invalid repository name", result.Error.Message())
			continue
		}

		assert.EqualValues(t, 123, result.Response.Id)
		assert.EqualValues(t, "", result.Response.Owner)
		assert.EqualValues(t, "", result.Response.Name)
	}
}

func TestCreatesReposAllSuccess(t *testing.T) {
	restclient.FlushMockups()
	restclient.AddMockup(restclient.Mock{
		Url:        "https://api.github.com/user/repos",
		HttpMethod: http.MethodPost,
		Response: &http.Response{
			StatusCode: http.StatusPartialContent,
			Body:       ioutil.NopCloser(strings.NewReader(`{"id":123}`)),
		},
	})

	requests := []repositories.CreateRepoRequest{
		{Name: "testing2"},
		{Name: "testing"},
	}
	result, err := RepositoryService.CreateRepos(requests)
	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.EqualValues(t, http.StatusPartialContent, result.StatusCode)
	assert.EqualValues(t, 2, len(result.Results))

	assert.EqualValues(t, 123, result.Results[0].Response.Id)
	assert.EqualValues(t, "", result.Results[0].Response.Owner)
	assert.EqualValues(t, "", result.Results[0].Response.Name)

}
