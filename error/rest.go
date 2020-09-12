package error_rest

import (
	"errors"
	"net/http"
)

const (
	bad_request           = "bad_request"
	conflict              = "conflict"
	not_found             = "not_found"
	internal_server_error = "internal_server_error"
)

type RestErr struct {
	Message string        `json:"message"`
	Status  int           `json:"status"`
	Error   string        `json:"error"`
	Causes  []interface{} `json:"causes"`
}

func NewError(msg string) error {
	return errors.New(msg)
}

func BadRequest(message string) *RestErr {
	return &RestErr{
		Message: message,
		Status:  http.StatusBadRequest,
		Error:   bad_request,
	}
}

func NotFound(message string) *RestErr {
	return &RestErr{
		Message: message,
		Status:  http.StatusNotFound,
		Error:   not_found,
	}
}

func Conflict(message string) *RestErr {
	return &RestErr{
		Message: message,
		Status:  http.StatusConflict,
		Error:   conflict,
	}
}

func InternalServerError(message string, err error) *RestErr {
	result := &RestErr{
		Message: message,
		Status:  http.StatusInternalServerError,
		Error:   internal_server_error,
	}

	if err != nil {
		result.Causes = append(result.Causes, err.Error())
	}

	return result
}
