package services

import (
	"github.com/shakilbd009/go-microsrvcs/mvc/domain"
	"github.com/shakilbd009/go-microsrvcs/mvc/utils"
)

//GetUser returns a User object or an error if no user exist with that ID.
func GetUser(userID int64) (*domain.User, *utils.ApplicationError) {
	return domain.GetUser(userID)
}
