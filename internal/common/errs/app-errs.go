package errs

import (
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
	// Auth
	ErrSvcAuth_Invalidtoken = BscError("invalid token")
	ErrSvcAuth_NotFoundRole = BscError("role doesn't exists")

	// User
	ErrSvcUser_ConflictEmail = BscError("email is already used")
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
