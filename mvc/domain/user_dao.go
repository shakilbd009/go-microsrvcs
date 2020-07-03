package domain

import (
	"fmt"
	"net/http"

	"github.com/shakilbd009/go-microsrvcs/mvc/utils"
)

var (
	users = map[int64]*User{
		123: {Id: 1, FirstName: "shakil", LastName: "akram", Email: "sha@sha.com"},
	}
)

//GetUser returns a User or an error.
func GetUser(userID int64) (*User, *utils.ApplicationError) {

	if user := users[userID]; user != nil {
		return user, nil
	}
	return nil, &utils.ApplicationError{
		Message:    fmt.Sprintf("user %v was not found", userID),
		StatusCode: http.StatusNotFound,
		Code:       "not_found",
	}

}
