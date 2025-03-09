package controllers

import (
	"time"

	"github.com/Ayeye11/AuthCache/internal/router/cache/rdb"
	"github.com/Ayeye11/AuthCache/internal/router/http"
	"github.com/Ayeye11/AuthCache/internal/router/middlewares"
	"github.com/Ayeye11/AuthCache/internal/services"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

type Controller interface {
	RegisterRoutes(*gin.Engine)
}

func InitController(services *services.Services, cacheDB *redis.Client, cacheTTL time.Duration) Controller {
	pkgHTTP := http.LoadPkgHTTP()
	cache := rdb.NewCache(cacheDB, cacheTTL)
	middls := middlewares.LoadMiddlewares(pkgHTTP, cache, services.Auth)

	return &handler{pkgHTTP.Req, pkgHTTP.Res, cache, middls, services}
}

type handler struct {
	req    http.RequestHTTP
	res    http.ResponseHTTP
	cache  rdb.RedisCache
	middls middlewares.Middlewares
	svc    *services.Services
}

func (h *handler) RegisterRoutes(r *gin.Engine) {
	api := r.Group("/api")

	auth := api.Group("/auth")
	auth.POST("/register", h.registerHandler)
	auth.POST("/login", h.loginHandler)
	auth.POST("/logout", h.logoutHandler)
}
