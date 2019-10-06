package domain

import (
	"fmt"
	"net/http"

	"github.com/fdiaz7/go-mricroservice/mvc/utils"
)

var (
	users = map[int64]*User{
		123: {Id: 123, FirstName: "Facu", LastName: "Diaz", Email: "mail@mail.com"},
	}
)

//GetUser FROM MOCKING DB
func GetUser(userID int64) (*User, *utils.ApplicationError) {
	if user := users[userID]; user != nil {
		return user, nil
	}
	return nil, &utils.ApplicationError{
		Message:    fmt.Sprintf("user %v was not found", userID),
		StatusCode: http.StatusNotFound,
		Code:       "not_found",
	}

}
