package errs

// Repository
var (
	// User
	ErrRepoUserInvalidUserID   = BscError("invalid user ID")
	ErrRepoUserDuplicatedEmail = BscError("email already used")
	ErrRepoUserNotFound        = BscError("user not found")
	ErrRepoUserPasswordHash    = BscError("password is not hashed")
	// Permissions
	ErrRepoPermInvalidRoleID = BscError("invalid role ID")
)

// Services
var (
	// Auth
	ErrSvcAuthInvalidtoken = BscError("invalid token")
	ErrSvcAuthNotFoundRole = BscError("role doesn't exists")
)
