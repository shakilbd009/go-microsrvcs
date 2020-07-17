package repositories

import (
	"strings"

	"github.com/shakilbd009/go-microsrvcs/src/api/utils/errors"
)

type (
	CreateRepoRequest struct {
		Name        string `json:"name"`
		Description string `json:"description"`
	}

	CreateRepoResponse struct {
		Id    int64  `json:"id"`
		Name  string `json:"name"`
		Owner string `json:"owner"`
	}

	CreateReposResponse struct {
		StatusCode int                        `json:"status"`
		Results    []CreateRepositoriesResult `json:"results"`
	}

	CreateRepositoriesResult struct {
		Response *CreateRepoResponse `json:"repo"`
		Error    errors.ApiError     `json:"error"`
	}
)

func (c *CreateRepoRequest) Validate() errors.ApiError {

	c.Name = strings.TrimSpace(c.Name)
	if c.Name == "" {
		return errors.NewBadRequestError("invalid repository name")
	}
	return nil
}
