package apperror

import (
	"errors"
	"fmt"
	"net/http"
)

var ErrUnexpectedNillValues = NewInternalError(errors.New("unexpected nil values supplied"))

type Generic struct {
	error
	Message    string
	StatusCode int
}

func NewGenericError(err error) *Generic {
	return &Generic{
		error:      err,
		Message:    "generic",
		StatusCode: http.StatusTeapot,
	}
}

func (e *Generic) Error() string {
	return fmt.Sprintf("%s: %s", e.Message, e.error.Error())
}

func (e *Generic) Is(err error) bool {
	fmt.Println("comparing errors")
	_, ok := err.(*Generic)
	return ok
}

type AuthenticationError struct {
	*Generic
}

func NewAuthenticationError(err error) *AuthenticationError {
	return &AuthenticationError{
		Generic: &Generic{
			Message:    "failed to authenticate",
			StatusCode: http.StatusUnauthorized,
			error:      err,
		},
	}
}

func (e *AuthenticationError) Is(err error) bool {
	fmt.Println("comparing errors")
	_, ok := err.(*AuthenticationError)
	return ok
}

type InternalError struct {
	*Generic
}

func (e *InternalError) Is(err error) bool {
	fmt.Println("comparing errors")
	_, ok := err.(*InternalError)
	return ok
}

func NewInternalError(err error) *InternalError {
	return &InternalError{
		Generic: &Generic{
			Message:    "unexpected server error encountered",
			StatusCode: http.StatusInternalServerError,
			error:      err,
		},
	}
}

type BadRequestError struct {
	*Generic
}

func (e *BadRequestError) Is(err error) bool {
	fmt.Println("comparing errors")
	_, ok := err.(*BadRequestError)
	return ok
}

func NewBadRequestError(err error) *BadRequestError {
	return &BadRequestError{
		Generic: &Generic{
			Message:    "bad request provided",
			StatusCode: http.StatusBadRequest,
			error:      err,
		},
	}
}
