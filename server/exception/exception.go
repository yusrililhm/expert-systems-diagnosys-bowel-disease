package exception

import "net/http"

type Exception interface {
	Status() int
	Message() string
	Error() string
}

type ErrorData struct {
	StatusError  int    `json:"status"`
	MessageError string `json:"message"`
	ErrorType    string `json:"error"`
}

// Error implements Exception.
func (e *ErrorData) Error() string {
	return e.ErrorType
}

// Message implements Exception.
func (e *ErrorData) Message() string {
	return e.MessageError
}

func (e *ErrorData) Status() int {
	return e.StatusError
}

func NewInternalServerError(message string) Exception {
	return &ErrorData{
		StatusError:  http.StatusInternalServerError,
		MessageError: message,
		ErrorType:    http.StatusText(http.StatusInternalServerError),
	}
}

func NewConflictError(message string) Exception {
	return &ErrorData{
		StatusError:  http.StatusConflict,
		MessageError: message,
		ErrorType:    http.StatusText(http.StatusConflict),
	}
}

func NewNotFound(message string) Exception {
	return &ErrorData{
		StatusError:  http.StatusNotFound,
		MessageError: message,
		ErrorType:    http.StatusText(http.StatusNotFound),
	}
}

func NewBadRequestError(message string) Exception {
	return &ErrorData{
		StatusError:  http.StatusBadRequest,
		MessageError: message,
		ErrorType:    http.StatusText(http.StatusBadRequest),
	}
}

func NewUnprocessableEntity(message string) Exception {
	return &ErrorData{
		StatusError:  http.StatusUnprocessableEntity,
		MessageError: message,
		ErrorType:    http.StatusText(http.StatusUnprocessableEntity),
	}
}

func NewUnauthenticatedError(message string) Exception {
	return &ErrorData{
		StatusError:  http.StatusUnauthorized,
		MessageError: message,
		ErrorType:    http.StatusText(http.StatusUnauthorized),
	}
}

func NewUnauthorizedError(message string) Exception {
	return &ErrorData{
		StatusError:  http.StatusForbidden,
		MessageError: message,
		ErrorType:    http.StatusText(http.StatusForbidden),
	}
}
