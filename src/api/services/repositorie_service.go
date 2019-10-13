package services

import (
	"strings"

	"github.com/fdiaz7/go-microservice/src/api/config"
	"github.com/fdiaz7/go-microservice/src/api/domain/github"
	"github.com/fdiaz7/go-microservice/src/api/domain/repositories"
	"github.com/fdiaz7/go-microservice/src/api/providers/github_provider"
	"github.com/fdiaz7/go-microservice/src/api/utils/errors"
)

type reposService struct{}

type reposServiceInterface interface {
	CreateRepo(request repositories.CreateRepoRequest) (*repositories.CreateReposResponse, errors.ApiError)
}

var (
	RepositoryService reposServiceInterface
)

func init() {
	RepositoryService = &reposService{}
}

func (s *reposService) CreateRepo(input repositories.CreateRepoRequest) (*repositories.CreateReposResponse, errors.ApiError) {
	input.Name = strings.TrimSpace(input.Name)
	if input.Name == "" {
		return nil, errors.NewBadRequestError("invalid repository name")
	}
	request := github.CreateRepoRequest{
		Name:        input.Name,
		Description: input.Description,
		Private:     false,
	}
	response, err := github_provider.CreateRepo(config.GetGithubAccesToken(), request)

	if err != nil {
		return nil, errors.NewApiError(err.StatusCode, err.Message)
	}

	result := repositories.CreateReposResponse{
		Id:    response.Id,
		Name:  response.Name,
		Owner: response.Owner.Login,
	}
	return &result, nil

}
