package errs

import "net/http"

type AppError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (a AppError) Error() string {
	return a.Message
}

func NewNotFoundError(msg string) AppError {
	return AppError{
		Code:    http.StatusNotFound,
		Message: msg,
	}
}

func NewUnexpectedError() AppError {
	return AppError{
		Code:    http.StatusInternalServerError,
		Message: "unexpected error",
	}
}

func NewInternalServerError() AppError {
	return AppError{
		Code:    http.StatusInternalServerError,
		Message: http.StatusText(http.StatusInternalServerError),
	}
}

func NewBadRequestError(msg string) AppError {
	if msg == "" {
		msg = http.StatusText(http.StatusBadRequest)
	}
	return AppError{
		Code:    http.StatusBadRequest,
		Message: msg,
	}
}

func NewUnauthorizedError(msg string) AppError {
	if msg == "" {
		msg = http.StatusText(http.StatusUnauthorized)
	}
	return AppError{
		Code:    http.StatusUnauthorized,
		Message: msg,
	}
}
