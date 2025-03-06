package services

import (
	"github.com/Ayeye11/se-thr/internal/common/errs"
	"github.com/Ayeye11/se-thr/internal/common/types"
	"github.com/Ayeye11/se-thr/internal/database/repository"
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
		if errs.CompareError(err, errs.ErrRepoUser_DuplicatedEmail) {
			return errs.ErrSvcUser_ConflictEmail
		}

		return errs.UnknownError(err)
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
		if errs.CompareError(err, errs.ErrRepoUser_NotFound) {
			return nil, errs.ErrSvcUser_NotFoundUser
		}

		return nil, errs.UnknownError(err)
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
