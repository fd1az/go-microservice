package app

import (
	"github.com/fdiaz7/go-microservice/oauth-api/src/api/controllers/oauth"
	"github.com/fdiaz7/go-microservice/src/api/controllers/polo"
)

func mapUrls() {
	router.GET("/marco", polo.Polo)
	router.POST("/oauth/access_token", oauth.CreateAccessToken)
	router.POST("/oauth/access_token/:token_id", oauth.GetAccessToken)
}
