package repositories

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/fdiaz7/go-microservice/src/api/clients/restclient"
	"github.com/fdiaz7/go-microservice/src/api/domain/repositories"
	"github.com/fdiaz7/go-microservice/src/api/utils/errors"
	"github.com/fdiaz7/go-microservice/src/api/utils/test_utils"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	restclient.StartMockups()
	os.Exit(m.Run())
}

func TestCreateRepoInvalidJsonRequest(t *testing.T) {
	request, _ := http.NewRequest(http.MethodPost, "/repositories", strings.NewReader(``))
	response := httptest.NewRecorder()
	c := test_utils.GetMockContext(request, response)
	CreateRepo(c)
	assert.EqualValues(t, http.StatusBadRequest, response.Code)
	fmt.Println(response.Body.String())
	apiErr, err := errors.NewApiErrFromBytes(response.Body.Bytes())

	assert.Nil(t, err)
	assert.NotNil(t, apiErr)
	assert.EqualValues(t, http.StatusBadRequest, apiErr.Status())
	assert.EqualValues(t, "invalid json body", apiErr.Message())
}

func TestCreateRepoErrorFromGithub(t *testing.T) {
	request, _ := http.NewRequest(http.MethodPost, "/repositories", strings.NewReader(`{"name":"test"}`))
	response := httptest.NewRecorder()
	c := test_utils.GetMockContext(request, response)
	restclient.AddMockup(restclient.Mock{
		Url:        "https://api.github.com/user/repos",
		HttpMethod: http.MethodPost,
		Response: &http.Response{
			StatusCode: http.StatusUnauthorized,
			Body:       ioutil.NopCloser(strings.NewReader(`{"message":"Requieres authentication", "documentation_url":"https://api.github.com/user/repos"}`)),
		},
	})
	CreateRepo(c)
	assert.EqualValues(t, http.StatusUnauthorized, response.Code)
	fmt.Println(response.Body.String())
	apiErr, err := errors.NewApiErrFromBytes(response.Body.Bytes())

	assert.Nil(t, err)
	assert.NotNil(t, apiErr)
	assert.EqualValues(t, http.StatusUnauthorized, apiErr.Status())
	assert.EqualValues(t, "Requieres authentication", apiErr.Message())
}

func TestCreateRepoNoError(t *testing.T) {
	request, _ := http.NewRequest(http.MethodPost, "/repositories", strings.NewReader(`{"name":"test"}`))
	response := httptest.NewRecorder()
	c := test_utils.GetMockContext(request, response)
	restclient.AddMockup(restclient.Mock{
		Url:        "https://api.github.com/user/repos",
		HttpMethod: http.MethodPost,
		Response: &http.Response{
			StatusCode: http.StatusCreated,
			Body:       ioutil.NopCloser(strings.NewReader(`{"id":123}`)),
		},
	})
	CreateRepo(c)
	assert.EqualValues(t, http.StatusCreated, response.Code)
	fmt.Println(response.Body.String())
	var result repositories.CreateReposResponse
	err := json.Unmarshal(response.Body.Bytes(), &result)
	assert.Nil(t, err)
	assert.EqualValues(t, 123, result.Id)
	assert.EqualValues(t, "", result.Name)
	assert.EqualValues(t, "", result.Owner)
}
