package services

import (
	"net/http"

	"github.com/shakilbd009/go-microsrvcs/mvc/domain"
	"github.com/shakilbd009/go-microsrvcs/mvc/utils"
)

type itemsService struct {
}

var (
	ItemsService itemsService
)

func (s *itemsService) GetItem(itemId string) (*domain.Item, *utils.ApplicationError) {
	return nil, &utils.ApplicationError{
		Message:    "implement me",
		StatusCode: http.StatusInternalServerError,
	}
}
