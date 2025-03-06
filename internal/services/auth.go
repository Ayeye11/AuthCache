package services

import (
	"strconv"

	"github.com/Ayeye11/se-thr/internal/common/errs"
	"github.com/Ayeye11/se-thr/internal/common/types"
	"github.com/Ayeye11/se-thr/internal/database/repository"
	"github.com/golang-jwt/jwt/v5"
)

type authSvc struct {
	permRepo repository.PermissionRepository
	tokenKey string
}

func setClaims(u *types.User) jwt.MapClaims {
	return jwt.MapClaims{
		"sub":       strconv.Itoa(int(u.ID)),
		"firstname": u.Firstname,
		"lastname":  u.Lastname,
		"role_id":   strconv.Itoa(int(u.Role.ID)),
	}
}

// Token
func (s *authSvc) CreateToken(u *types.User) (string, error) {
	claims := setClaims(u)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(s.tokenKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (s *authSvc) CheckToken(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (any, error) {

		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errs.ErrSvcAuth_Invalidtoken
		}

		return []byte(s.tokenKey), nil
	})
	if err != nil {
		return nil, errs.ErrSvcAuth_Invalidtoken
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, errs.ErrSvcAuth_Invalidtoken
}

// Permissions
func (s *authSvc) GetRole(ident any) (*types.Role, error) {
	if ident == nil {
		return nil, errs.ErrSvcAuth_InvalidRoleID
	}

	checkFunc := func(v *types.Role, err error) (*types.Role, error) {
		if err == nil {
			return v, nil
		}

		return nil, errs.IsErrDoX(err, errs.ErrRepoPerm_InvalidRoleID, errs.ErrSvcAuth_InvalidRoleID)
	}

	switch v := ident.(type) {
	case int:
		return checkFunc(s.permRepo.GetRoleByID(v))

	case string:
		return checkFunc(s.permRepo.GetRoleByName(v))

	default:
		return nil, errs.ErrSvcAuth_InvalidRoleID
	}
}

func (s *authSvc) GetPermissions(roleID int) ([]*types.Permission, error) {

	perms, err := s.permRepo.GetPermissions(roleID)
	if err != nil {
		errs.IsErrDoX(err, errs.ErrRepoPerm_InvalidRoleID, errs.ErrSvcAuth_NotFoundRole)
	}

	return perms, nil
}
