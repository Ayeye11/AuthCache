package errs

import (
	"errors"
	"fmt"
)

func ToHTTP(err error) *errorHTTP {
	if errHTTP, ok := err.(*errorHTTP); ok {
		return errHTTP
	}

	msg := fmt.Sprintf("error: Failed to convert error to HTTP: %s\n", err)
	return &errorHTTP{500, msg}
}

func ErrIs(err, with error) bool {
	return errors.Is(err, with)
}

func BscError(message string) error {
	return errors.New(message)
}
