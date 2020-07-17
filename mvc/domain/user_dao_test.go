package domain

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetUserNoUserFound(t *testing.T) {

	user, err := UserDao.GetUser(0)
	assert.Nil(t, user, "we were not expecting a user with id 0")
	assert.NotNil(t, err, "we are expecting an error when user id is 0")
	assert.EqualValues(t, "not_found", err.Code)
	assert.EqualValues(t, "user 0 was not found", err.Message)
	assert.EqualValues(t, http.StatusNotFound, err.StatusCode, "we were expecting 404 when user is not found")

}

func TestGetUserNoError(t *testing.T) {

	user, err := UserDao.GetUser(123)
	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.EqualValues(t, "shakil", user.FirstName)
}
