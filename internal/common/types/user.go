package types

import (
	"github.com/Ayeye11/AuthCache/internal/common/errs"
	"github.com/Ayeye11/AuthCache/internal/common/validations"
)

// Type
type User struct {
	ID        uint   `json:"id,omitempty"`
	Email     string `json:"email,omitempty"`
	Password  string `json:"password,omitempty"`
	Firstname string `json:"firstname,omitempty"`
	Lastname  string `json:"lastname,omitempty"`
	Age       int    `json:"age,omitempty"`
	Role      *Role  `json:"role"`
}

// Reference fields
const (
	// User
	UserID        string = "user_id"
	UserEmail     string = "user_email"
	UserPassword  string = "user_password"
	UserFirstname string = "user_firstname"
	UserLastname  string = "user_lastname"
	UserAge       string = "user_age"
	UserRole      string = "user_role"
)

// Specifications
var (
	specEmail        = validations.NewSpec(false, 2, 100, errs.ErrValidation_InvalidEmail, validations.PatternEmail)
	specPassword     = validations.NewSpec(false, 2, 100, errs.ErrValidation_InvalidPassword)
	specPasswordHash = validations.NewSpec(false, 60, 60, errs.ErrValidation_InvalidPassword, validations.PatternBcryptHash)
	specFirstname    = validations.NewSpec(false, 2, 100, errs.ErrValidation_InvalidFirstname)
	specLastname     = validations.NewSpec(false, 2, 100, errs.ErrValidation_InvalidLastname)
	specAge          = validations.NewSpec(false, 18, 150, errs.ErrValidation_InvalidAge)
)

// ValideFields
var valideFields = map[string]bool{
	UserID:        true,
	UserEmail:     true,
	UserPassword:  true,
	UserFirstname: true,
	UserLastname:  true,
	UserAge:       true,
	UserRole:      true,
}

// Methods
func (u *User) Validate(whitelist bool, targets ...string) error {
	list := createList(valideFields, whitelist, targets...)

	if isExist(UserEmail, list) {
		if err := validations.ValidateField(u.Email, specEmail); err != nil {
			return err
		}
	}

	if isExist(UserPassword, list) {
		if err := validations.ValidateField(u.Password, specPassword); err != nil {
			return err
		}
	}

	if isExist(UserFirstname, list) {
		if err := validations.ValidateField(u.Firstname, specFirstname); err != nil {
			return err
		}
	}

	if isExist(UserLastname, list) {
		if err := validations.ValidateField(u.Lastname, specLastname); err != nil {
			return err
		}
	}

	if isExist(UserAge, list) {
		if err := validations.ValidateField(u.Age, specAge); err != nil {
			return err
		}
	}

	return nil
}

func (u *User) IsPasswordHashed() bool {
	err := validations.ValidateField(u.Password, specPasswordHash)
	return err == nil
}
