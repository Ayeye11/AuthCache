package repository

import (
	"github.com/Ayeye11/se-thr/internal/database/models"
	"gorm.io/gorm"
)

type authStore struct {
	db *gorm.DB
}

func (s *authStore) GetPermissions(role string) ([]models.Permission, error) {
	var permissions []models.Permission

	err := s.db.
		Table("ac_relations").
		Select("role", "category", "action").
		Joins("JOIN ac_roles r ON role_id = r.id").
		Joins("JOIN ac_categories c ON category_id = c.id").
		Joins("JOIN ac_actions a ON action_id = a.id").
		Where("r.role = ?", role).
		Find(&permissions).Error
	if err != nil {
		return nil, err
	}

	return permissions, nil
}
