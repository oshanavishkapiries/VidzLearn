package response

import "net/http"

type APIError struct {
	StatusCode int
	Message    string
	Err        error
}

func (e *APIError) Error() string {
	return e.Message
}

func NewAPIError(status int, msg string, err error) *APIError {
	return &APIError{
		StatusCode: status,
		Message:    msg,
		Err:        err,
	}
}

/* Common helpers */

// 400 Bad Request
func BadRequestException(msg string, err error) *APIError {
	return NewAPIError(http.StatusBadRequest, msg, err)
}

// 404 Not Found
func NotFoundException(msg string, err error) *APIError {
	return NewAPIError(http.StatusNotFound, msg, err)
}

// 500 Internal Server Error
func InternalServerException(msg string, err error) *APIError {
	return NewAPIError(http.StatusInternalServerError, msg, err)
}

// 401 Unauthorized
func UnauthorizedException(msg string, err error) *APIError {
	return NewAPIError(http.StatusUnauthorized, msg, err)
}

// 403 Forbidden
func ForbiddenException(msg string, err error) *APIError {
	return NewAPIError(http.StatusForbidden, msg, err)
}

// 409 Conflict
func ConflictException(msg string, err error) *APIError {
	return NewAPIError(http.StatusConflict, msg, err)
}

// 422 Unprocessable Entity
func UnprocessableEntityException(msg string, err error) *APIError {
	return NewAPIError(http.StatusUnprocessableEntity, msg, err)
}
