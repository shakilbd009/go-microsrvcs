package app

import (
	"github.com/shakilbd009/go-microsrvcs/mvc/controllers"
)

func getUrls() {
	router.GET("/users/:user_id", controllers.GetUser)
}
