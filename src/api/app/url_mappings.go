package app

import (
	"github.com/shakilbd009/go-microsrvcs/src/api/controllers/polo"
	"github.com/shakilbd009/go-microsrvcs/src/api/controllers/repositories"
)

func mapUrls() {
	router.GET("/marco", polo.Polo)
	router.POST("/repository", repositories.CreateRepo)
	router.POST("/repositories", repositories.CreateRepos)
}
