package services

import (
	"github.com/Ayeye11/AuthCache/internal/common/errs"
	"github.com/Ayeye11/AuthCache/internal/common/types"
	"github.com/Ayeye11/AuthCache/internal/database/repository"
)

type userSvc struct {
	userRepo repository.UserRepository
}

func (s *userSvc) RegisterUser(u *types.User) error {
	if err := u.Validate(false, types.UserPassword); err != nil {
		return err
	}

	if !u.IsPasswordHashed() {
		return errs.ErrValidation_PasswordNotHash
	}

	if err := s.userRepo.CreateUser(u); err != nil {
		return errs.IsErrDoX(err, errs.ErrRepoUser_DuplicatedEmail, errs.ErrSvcUser_ConflictEmail)
	}

	return nil
}

func (s *userSvc) GetUser(ident any) (*types.User, error) {
	if ident == nil {
		return nil, errs.ErrSvc_InvalidID
	}

	checkFunc := func(v *types.User, err error) (*types.User, error) {
		if err == nil {
			return v, nil
		}

		return nil, errs.IsErrDoX(err, errs.ErrRepoUser_NotFound, errs.ErrSvcUser_NotFoundUser)
	}

	switch v := ident.(type) {

	case string:
		return checkFunc(s.userRepo.GetUserByEmail(v))

	case int:
		return checkFunc(s.userRepo.GetUserByID(v))

	default:
		return nil, errs.ErrSvc_InvalidID
	}
}
