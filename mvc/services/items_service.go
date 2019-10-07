package services

import (
	"net/http"

	"github.com/fdiaz7/go-mricroservice/mvc/domain"
	"github.com/fdiaz7/go-mricroservice/mvc/utils"
)

type itemsService struct {
}

var (
	ItemsService itemsService
)

func (s *itemsService) GetItem(itemId string) (*domain.User, *utils.ApplicationError) {
	return nil, &utils.ApplicationError{
		Message:    "Implement me",
		StatusCode: http.StatusInternalServerError,
	}
}
