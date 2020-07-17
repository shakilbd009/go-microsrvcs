package services

import (
	"net/http"
	"sync"

	"github.com/shakilbd009/go-microsrvcs/src/api/config"
	"github.com/shakilbd009/go-microsrvcs/src/api/domain/github"
	"github.com/shakilbd009/go-microsrvcs/src/api/domain/repositories"
	"github.com/shakilbd009/go-microsrvcs/src/api/log/optionB"
	"github.com/shakilbd009/go-microsrvcs/src/api/providers/github_provider"
	"github.com/shakilbd009/go-microsrvcs/src/api/utils/errors"
)

type reposService struct {
}
type reposServiceInterface interface {
	CreateRepo(clientId string, request repositories.CreateRepoRequest) (*repositories.CreateRepoResponse, errors.ApiError)
	CreateRepos(requests []repositories.CreateRepoRequest) (repositories.CreateReposResponse, errors.ApiError)
}

var (
	RepositoryService reposServiceInterface
)

func init() {
	RepositoryService = &reposService{}
}

func (s *reposService) CreateRepo(clientId string, input repositories.CreateRepoRequest) (*repositories.CreateRepoResponse, errors.ApiError) {

	if err := input.Validate(); err != nil {
		return nil, err
	}
	request := github.CreateRepoRequest{
		Name:        input.Name,
		Description: input.Description,
		Private:     false,
	}

	optionB.Info("about to send request to external api",
		optionB.Field("client_id", clientId),
		optionB.Field("status", "pendind"))
	response, err := github_provider.CreateRepo(config.GetGithubAccessToken(), request)
	if err != nil {
		optionB.Error("response obtained from external api", err,
			optionB.Field("client_id", clientId),
			optionB.Field("status", "error"))
		return nil, errors.NewApiError(err.StatusCode, err.Message)
	}
	optionB.Info("response obtained from external api",
		optionB.Field("client_id", clientId),
		optionB.Field("status", "success"))
	result := repositories.CreateRepoResponse{
		Id:    response.Id,
		Name:  response.Name,
		Owner: response.Owner.Login,
	}
	return &result, nil
}

func (s *reposService) CreateRepos(requests []repositories.CreateRepoRequest) (repositories.CreateReposResponse, errors.ApiError) {

	input := make(chan repositories.CreateRepositoriesResult)
	output := make(chan repositories.CreateReposResponse)
	defer close(output)
	var wg sync.WaitGroup
	go s.handleRepoResults(&wg, input, output)

	for _, request := range requests {
		wg.Add(1)
		go s.createRepoConcurrent(request, input)
	}
	wg.Wait()
	close(input)
	result := <-output
	successCreation := 0
	for _, req := range result.Results {
		if req.Response != nil {
			successCreation++
		}
	}
	if successCreation == 0 {
		result.StatusCode = result.Results[0].Error.Status()
	} else if successCreation == len(requests) {
		result.StatusCode = http.StatusCreated
	} else {
		result.StatusCode = http.StatusPartialContent
	}
	return result, nil
}

func (s *reposService) handleRepoResults(wg *sync.WaitGroup, input chan repositories.CreateRepositoriesResult, output chan repositories.CreateReposResponse) {

	var results repositories.CreateReposResponse

	for result := range input {
		repoResult := repositories.CreateRepositoriesResult{
			Response: result.Response,
			Error:    result.Error,
		}
		results.Results = append(results.Results, repoResult)
		wg.Done()
	}
	output <- results
}

func (s *reposService) createRepoConcurrent(input repositories.CreateRepoRequest, output chan repositories.CreateRepositoriesResult) {

	if err := input.Validate(); err != nil {
		output <- repositories.CreateRepositoriesResult{
			Error: err,
		}
		return
	}
	result, err := s.CreateRepo("", input)
	if err != nil {
		output <- repositories.CreateRepositoriesResult{
			Error: err,
		}
		return
	}
	output <- repositories.CreateRepositoriesResult{Response: result}
}
