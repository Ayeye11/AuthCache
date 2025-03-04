package errs

import "net/http"

type errorHTTP struct {
	status  int
	message string
}

func newErrorHTTP(status int, message string) error {
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
	InvalidRequest = newErrorHTTP(http.StatusBadRequest, "invalid request")
	MissingRequest = newErrorHTTP(http.StatusBadRequest, "missing request")

	InvalidCredentials = newErrorHTTP(http.StatusBadRequest, "invalid credentials")
	MissingCredentials = newErrorHTTP(http.StatusBadRequest, "missing credentials")

	InvalidID        = newErrorHTTP(http.StatusBadRequest, "invalid id")
	InvalidEmail     = newErrorHTTP(http.StatusBadRequest, "invalid email")
	InvalidPassword  = newErrorHTTP(http.StatusBadRequest, "invalid password")
	InvalidFirstname = newErrorHTTP(http.StatusBadRequest, "invalid firstname")
	InvalidLastname  = newErrorHTTP(http.StatusBadRequest, "invalid lastname")
	InvalidAge       = newErrorHTTP(http.StatusBadRequest, "invalid age")

	// Status: 401
	InvalidToken = newErrorHTTP(http.StatusUnauthorized, "invalid token")
	MissingToken = newErrorHTTP(http.StatusUnauthorized, "missing token")

	InvalidLogin = newErrorHTTP(http.StatusUnauthorized, "invalid email or password")

	// Status: 403
	Forbidden = newErrorHTTP(http.StatusForbidden, "access denied")

	// Status: 404
	NotFoundUser = newErrorHTTP(http.StatusNotFound, "user not found")

	// Status: 409
	AlreadyExistEmail = newErrorHTTP(http.StatusConflict, "email already used")
	AlreadyExistUser  = newErrorHTTP(http.StatusContinue, "user already exist")
)

func InternalX(err error) error {
	return newErrorHTTP(http.StatusInternalServerError, err.Error())
}

var serverErrorMessages = map[int]string{
	http.StatusInternalServerError: "internal server error",
	http.StatusNotImplemented:      "not implemented",
	http.StatusBadGateway:          "bad gateway",
	http.StatusServiceUnavailable:  "service unavailable",
	http.StatusGatewayTimeout:      "timeout gateway",
}
