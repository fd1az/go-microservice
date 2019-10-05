package domain

import (
	"errors"
	"fmt"
)

var (
	users = map[int64]*User{
		123: {Id: 123, FirstName: "Facu", LastName: "Diaz", Email: "mail@mail.com"},
	}
)

func GetUser(userID int64) (*User, error) {
	if user := users[userID]; user != nil {
		return user, nil
	}
	return nil, errors.New(fmt.Sprintf("User %v not found", userID))

}
