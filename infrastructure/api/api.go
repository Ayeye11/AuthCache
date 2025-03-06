package api

import (
	"github.com/Ayeye11/se-thr/internal/database/repository"
	"github.com/Ayeye11/se-thr/internal/router/controllers"
	"github.com/Ayeye11/se-thr/internal/services"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type api struct {
	handler  *gin.Engine
	db       *gorm.DB
	tokenKey string
}

func NewRouter(db *gorm.DB, tokenKey string) *api {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	return &api{r, db, tokenKey}
}

func (r *api) RegisterRoutes() *gin.Engine {
	// Repository & Services:
	repo := repository.LoadRepository(r.db)
	services := services.LoadServices(repo, r.tokenKey)

	// Controller:
	controller := controllers.InitController(services)
	controller.RegisterRoutes(r.handler)

	// Return: 'Handler'
	return r.handler
}
