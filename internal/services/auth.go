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
		"sub":     strconv.Itoa(int(u.ID)),
		"role_id": strconv.Itoa(int(u.Role.ID)),
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
func (s *authSvc) GetPermissions(roleID int) ([]*types.Permission, error) {

	perms, err := s.permRepo.GetPermissions(roleID)
	if err != nil {

		if errs.CompareError(err, errs.ErrRepoPerm_InvalidRoleID) {
			return nil, errs.ErrSvcAuth_NotFoundRole
		}

		return nil, err
	}

	return perms, nil

}
