package rest_error

import "net/http"

type RestError struct {
	Message    string       `json:"message"`
	Err        string       `json:"error"`
	Code       int          `json:"code"`
	Error_type []Error_type `json:"error_type,omitempty"`
}

type Error_type struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

// Error returns the error message of the RestError.
//
// This function does not take any parameters.
// It returns a string, which is the error message of the RestError.
func (e *RestError) Error() string {
	return e.Message
}

// NewRestError creates a new RestError object with the given message, error, code, and error type.
//
// message: The error message.
// err: The error string.
// code: The error code.
// error_type: The error type(s).
// Returns a pointer to the newly created RestError object.
func NewRestError(message string, err string, code int, error_type []Error_type) *RestError {
	return &RestError{
		Message:    message,
		Err:        err,
		Code:       code,
		Error_type: error_type,
	}
}

// NewBadRequestError creates a new BadRequestError with the given message.
//
// Parameters:
// - message: The error message.
//
// Returns:
// - *RestError: A pointer to the newly created BadRequestError.
func NewBadRequestError(message string) *RestError {
	return &RestError{
		Message: message,
		Err:     "bad_request",
		Code:    http.StatusBadRequest,
	}
}

// NewBadRequestValidationError creates a new BadRequestValidationError RestError.
//
// message: the error message.
// error_type: an array of error types.
// Returns a pointer to a RestError.
func NewBadRequestValidationError(message string, error_type []Error_type) *RestError {
	return &RestError{
		Message:    message,
		Err:        "bad_request",
		Code:       http.StatusBadRequest,
		Error_type: error_type,
	}
}

// NewInternalServerError creates a new internal server error RestError.
//
// Parameters:
// - message: the error message.
//
// Returns:
// - *RestError: a pointer to the newly created RestError.
func NewInternalServerError(message string) *RestError {
	return &RestError{
		Message: message,
		Err:     "internal_server_error",
		Code:    http.StatusInternalServerError,
	}
}

// NewNotFoundError creates a new RestError with a not found message.
//
// Parameters:
// - message: The error message.
//
// Returns:
// - *RestError: A pointer to a RestError object.
func NewNotFoundError(message string) *RestError {
	return &RestError{
		Message: message,
		Err:     "not_found",
		Code:    http.StatusNotFound,
	}
}

// NewUnauthorizedError creates a new instance of RestError with an unauthorized error.
//
// Parameters:
//   - message: the error message.
//
// Returns:
//   - *RestError: the new instance of RestError.
func NewUnauthorizedError(message string) *RestError {
	return &RestError{
		Message: message,
		Err:     "unauthorized",
		Code:    http.StatusUnauthorized,
	}
}
