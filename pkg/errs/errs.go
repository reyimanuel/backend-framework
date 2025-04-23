package errs

import (
	"net/http"
)

// Response is a struct that represents the response structure for the API.
// It contains the status code, message, and data fields.
type MessageError interface {
	Status() int
	Error() string
	Message() string
}

// ErrorData is a struct that implements the MessageError interface.
// It contains the status code, error message, and additional message fields.
type ErrorData struct {
	ErrStatus  int    `json:"status"`
	ErrError   string `json:"error"`
	ErrMessage string `json:"message"`
}

// Returns the status code of the response.
func (e *ErrorData) Status() int {
	return e.ErrStatus
}

// Return the error.
func (e *ErrorData) Error() string {
	return e.ErrError
}

// Return a message associated with the error.
func (e *ErrorData) Message() string {
	return e.ErrMessage
}

// Client Error Responses (400s)
func BadRequest(message string) MessageError {
	return &ErrorData{
		ErrStatus:  http.StatusBadRequest,
		ErrError:   "Bad Request",
		ErrMessage: message,
	}
}

func Unauthorized(message string) MessageError {
	return &ErrorData{
		ErrMessage: message,
		ErrStatus:  http.StatusUnauthorized,
		ErrError:   "Unauthorized",
	}
}

func Forbidden(message string) MessageError {
	return &ErrorData{
		ErrMessage: message,
		ErrStatus:  http.StatusForbidden,
		ErrError:   "Forbidden",
	}
}

func NotFound(message string) MessageError {
	return &ErrorData{
		ErrMessage: message,
		ErrStatus:  http.StatusNotFound,
		ErrError:   "Not Found",
	}
}

func InternalServerError(message string) MessageError {
	return &ErrorData{
		ErrMessage: message,
		ErrStatus:  http.StatusInternalServerError,
		ErrError:   "Internal Server Error",
	}
}
