package repositories

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/shakilbd009/go-microsrvcs/src/api/domain/repositories"
	"github.com/shakilbd009/go-microsrvcs/src/api/services"
	"github.com/shakilbd009/go-microsrvcs/src/api/utils/errors"
	"github.com/shakilbd009/go-microsrvcs/src/api/utils/test_utils"
	"github.com/stretchr/testify/assert"
)

var (
	funcCreateRepo  func(clientId string, request repositories.CreateRepoRequest) (*repositories.CreateRepoResponse, errors.ApiError)
	funcCreateRepos func(requests []repositories.CreateRepoRequest) (repositories.CreateReposResponse, errors.ApiError)
)

type repoServiceMock struct {
}

func (s *repoServiceMock) CreateRepo(clientId string, request repositories.CreateRepoRequest) (*repositories.CreateRepoResponse, errors.ApiError) {
	return funcCreateRepo(clientId, request)
}

func (s *repoServiceMock) CreateRepos(requests []repositories.CreateRepoRequest) (repositories.CreateReposResponse, errors.ApiError) {
	return funcCreateRepos(requests)
}

func TestCreateRepoNoErrorMockingTheEntireService(t *testing.T) {

	services.RepositoryService = &repoServiceMock{}
	funcCreateRepo = func(clientId string, request repositories.CreateRepoRequest) (*repositories.CreateRepoResponse, errors.ApiError) {
		return &repositories.CreateRepoResponse{
			Id:    321,
			Name:  "mocked service",
			Owner: "golang",
		}, nil
	}
	resp := httptest.NewRecorder()
	request, _ := http.NewRequest(http.MethodPost, "/repositories", strings.NewReader(`{"name": "testing"}`))
	c := test_utils.GetMockedContext(request, resp)
	CreateRepo(c)
	assert.EqualValues(t, http.StatusCreated, resp.Code)
	var result repositories.CreateRepoResponse
	err := json.Unmarshal(resp.Body.Bytes(), &result)
	assert.Nil(t, err)
	assert.EqualValues(t, int64(321), result.Id)
	assert.EqualValues(t, "mocked service", result.Name)
	assert.EqualValues(t, "golang", result.Owner)
}

func TestCreateRepoErrorFromGithubMockingTheEntireService(t *testing.T) {

	services.RepositoryService = &repoServiceMock{}
	funcCreateRepo = func(clientId string, request repositories.CreateRepoRequest) (*repositories.CreateRepoResponse, errors.ApiError) {
		return nil, errors.NewBadRequestError("invalid repository name")
	}
	resp := httptest.NewRecorder()
	request, _ := http.NewRequest(http.MethodPost, "/repositories", strings.NewReader(`{"name": "testing"}`))
	c := test_utils.GetMockedContext(request, resp)
	CreateRepo(c)
	assert.EqualValues(t, http.StatusBadRequest, resp.Code)
	apiErr, err := errors.NewApiErrorFromBytes(resp.Body.Bytes())
	assert.Nil(t, err)
	assert.NotNil(t, apiErr)
	assert.EqualValues(t, http.StatusBadRequest, apiErr.Status())
	assert.EqualValues(t, "invalid repository name", apiErr.Message())
}
