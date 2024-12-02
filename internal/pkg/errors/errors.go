package errors

import "net/http"

type Errors interface {
	Status() int
	Message() string
	Error() string
}

type ErrorsData struct {
	StatusError  int    `json:"status"`
	MessageError string `json:"message"`
	ErrorType    string `json:"error"`
}

// Error implements Errors.
func (e *ErrorsData) Error() string {
	return e.ErrorType
}

// Message implements Errors.
func (e *ErrorsData) Message() string {
	return e.MessageError
}

// Status implements Errors.
func (e *ErrorsData) Status() int {
	return e.StatusError
}

func NewInternalServerError(message string) Errors {
	return &ErrorsData{
		StatusError:  http.StatusInternalServerError,
		MessageError: message,
		ErrorType:    http.StatusText(http.StatusInternalServerError),
	}
}

func NewNotFound(message string) Errors {
	return &ErrorsData{
		StatusError:  http.StatusNotFound,
		MessageError: message,
		ErrorType:    http.StatusText(http.StatusNotFound),
	}
}
