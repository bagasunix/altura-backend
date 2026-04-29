package errors

import (
	"errors"
	"fmt"

	"altura-property/shared/constants"
)

type AppError struct {
	Code        constants.ErrorCode `json:"code"`
	Message     string              `json:"message"`
	Details     any                 `json:"details,omitempty"`
	DebugData   map[string]any      `json:"-"`
	Err         error               `json:"-"`
	AuditDetail string              `json:"-"`
}

func (e *AppError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("[%s] %s: %v", e.Code, e.Message, e.Err)
	}
	return fmt.Sprintf("[%s] %s", e.Code, e.Message)
}

func (e *AppError) Unwrap() error {
	return e.Err
}

func (e *AppError) HTTPStatus() int {
	if status, ok := constants.HTTPStatusMap[e.Code]; ok {
		return status
	}
	return 500
}

func (e *AppError) WithDetails(details any) *AppError {
	e.Details = details
	return e
}

func (e *AppError) WithDebug(key string, value any) *AppError {
	if e.DebugData == nil {
		e.DebugData = make(map[string]any)
	}
	e.DebugData[key] = value
	return e
}

func (e *AppError) WithAuditDetail(detail string) *AppError {
	e.AuditDetail = detail
	return e
}

func (e *AppError) Is(target error) bool {
	t, ok := target.(*AppError)
	if !ok {
		return false
	}
	return e.Code == t.Code
}

func New(code constants.ErrorCode, message string) *AppError {
	return &AppError{Code: code, Message: message}
}

func Newf(code constants.ErrorCode, format string, args ...any) *AppError {
	return &AppError{Code: code, Message: fmt.Sprintf(format, args...)}
}

func Wrap(code constants.ErrorCode, message string, err error) *AppError {
	return &AppError{Code: code, Message: message, Err: err}
}

func AsAppError(err error) *AppError {
	var appErr *AppError
	if errors.As(err, &appErr) {
		return appErr
	}
	return nil
}

func IsCode(err error, code constants.ErrorCode) bool {
	appErr := AsAppError(err)
	return appErr != nil && appErr.Code == code
}

func ErrBadRequest(message string) *AppError {
	return New(constants.ErrCodeBadRequest, message)
}

func ErrValidation(message string, details any) *AppError {
	return New(constants.ErrCodeValidation, message).WithDetails(details)
}

func ErrUnauthorized(message string) *AppError {
	return New(constants.ErrCodeUnauthorized, message)
}

func ErrForbidden(message string) *AppError {
	return New(constants.ErrCodeForbidden, message)
}

func ErrNotFound(resource string) *AppError {
	return Newf(constants.ErrCodeNotFound, "%s not found", resource)
}

func ErrConflict(message string) *AppError {
	return New(constants.ErrCodeConflict, message)
}

func ErrInternal(err error) *AppError {
	return Wrap(constants.ErrCodeInternal, "internal server error", err)
}

func ErrDatabase(err error) *AppError {
	return Wrap(constants.ErrCodeDatabase, "database error", err)
}

func ErrCache(err error) *AppError {
	return Wrap(constants.ErrCodeCache, "cache error", err)
}

func ErrTimeout(message string) *AppError {
	return New(constants.ErrCodeTimeout, message)
}

func ErrExternalService(service string, err error) *AppError {
	return Wrap(constants.ErrCodeExternalService, fmt.Sprintf("external service error: %s", service), err)
}
