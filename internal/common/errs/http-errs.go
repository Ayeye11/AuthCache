package errs

import "net/http"

type errorHTTP struct {
	status  int
	message string
}

func NewErrorHTTP(status int, message string) error {
	return &errorHTTP{status, message}
}

func (err *errorHTTP) Error() string {
	return err.message
}

func (err *errorHTTP) Status() int {
	return err.status
}

func (err *errorHTTP) SafeMessage() string {
	if err.Status() < 500 {
		return err.Error()
	}

	if msg, exists := serverErrorMessages[err.Status()]; exists {
		return msg
	}

	return "something went wrong"
}

// Client errors
var (
	// Status: 400
	ErrHttpInvalidRequest = NewErrorHTTP(http.StatusBadRequest, "invalid request")
	ErrHttpMissingRequest = NewErrorHTTP(http.StatusBadRequest, "missing request")

	ErrHttpInvalidCredentials = NewErrorHTTP(http.StatusBadRequest, "invalid credentials")
	ErrHttpMissingCredentials = NewErrorHTTP(http.StatusBadRequest, "missing credentials")

	// Status: 401
	ErrHttpInvalidToken = NewErrorHTTP(http.StatusUnauthorized, "invalid token")
	ErrHttpMissingToken = NewErrorHTTP(http.StatusUnauthorized, "missing token")

	ErrHttpInvalidLogin = NewErrorHTTP(http.StatusUnauthorized, "invalid email or password")

	// Status: 403
	ErrHttpForbidden = NewErrorHTTP(http.StatusForbidden, "access denied")

	// Status: 404
	ErrHttpNotFoundUser = NewErrorHTTP(http.StatusNotFound, "user not found")

	// Status: 409
	ErrHttpAlreadyExistEmail = NewErrorHTTP(http.StatusConflict, "email already used")
	ErrHttpAlreadyExistUser  = NewErrorHTTP(http.StatusContinue, "user already exist")
)

func InternalX(err error) error {
	return NewErrorHTTP(http.StatusInternalServerError, err.Error())
}

var serverErrorMessages = map[int]string{
	http.StatusInternalServerError: "internal server error",
	http.StatusNotImplemented:      "not implemented",
	http.StatusBadGateway:          "bad gateway",
	http.StatusServiceUnavailable:  "service unavailable",
	http.StatusGatewayTimeout:      "timeout gateway",
}
