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

func CompareError(err, with error) bool {
	return errors.Is(err, with)
}

func CompareStatus(err error, with int) bool {
	errHTTP := ToHTTP(err)
	return errHTTP.status == with
}

func CompareErrCode(err error, code string) bool {
	errSvc, ok := err.(*appError)
	if !ok {
		return false
	}

	return errSvc.code == code
}

func BscError(message string) error {
	return errors.New(message)
}
