package controllers

import (
	"github.com/Ayeye11/se-thr/internal/router/http"
	"github.com/Ayeye11/se-thr/internal/router/middlewares"
	"github.com/Ayeye11/se-thr/internal/services"
	"github.com/gin-gonic/gin"
)

type Controller interface {
	RegisterRoutes(*gin.Engine)
}

func InitController(services *services.Services) Controller {
	pkgHTTP := http.LoadPkgHTTP()
	middls := middlewares.LoadMiddlewares(pkgHTTP, services.Auth)

	return &handler{pkgHTTP.Req, pkgHTTP.Res, middls, services}
}

type handler struct {
	req    http.RequestHTTP
	res    http.ResponseHTTP
	middls middlewares.Middlewares
	svc    *services.Services
}

func (h *handler) RegisterRoutes(r *gin.Engine) {
	api := r.Group("/api")

	auth := api.Group("/auth")
	auth.POST("/register", h.registerHandler)
	auth.POST("/login", h.loginHandler)
}
