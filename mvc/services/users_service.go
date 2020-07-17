package services

import (
	"github.com/shakilbd009/go-microsrvcs/mvc/domain"
	"github.com/shakilbd009/go-microsrvcs/mvc/utils"
)

type usersService struct {
}

var (
	UsersService usersService
)

//GetUser returns a User object or an error if no user exist with that ID.
func (u *usersService) GetUser(userID int64) (*domain.User, *utils.ApplicationError) {

	user, err := domain.UserDao.GetUser(userID)
	if err != nil {
		return nil, err
	}
	return user, nil
}
