package services

import (
	"github.com/Ayeye11/se-thr/internal/common/types"
	"github.com/Ayeye11/se-thr/internal/database/repository"
	"github.com/golang-jwt/jwt/v5"
)

type Services struct {
	Auth AuthService
	Hash HashService
	User UserService
}

func LoadServices(repo *repository.Repository, tokenKey string) *Services {
	return &Services{
		Auth: &authSvc{repo.Perm, tokenKey},
		Hash: &hashSvc{},
		User: &userSvc{repo.User},
	}
}

type AuthService interface {
	// Token
	CreateToken(u *types.User) (string, error)
	CheckToken(tokenString string) (jwt.MapClaims, error)
	// Permissions
	GetRole(ident any) (*types.Role, error)
	GetPermissions(roleID int) ([]*types.Permission, error)
}

type HashService interface {
	HashPassword(password string) (string, error)
	ComparePasswords(hash, password string) bool
}

type UserService interface {
	RegisterUser(u *types.User) error
	GetUser(ident any) (*types.User, error)
}
