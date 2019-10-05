package services

import "github.com/fdiaz7/go-mricroservice/mvc/domain"

func GetUser(userId int64) (*domain.User, error) {
	return domain.GetUser(userId)
}
