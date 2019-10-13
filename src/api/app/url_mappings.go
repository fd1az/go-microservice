package app

import (
	"github.com/fdiaz7/go-microservice/src/api/controllers/polo"
	"github.com/fdiaz7/go-microservice/src/api/controllers/repositories"
)

func mapUrls() {
	router.GET("/marco", polo.Polo)
	router.POST("/repositories", repositories.CreateRepo)
}
