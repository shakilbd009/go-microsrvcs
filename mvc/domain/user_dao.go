package domain

import (
	"fmt"
	"log"
	"net/http"

	"github.com/shakilbd009/go-microsrvcs/mvc/utils"
)

var (
	users = map[int64]*User{
		123: {Id: 123, FirstName: "shakil", LastName: "akram", Email: "sha@sha.com"},
	}
	UserDao userDaoInterface
)

type userDao struct{}

type userDaoInterface interface {
	GetUser(int64) (*User, *utils.ApplicationError)
}

func init() {
	UserDao = &userDao{}
}

//GetUser returns a User or an error.
func (u *userDao) GetUser(userID int64) (*User, *utils.ApplicationError) {
	log.Println("we are calling the database")
	if user := users[userID]; user != nil {
		return user, nil
	}
	return nil, &utils.ApplicationError{
		Message:    fmt.Sprintf("user %v was not found", userID),
		StatusCode: http.StatusNotFound,
		Code:       "not_found",
	}

}
