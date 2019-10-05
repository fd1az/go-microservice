package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/fdiaz7/go-mricroservice/mvc/services"
)

func GetUser(res http.ResponseWriter, req *http.Request) {
	userId, err := strconv.ParseInt(req.URL.Query().Get("user_id"), 10, 64)

	if err != nil {
		panic(err)
	}
	user, err := services.GetUser(userId)
	if err != nil {
		res.WriteHeader(http.StatusNotFound)
		res.Write([]byte(err.Error()))
		return
	}
	//return user to client
	jsonValue, _ := json.Marshal(user)
	res.Write(jsonValue)
}
