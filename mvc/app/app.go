package app

import (
	"net/http"

	controllers "github.com/fdiaz7/go-mricroservice/mvc/controllers/user_controller"
)

func StartApp() {
	http.HandleFunc("/users", controllers.GetUser)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
