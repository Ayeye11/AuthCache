package repository

import (
	"errors"

	"github.com/Ayeye11/se-thr/internal/common/errs"
	"github.com/Ayeye11/se-thr/internal/common/types"
	"github.com/Ayeye11/se-thr/internal/database/models"
	"gorm.io/gorm"
)

type permStore struct {
	db *gorm.DB
}

func (s *permStore) GetRoleByID(roleID int) (*models.AcRole, error) {
	if roleID < 1 {
		return nil, errs.ErrRepoPermInvalidRoleID
	}

	role := models.AcRole{}
	if err := s.db.First(&role, roleID).Error; err != nil {

		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errs.ErrRepoPermInvalidRoleID
		}

		return nil, err
	}

	return &role, nil
}

func (s *permStore) GetRoleByName(roleName string) (*models.AcRole, error) {
	role := models.AcRole{}
	if err := s.db.First(&role, "role = ?", roleName).Error; err != nil {

		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errs.ErrRepoPermInvalidRoleID
		}

		return nil, err
	}

	return &role, nil
}

func (s *permStore) GetPermissions(roleID int) ([]*types.Permission, error) {
	model := []*models.PermissionModel{}

	err := s.db.
		Table("ac_relations").
		Select("role", "category", "action").
		Joins("JOIN ac_roles r ON role_id = r.id").
		Joins("JOIN ac_categories c ON category_id = c.id").
		Joins("JOIN ac_actions a ON action_id = a.id").
		Where("r.id = ?", roleID).
		Find(&model).Error
	if err != nil {
		return nil, err
	}

	permissions := []*types.Permission{}
	for _, val := range model {
		perm := &types.Permission{
			Category: val.Category,
			Action:   val.Action,
		}
		permissions = append(permissions, perm)
	}

	return permissions, nil
}
