package rest_errors

import "net/http"

type RestError struct {
	Message    string   `json:"message"`
	StatusCode int      `json:"status_code"`
	Err        string   `json:"error"`
	Causes     []Causes `json:"causes"`
}

type Causes struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func (r *RestError) Error() string {
	return r.Message
}

func NewRestError(message string, status int, err string, causes []Causes) *RestError {
	return &RestError{
		Message:    message,
		StatusCode: status,
		Err:        err,
		Causes:     causes,
	}
}

func NewBadRequestError(msg string) *RestError {
	return &RestError{
		Message:    msg,
		StatusCode: http.StatusBadRequest,
		Err:        "bad_request",
	}
}

func NewBadRequestValidateError(msg string, causes []Causes) *RestError {
	return &RestError{
		Message:    msg,
		StatusCode: http.StatusBadRequest,
		Err:        "bad_request",
		Causes:     causes,
	}
}

func NewInternalServerError(msg string) *RestError {
	return &RestError{
		Message:    msg,
		StatusCode: http.StatusInternalServerError,
		Err:        "internal_server_error",
	}
}

func NewNotFoundError(msg string) *RestError {
	return &RestError{
		Message:    msg,
		StatusCode: http.StatusNotFound,
		Err:        "not_found",
	}
}

func NewUnauthorizedError(msg string) *RestError {
	return &RestError{
		Message:    msg,
		StatusCode: http.StatusUnauthorized,
		Err:        "unauthorized",
	}
}

func NewForbiddenError(msg string) *RestError {
	return &RestError{
		Message:    msg,
		StatusCode: http.StatusForbidden,
		Err:        "forbidden",
	}
}
