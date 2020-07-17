package repositories

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/shakilbd009/go-microsrvcs/src/api/domain/repositories"
	"github.com/shakilbd009/go-microsrvcs/src/api/services"
	"github.com/shakilbd009/go-microsrvcs/src/api/utils/errors"
)

func CreateRepo(c *gin.Context) {

	var request repositories.CreateRepoRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		apiErr := errors.NewBadRequestError("invalid json body")
		c.JSON(apiErr.Status(), apiErr)
		return
	}
	ClientId := c.GetHeader("X-Client-Id")
	result, err := services.RepositoryService.CreateRepo(ClientId, request)
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusCreated, result)
}

func CreateRepos(c *gin.Context) {

	var request []repositories.CreateRepoRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		apiErr := errors.NewBadRequestError("invalid json body")
		c.JSON(apiErr.Status(), apiErr)
		return
	}
	result, err := services.RepositoryService.CreateRepos(request)
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(result.StatusCode, result)
}
