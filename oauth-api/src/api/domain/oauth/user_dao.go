package oauth

import (
	"github.com/shakilbd009/go-microsrvcs/src/api/utils/errors"
)

const (
	queryGetuserByUsernameAndPassword = "SELECT id, username from users WHERE username=? AND password=?"
)

var (
	users = map[string]*User{
		"shak": {Id: 123, Username: "shak"},
	}
)

func GetUserByUsernameAndPassword(username, password string) (*User, errors.ApiError) {
	user := users[username]
	if user == nil {
		return nil, errors.NewNotFoundApiError("no user found with given details")
	}
	return user, nil
}
