package repository

import "errors"

var (
	// User
	errUserInvalidUserID   = errors.New("invalid user ID")
	errUserDuplicatedEmail = errors.New("email already used")
	errUserNotFoundUser    = errors.New("user not found")
	errUserPasswordHash    = errors.New("password is not hashed")
	// Permissions
	errPermInvalidRoleID = errors.New("invalid role ID")
)
