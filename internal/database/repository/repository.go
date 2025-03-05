package repository

import (
	"github.com/Ayeye11/se-thr/internal/common/types"
	"github.com/Ayeye11/se-thr/internal/database/models"
	"gorm.io/gorm"
)

type Repository struct {
	Perm PermissionRepository
	User UserRepository
}

func LoadRepository(db *gorm.DB) *Repository {
	return &Repository{
		Perm: &permStore{db},
		User: &userStore{db},
	}
}

type PermissionRepository interface {
	GetRoleByID(roleID int) (*models.AcRole, error)
	GetRoleByName(roleName string) (*models.AcRole, error)
	GetPermissions(roleID int) ([]*types.Permission, error)
}

type UserRepository interface {
	// Create
	CreateUser(u *types.User) error
	// Read
	GetUserByID(id int) (*types.User, error)
	GetUserByEmail(email string) (*types.User, error)
}
