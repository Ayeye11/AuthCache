package errs

import (
	"errors"
	"log"
	"runtime"
)

// Repository
var (
	// User
	ErrRepoUser_InvalidUserID   = BscError("invalid user ID")
	ErrRepoUser_DuplicatedEmail = BscError("email already used")
	ErrRepoUser_NotFound        = BscError("user not found")
	// Permissions
	ErrRepoPerm_InvalidRoleID = BscError("invalid role ID")
)

// Services
var (
	ErrSvc_InvalidID = BscError("invalid id")
	// Auth
	ErrSvcAuth_Invalidtoken  = BscError("invalid token")
	ErrSvcAuth_NotFoundRole  = BscError("role doesn't exists")
	ErrSvcAuth_InvalidRoleID = BscError("invalid role ID")

	// User
	ErrSvcUser_ConflictEmail = BscError("email is already used")
	ErrSvcUser_NotFoundUser  = BscError("user not found")
)

// Types validations
var (
	// All
	ErrValidation_MissingSpec        = BscError("missing field specifications")
	ErrValidation_MissingCredentials = BscError("missing credentials")
	// User
	ErrValidation_InvalidID        = BscError("invalid id")
	ErrValidation_InvalidEmail     = BscError("invalid email")
	ErrValidation_InvalidPassword  = BscError("invalid password")
	ErrValidation_InvalidFirstname = BscError("invalid firstname")
	ErrValidation_InvalidLastname  = BscError("invalid lastname")
	ErrValidation_InvalidAge       = BscError("invalid age")
	ErrValidation_PasswordNotHash  = BscError("the password must be hashed")
)

func UnknownError(err error) error {
	_, file, line, _ := runtime.Caller(1)
	log.Printf("unknown error at %s:%d: %v", file, line, err)
	return err
}

// IsErrDoX(MainERROR, Is, Do, Is, Do, Is, Do...)
func IsErrDoX(main error, isDo ...error) error {
	if main == nil {
		return nil
	}

	if len(isDo) == 0 || len(isDo)%2 != 0 {
		return UnknownError(BscError("expected pairs of (Is, Do) errors"))
	}

	for i := 0; i < len(isDo); i += 2 {
		if errors.Is(main, isDo[i]) {
			return isDo[i+1]
		}
	}

	return UnknownError(main)
}
