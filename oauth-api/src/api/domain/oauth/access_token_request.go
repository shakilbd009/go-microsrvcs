package oauth

import (
	"strings"

	"github.com/shakilbd009/go-microsrvcs/src/api/utils/errors"
)

type (
	AccessTokenRequest struct {
		UserName string `json:"username"`
		Password string `json:"password"`
	}
)

func (r *AccessTokenRequest) Validate() errors.ApiError {

	r.UserName = strings.TrimSpace(r.UserName)
	if r.UserName == "" {
		return errors.NewBadRequestError("invalid username")
	}

	r.Password = strings.TrimSpace(r.Password)
	if r.Password == "" {
		return errors.NewBadRequestError("invalid password")
	}
	return nil
}
