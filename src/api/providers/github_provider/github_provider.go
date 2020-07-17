package github_provider

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/shakilbd009/go-microsrvcs/src/api/clients/restclient"
	"github.com/shakilbd009/go-microsrvcs/src/api/domain/github"
)

const (
	headerAuthorization       = "Authorization"
	headerAuthorizationFormat = "token %s"
	urlCreateRepo             = "https://api.github.com/user/repos"
)

func getAuthorizationHeader(token string) string {
	return fmt.Sprintf(headerAuthorizationFormat, token)
}

func CreateRepo(token string, request github.CreateRepoRequest) (*github.CreateRepoResponse, *github.GithubErrorResponse) {

	headers := http.Header{}
	headers.Set(headerAuthorization, getAuthorizationHeader(token))
	response, err := restclient.Post(urlCreateRepo, request, headers)
	if err != nil {
		log.Printf("error when trying to create new repo in github: %s\n", err.Error())
		return nil, &github.GithubErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		}
	}
	bytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, &github.GithubErrorResponse{StatusCode: http.StatusInternalServerError, Message: "invalid response body"}
	}
	defer response.Body.Close()
	if response.StatusCode > 299 {
		var errResponse github.GithubErrorResponse
		if err := json.Unmarshal(bytes, &errResponse); err != nil {
			return nil, &github.GithubErrorResponse{StatusCode: http.StatusInternalServerError, Message: "invalid json response body"}
		}
		errResponse.StatusCode = response.StatusCode
		return nil, &errResponse
	}
	var result github.CreateRepoResponse
	if err := json.Unmarshal(bytes, &result); err != nil {
		log.Printf("error when trying to unmarshal create repo sucessful response: %s\n", err.Error())
		return nil, &github.GithubErrorResponse{StatusCode: http.StatusInternalServerError, Message: "error when trying to unmarshal github create repo response"}
	}
	return &result, nil
}
