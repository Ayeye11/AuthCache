package types

import (
	"github.com/Ayeye11/se-thr/internal/common/constants"
	"github.com/Ayeye11/se-thr/internal/common/errs"
	"github.com/Ayeye11/se-thr/internal/common/validations"
)

// Type
type User struct {
	ID        int    `json:"id,omitempty"`
	Email     string `json:"email,omitempty"`
	Password  string `json:"password,omitempty"`
	Firstname string `json:"firstname,omitempty"`
	Lastname  string `json:"lastname,omitempty"`
	Age       int    `json:"age,omitempty"`
	RoleID    int    `json:"role,omitempty"`
}

// Specifications
var (
	specEmail        = validations.NewSpec(false, 2, 100, errs.InvalidEmail, validations.PatternEmail)
	specPassword     = validations.NewSpec(false, 2, 100, errs.InvalidPassword)
	specPasswordHash = validations.NewSpec(false, 60, 60, errs.InvalidPassword, validations.PatternBcryptHash)
	specFirstname    = validations.NewSpec(false, 2, 100, errs.InvalidFirstname)
	specLastname     = validations.NewSpec(false, 2, 100, errs.InvalidLastname)
	specAge          = validations.NewSpec(false, 18, 150, errs.InvalidAge)
)

// ValideFields
var valideFields = map[string]bool{
	constants.TypeUserID:        true,
	constants.TypeUserEmail:     true,
	constants.TypeUserPassword:  true,
	constants.TypeUserFirstname: true,
	constants.TypeUserLastname:  true,
	constants.TypeUserAge:       true,
	constants.TypeUserRole:      true,
}

// Methods
func (u *User) Validate(whitelist bool, targets ...string) error {
	list := createList(valideFields, whitelist, targets...)

	if isExist(constants.TypeUserEmail, list) {
		if err := validations.ValidateField(u.Email, specEmail); err != nil {
			return err
		}
	}

	if isExist(constants.TypeUserPassword, list) {
		if err := validations.ValidateField(u.Password, specPassword); err != nil {
			return err
		}
	}

	if isExist(constants.TypeUserFirstname, list) {
		if err := validations.ValidateField(u.Firstname, specFirstname); err != nil {
			return err
		}
	}

	if isExist(constants.TypeUserLastname, list) {
		if err := validations.ValidateField(u.Lastname, specLastname); err != nil {
			return err
		}
	}

	if isExist(constants.TypeUserAge, list) {
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
