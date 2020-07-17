package app

import "github.com/gin-gonic/gin"

var router *gin.Engine

func StartApp() {
	router = gin.Default()

	getUrls()
	if err := router.Run(":8081"); err != nil {
		panic(err)
	}
}
