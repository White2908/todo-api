package utils

import (
	"fmt"
	"net/http"

	"github.com/go-fuego/fuego"
	"github.com/go-playground/validator/v10"
)

// Notice when not found data
type NotFoundError struct {
	Message string
	Err     error
}

// Combine message and return error
func (e NotFoundError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("%s: %v", e.Message, e.Err)
	}
	return e.Message
}

// Request is invalid
type BadRequestError struct {
	Message string
	Err     error
}

// Combine message and return error
func (e BadRequestError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("%s: %v", e.Message, e.Err)
	}
	return e.Message
}

// Internal server error
type InternalServerError struct {
	Message string
	Err     error
}

// Combine message and return error
func (e InternalServerError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("%s: %v", e.Message, e.Err)
	}
	return e.Message
}

// Error constructors
func NewNotFoundError(message string, err error) error {
	return NotFoundError{Message: message, Err: err}
}

func NewBadRequestError(message string, err error) error {
	return BadRequestError{Message: message, Err: err}
}

func NewInternalServerError(message string, err error) error {
	return InternalServerError{Message: message, Err: err}
}

// Convert custom errors to fuego.HTTPError
func ToFuegoError(err error) error {
	switch e := err.(type) {
	case NotFoundError:
		return fuego.HTTPError{
			Status: http.StatusNotFound,
			Err:    e,
		}
	case BadRequestError:
		return fuego.HTTPError{
			Status: http.StatusBadRequest,
			Err:    e,
		}
	case InternalServerError:
		return fuego.HTTPError{
			Status: http.StatusInternalServerError,
			Err:    e,
		}
	default:
		return fuego.HTTPError{
			Status: http.StatusInternalServerError,
			Err:    fmt.Errorf("internal server error"),
		}
	}
}

// Struct validation errors
func ValidateStruct(s interface{}) error {
	validate := validator.New()
	return validate.Struct(s)
}
