package apperror_test

import (
	"errors"
	"net/http"
	"testing"

	"github.com/LevinStrike/lux-backend/internal/core/apperror"
	"github.com/stretchr/testify/assert"
)

type testError struct {
	*apperror.Generic
}

func NewTestError(err string) *testError {
	return &testError{
		Generic: apperror.NewGenericError(errors.New(err)),
	}
}

func TestNewGenericError(t *testing.T) {
	tests := map[string]struct {
		err        error
		expected   string
		statusCode int
	}{
		"happy - expect error to return correct error string and status code": {
			err:        errors.New("test error"),
			expected:   "generic: test error",
			statusCode: http.StatusTeapot,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			err := apperror.NewGenericError(tc.err)
			assert.Equal(t, tc.expected, err.Error())
			assert.Equal(t, tc.statusCode, err.StatusCode)
			assert.True(t, err.Is(apperror.NewGenericError(tc.err)))
			assert.True(t, err.Is(apperror.NewGenericError(errors.New("XXyyXyYYx"))))
			assert.False(t, err.Is(apperror.NewInternalError(errors.New("XXyyXyYYx"))))
		})
	}
}

func TestNewAuthenticationError(t *testing.T) {
	tests := map[string]struct {
		err        error
		expected   string
		statusCode int
	}{
		"happy - expect error to return correct error string and status code": {
			err:        errors.New("test error"),
			expected:   "failed to authenticate: test error",
			statusCode: http.StatusUnauthorized,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			err := apperror.NewAuthenticationError(tc.err)
			assert.Equal(t, tc.expected, err.Error())
			assert.Equal(t, tc.statusCode, err.StatusCode)
			assert.True(t, err.Is(apperror.NewAuthenticationError(tc.err)))
			assert.True(t, err.Is(apperror.NewAuthenticationError(errors.New("XXyyXyYYx"))))
			assert.False(t, err.Is(apperror.NewInternalError(errors.New("XXyyXyYYx"))))
		})
	}
}

func TestNewInternalError(t *testing.T) {
	tests := map[string]struct {
		err        error
		expected   string
		statusCode int
	}{
		"happy - expect error to return correct error string and status code": {
			err:        errors.New("test error"),
			expected:   "unexpected server error encountered: test error",
			statusCode: http.StatusInternalServerError,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			err := apperror.NewInternalError(tc.err)
			assert.Equal(t, tc.expected, err.Error())
			assert.Equal(t, tc.statusCode, err.StatusCode)
			assert.True(t, err.Is(apperror.NewInternalError(tc.err)))
			assert.True(t, err.Is(apperror.NewInternalError(errors.New("XXyyXyYYx"))))
			assert.False(t, err.Is(apperror.NewAuthenticationError(errors.New("XXyyXyYYx"))))
		})
	}
}

func TestNewBadRequestError(t *testing.T) {
	tests := map[string]struct {
		err        error
		expected   string
		statusCode int
	}{
		"happy - expect error to return correct error string and status code": {
			err:        errors.New("test error"),
			expected:   "bad request provided: test error",
			statusCode: http.StatusBadRequest,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			err := apperror.NewBadRequestError(tc.err)
			assert.Equal(t, tc.expected, err.Error())
			assert.Equal(t, tc.statusCode, err.StatusCode)
			assert.True(t, err.Is(apperror.NewBadRequestError(tc.err)))
			assert.True(t, err.Is(apperror.NewBadRequestError(errors.New("XXyyXyYYx"))))
			assert.False(t, err.Is(apperror.NewAuthenticationError(errors.New("XXyyXyYYx"))))
		})
	}
}
