package constants

type ErrorCode string

const (
	ErrCodeBadRequest          ErrorCode = "BAD_REQUEST"
	ErrCodeValidation          ErrorCode = "VALIDATION_ERROR"
	ErrCodeUnauthorized        ErrorCode = "UNAUTHORIZED"
	ErrCodeForbidden           ErrorCode = "FORBIDDEN"
	ErrCodeNotFound            ErrorCode = "NOT_FOUND"
	ErrCodeConflict            ErrorCode = "CONFLICT"
	ErrCodeUnprocessableEntity ErrorCode = "UNPROCESSABLE_ENTITY"
	ErrCodeTooManyRequests     ErrorCode = "TOO_MANY_REQUESTS"

	ErrCodeInternal        ErrorCode = "INTERNAL_SERVER_ERROR"
	ErrCodeDatabase        ErrorCode = "DATABASE_ERROR"
	ErrCodeCache           ErrorCode = "CACHE_ERROR"
	ErrCodeExternalService ErrorCode = "EXTERNAL_SERVICE_ERROR"
	ErrCodeTimeout         ErrorCode = "TIMEOUT"
)

var HTTPStatusMap = map[ErrorCode]int{
	ErrCodeBadRequest:          400,
	ErrCodeValidation:          422,
	ErrCodeUnauthorized:        401,
	ErrCodeForbidden:           403,
	ErrCodeNotFound:            404,
	ErrCodeConflict:            409,
	ErrCodeUnprocessableEntity: 422,
	ErrCodeTooManyRequests:     429,
	ErrCodeInternal:            500,
	ErrCodeDatabase:            500,
	ErrCodeCache:               500,
	ErrCodeExternalService:     502,
	ErrCodeTimeout:             504,
}
