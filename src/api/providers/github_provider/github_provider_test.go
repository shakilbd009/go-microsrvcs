package github_provider

import (
	"errors"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"testing"

	"github.com/shakilbd009/go-microsrvcs/src/api/clients/restclient"
	"github.com/shakilbd009/go-microsrvcs/src/api/domain/github"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	restclient.StartMockup()
	os.Exit(m.Run())
}

func TestConstants(t *testing.T) {
	assert.EqualValues(t, "Authorization", headerAuthorization)
	assert.EqualValues(t, "token %s", headerAuthorizationFormat)
	assert.EqualValues(t, "https://api.github.com/user/repos", urlCreateRepo)
}
func TestGetAuthorizationHeader(t *testing.T) {

	token := getAuthorizationHeader("abc123")
	assert.EqualValues(t, "token abc123", token)
}

func TestCreateRepoErrorRestClient(t *testing.T) {

	restclient.FlushMockups()
	restclient.AddMockup(restclient.Mock{
		Url:        "https://api.github.com/user/repos",
		HttpMethod: http.MethodPost,
		Err:        errors.New("invalid restclient response"),
	})
	resp, err := CreateRepo("", github.CreateRepoRequest{})
	assert.Nil(t, resp)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.StatusCode)
	assert.EqualValues(t, "invalid restclient response", err.Message)

}

func TestCreateRepoInvalidResponseBody(t *testing.T) {

	restclient.FlushMockups()
	invalidCloser, _ := os.Open("-ash")
	restclient.AddMockup(restclient.Mock{
		Url:        "https://api.github.com/user/repos",
		HttpMethod: http.MethodPost,
		Response: &http.Response{
			StatusCode: http.StatusCreated,
			Body:       invalidCloser,
		},
	})
	resp, err := CreateRepo("", github.CreateRepoRequest{})
	assert.Nil(t, resp)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.StatusCode)
	assert.EqualValues(t, "invalid response body", err.Message)

}

func TestCreateRepoInvalidErrorInterface(t *testing.T) {

	restclient.FlushMockups()
	restclient.AddMockup(restclient.Mock{
		Url:        "https://api.github.com/user/repos",
		HttpMethod: http.MethodPost,
		Response: &http.Response{
			StatusCode: http.StatusUnauthorized,
			Body:       ioutil.NopCloser(strings.NewReader(`{"message": 1}`)),
		},
	})
	resp, err := CreateRepo("", github.CreateRepoRequest{})
	assert.Nil(t, resp)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.StatusCode)
	assert.EqualValues(t, "invalid json response body", err.Message)

}

func TestCreateRepoUnauthorized(t *testing.T) {

	restclient.FlushMockups()
	restclient.AddMockup(restclient.Mock{
		Url:        "https://api.github.com/user/repos",
		HttpMethod: http.MethodPost,
		Response: &http.Response{
			StatusCode: http.StatusUnauthorized,
			Body:       ioutil.NopCloser(strings.NewReader(`{"message": "Requires authentication","documentation_url": "https://developer.github.com/v3/repos/#create"}`)),
		},
	})
	resp, err := CreateRepo("", github.CreateRepoRequest{})
	assert.Nil(t, resp)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusUnauthorized, err.StatusCode)
	assert.EqualValues(t, "Requires authentication", err.Message)
}

func TestCreateRepoInvalidSucessResponse(t *testing.T) {

	restclient.FlushMockups()
	restclient.AddMockup(restclient.Mock{
		Url:        "https://api.github.com/user/repos",
		HttpMethod: http.MethodPost,
		Response: &http.Response{
			StatusCode: http.StatusCreated,
			Body:       ioutil.NopCloser(strings.NewReader(`{"id": "123"}`)),
		},
	})
	resp, err := CreateRepo("", github.CreateRepoRequest{})
	assert.Nil(t, resp)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.StatusCode)
	assert.EqualValues(t, "error when trying to unmarshal github create repo response", err.Message)
}

func TestCreateRepoNoError(t *testing.T) {

	restclient.FlushMockups()
	restclient.AddMockup(restclient.Mock{
		Url:        "https://api.github.com/user/repos",
		HttpMethod: http.MethodPost,
		Response: &http.Response{
			StatusCode: http.StatusCreated,
			Body:       ioutil.NopCloser(strings.NewReader(`{"id": 123,"name": "golang-tutorial","full_name": "shakilfc009/golang-tutorial"}`)),
		},
	})
	resp, err := CreateRepo("", github.CreateRepoRequest{})
	assert.Nil(t, err)
	assert.NotNil(t, resp)
	assert.EqualValues(t, 123, resp.Id)
	assert.EqualValues(t, "golang-tutorial", resp.Name)
	assert.EqualValues(t, "shakilfc009/golang-tutorial", resp.FullName)
}
