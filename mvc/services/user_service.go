package services

import (
	"github.com/fdiaz7/go-mricroservice/mvc/domain"
	"github.com/fdiaz7/go-mricroservice/mvc/utils"
)

type usersService struct{}

var (
	//UsersService Exported
	UsersService usersService
)

func (u *usersService) GetUser(userId int64) (*domain.User, *utils.ApplicationError) {
	user, err := domain.UserDao.GetUser(userId)

	if err != nil {
		return nil, err
	}

	return user, nil
}
