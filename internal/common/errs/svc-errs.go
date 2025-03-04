package errs

type appError struct {
	code    string
	message string
}

func newAppError(code, message string) error {
	return &appError{code, message}
}

func (err *appError) Error() string {
	return err.message
}

// Codes:
const (
	CodeBadRequest   string = "CODE_BAD_REQUEST"
	CodeUnauthorized string = "CODE_UNAUTHORIZED"
	CodeNotFound     string = "CODE_NOT_FOUND"
	CodeForbidden    string = "CODE_FORBIDDEN"
	CodeConflict     string = "CODE_CONFLICT"

	CodeInternal string = "CODE_INTERNAL"
)

// Auth
var (
	ErrBadRequest   = func(msg string) error { return newAppError(CodeBadRequest, msg) }
	ErrUnauthorized = func(msg string) error { return newAppError(CodeUnauthorized, msg) }
	ErrForbidden    = func(msg string) error { return newAppError(CodeForbidden, msg) }
	ErrNotfound     = func(msg string) error { return newAppError(CodeNotFound, msg) }
	ErrConflict     = func(msg string) error { return newAppError(CodeConflict, msg) }
)

var (
	ErrInternal = func(msg string) error { return newAppError(CodeInternal, msg) }
)
