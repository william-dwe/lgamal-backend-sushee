package errorlist

import "net/http"

type AppError struct {
	StatusCode int    `json:"status_code"`
	Code       string `json:"code"`
	Message    string `json:"message"`
}

func (err AppError) Error() string {
	return err.Message
}

func BadRequestError(message string, code string) AppError {
	if code == "" {
		code = "BAD_REQUEST"
	}
	return AppError{
		Code:       code,
		Message:    message,
		StatusCode: http.StatusBadRequest,
	}
}

func NotFoundError(message string) AppError {
	return AppError{
		Code:       "NOT_FOUND_ERROR",
		Message:    message,
		StatusCode: http.StatusNotFound,
	}
}

func InternalServerError() AppError {
	return AppError{
		Code:       "INTERNAL_SERVER_ERROR",
		Message:    "there's an internal server error",
		StatusCode: http.StatusInternalServerError,
	}
}

func UnauthorizedError() AppError {
	return AppError{
		Code:       "UNAUTHORIZED_ERROR",
		Message:    "Unauthorized",
		StatusCode: http.StatusUnauthorized,
	}
}

func ForbiddenError() AppError {
	return AppError{
		Code:       "FORBIDDEN",
		Message:    "You don't have permission to access",
		StatusCode: http.StatusForbidden,
	}
}