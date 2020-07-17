package repositories

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/shakilbd009/go-microsrvcs/src/api/clients/restclient"
	"github.com/shakilbd009/go-microsrvcs/src/api/domain/repositories"
	"github.com/shakilbd009/go-microsrvcs/src/api/utils/errors"
	"github.com/shakilbd009/go-microsrvcs/src/api/utils/test_utils"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	restclient.StartMockup()
	os.Exit(m.Run())
}

func TestCreateRepoInvalidJsonRequest(t *testing.T) {

	resp := httptest.NewRecorder()
	request, _ := http.NewRequest(http.MethodPost, "/repositories", strings.NewReader(""))
	c := test_utils.GetMockedContext(request, resp)
	CreateRepo(c)
	assert.EqualValues(t, http.StatusBadRequest, resp.Code)
	apiErr, err := errors.NewApiErrorFromBytes(resp.Body.Bytes())
	assert.Nil(t, err)
	assert.NotNil(t, apiErr)
	assert.EqualValues(t, http.StatusBadRequest, apiErr.Status())
	assert.EqualValues(t, "invalid json body", apiErr.Message())
}

func TestCreateRepoErrorFromGithub(t *testing.T) {

	resp := httptest.NewRecorder()
	request, _ := http.NewRequest(http.MethodPost, "/repositories", strings.NewReader(`{"name": "testing"}`))
	c := test_utils.GetMockedContext(request, resp)
	restclient.FlushMockups()
	restclient.AddMockup(restclient.Mock{
		Url:        "https://api.github.com/user/repos",
		HttpMethod: http.MethodPost,
		Response: &http.Response{
			StatusCode: http.StatusUnauthorized,
			Body:       ioutil.NopCloser(strings.NewReader(`{"message": "Requires authentication","documentation_url": "https://developer.github.com/docs"}`)),
		},
	})
	CreateRepo(c)
	assert.EqualValues(t, http.StatusUnauthorized, resp.Code)
	apiErr, err := errors.NewApiErrorFromBytes(resp.Body.Bytes())
	assert.Nil(t, err)
	assert.NotNil(t, apiErr)
	assert.EqualValues(t, http.StatusUnauthorized, apiErr.Status())
	assert.EqualValues(t, "Requires authentication", apiErr.Message())
}

func TestCreateRepoNoError(t *testing.T) {

	resp := httptest.NewRecorder()
	request, _ := http.NewRequest(http.MethodPost, "/repositories", strings.NewReader(`{"name": "testing"}`))
	c := test_utils.GetMockedContext(request, resp)

	restclient.FlushMockups()
	restclient.AddMockup(restclient.Mock{
		Url:        "https://api.github.com/user/repos",
		HttpMethod: http.MethodPost,
		Response: &http.Response{
			StatusCode: http.StatusCreated,
			Body:       ioutil.NopCloser(strings.NewReader(`{"id": 123}`)),
		},
	})
	CreateRepo(c)
	assert.EqualValues(t, http.StatusCreated, resp.Code)
	var result repositories.CreateRepoResponse
	err := json.Unmarshal(resp.Body.Bytes(), &result)
	assert.Nil(t, err)
	assert.EqualValues(t, 123, result.Id)
	assert.EqualValues(t, "", result.Name)
	assert.EqualValues(t, "", result.Owner)
}
