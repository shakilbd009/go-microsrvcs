package errors

import (
	"encoding/json"
	"errors"
	"net/http"
)

type ApiError interface {
	Status() int
	Message() string
	Error() string
}

type apiError struct {
	AStatus  int    `json:"status"`
	AMessage string `json:"message"`
	AnError  string `json:"error,omitempty"`
}

func (a *apiError) Status() int {
	return a.AStatus
}

func (a *apiError) Message() string {
	return a.AMessage
}

func (a *apiError) Error() string {
	return a.AnError
}

func NewApiError(statuscode int, msg string) ApiError {
	return &apiError{
		AStatus:  statuscode,
		AMessage: msg,
	}
}

func NewNotFoundApiError(msg string) ApiError {
	return &apiError{
		AStatus:  http.StatusNotFound,
		AMessage: msg,
	}
}
func NewApiErrorFromBytes(body []byte) (ApiError, error) {
	var result apiError
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, errors.New("invalid json body")
	}
	return &result, nil
}

func NewInternalServerError(msg string) ApiError {
	return &apiError{
		AStatus:  http.StatusInternalServerError,
		AMessage: msg,
	}
}

func NewBadRequestError(msg string) ApiError {
	return &apiError{
		AStatus:  http.StatusBadRequest,
		AMessage: msg,
	}
}
