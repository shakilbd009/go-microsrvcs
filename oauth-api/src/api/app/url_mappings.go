package app

import (
	"github.com/shakilbd009/go-microsrvcs/oauth-api/src/api/controller/oauth"
	"github.com/shakilbd009/go-microsrvcs/src/api/controllers/polo"
)

func getUrls() {

	router.GET("/marco", polo.Polo)

	router.POST("/oauth/access_token", oauth.CreateAccessToken)
	router.GET("oauth/access_token/:token_id", oauth.GetAccessToken)
}
