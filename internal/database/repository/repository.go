package repository

import (
	"github.com/Ayeye11/se-thr/internal/database/models"
	"gorm.io/gorm"
)

type Repository struct {
	Auth AuthRepository
}

func LoadRepository(db *gorm.DB) *Repository {
	return &Repository{
		Auth: &authStore{db},
	}
}

type AuthRepository interface {
	GetPermissions(role string) ([]models.Permission, error)
}
