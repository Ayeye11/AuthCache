package api

import (
	"time"

	"github.com/Ayeye11/se-thr/internal/database/repository"
	"github.com/Ayeye11/se-thr/internal/router/controllers"
	"github.com/Ayeye11/se-thr/internal/services"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type api struct {
	handler  *gin.Engine
	db       *gorm.DB
	cacheDB  *redis.Client
	cacheTTL time.Duration
	tokenKey string
}

func NewRouter(db *gorm.DB, cacheDB *redis.Client, cacheTTL time.Duration, tokenKey string) *api {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	return &api{r, db, cacheDB, cacheTTL, tokenKey}
}

func (r *api) RegisterRoutes() *gin.Engine {
	// Repository & Services:
	repo := repository.LoadRepository(r.db)
	services := services.LoadServices(repo, r.tokenKey)

	// Controller:
	controller := controllers.InitController(services, r.cacheDB, r.cacheTTL)
	controller.RegisterRoutes(r.handler)

	// Return: 'Handler'
	return r.handler
}
