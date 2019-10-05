package services

import (
	"github.com/fdiaz7/go-mricroservice/mvc/domain"
	"github.com/fdiaz7/go-mricroservice/mvc/utils"
)

func GetUser(userId int64) (*domain.User, *utils.ApplicationError) {
	return domain.GetUser(userId)
}
