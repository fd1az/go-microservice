package domain

import (
	"fmt"
	"log"
	"net/http"

	"github.com/fdiaz7/go-mricroservice/mvc/utils"
)

var (
	users = map[int64]*User{
		123: {Id: 123, FirstName: "Facu", LastName: "Diaz", Email: "mail@mail.com"},
	}
	UserDao usersDaoInterface
)

func init() {
	UserDao = &userDao{}
}

type usersDaoInterface interface {
	GetUser(int64) (*User, *utils.ApplicationError)
}

type userDao struct{}

//GetUser FROM MOCKING DB
func (u *userDao) GetUser(userID int64) (*User, *utils.ApplicationError) {
	log.Println("we're accessing the database")
	if user := users[userID]; user != nil {
		return user, nil
	}
	return nil, &utils.ApplicationError{
		Message:    fmt.Sprintf("user %v was not found", userID),
		StatusCode: http.StatusNotFound,
		Code:       "not_found",
	}

}
