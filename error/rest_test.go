package error_rest

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

const (
	error_message = "error message"
)

func TestNewInternalServerError(t *testing.T) {
	cause := "database error"
	err := InternalServerError(error_message, errors.New(cause))
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.Status)
	assert.EqualValues(t, error_message, err.Message)
	assert.EqualValues(t, internal_server_error, err.Error)

	assert.NotNil(t, err.Causes)
	assert.EqualValues(t, 1, len(err.Causes))
	assert.EqualValues(t, cause, err.Causes[0])
}

func TestNewBadRequestError(t *testing.T) {
	err := BadRequest(error_message)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusBadRequest, err.Status)
	assert.EqualValues(t, error_message, err.Message)
	assert.EqualValues(t, bad_request, err.Error)
}

func TestNewNotFoundError(t *testing.T) {
	err := NotFound(error_message)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusNotFound, err.Status)
	assert.EqualValues(t, error_message, err.Message)
	assert.EqualValues(t, not_found, err.Error)
}

func TestNewConflictError(t *testing.T) {
	err := Conflict(error_message)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusConflict, err.Status)
	assert.EqualValues(t, error_message, err.Message)
	assert.EqualValues(t, conflict, err.Error)
}

func TestNewError(t *testing.T) {
	err := NewError(error_message)
	assert.NotNil(t, err)
}