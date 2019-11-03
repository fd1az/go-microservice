package app

import (
	controllers "github.com/fdiaz7/go-microservice/mvc/controllers/user_controller"
)

func mapUrls() {
	router.GET("/users/:user_id", controllers.GetUser)
}
