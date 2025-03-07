package controllers

import (
	"fmt"

	"github.com/Ayeye11/se-thr/internal/router/http"
	"github.com/Ayeye11/se-thr/internal/router/middlewares"
	"github.com/Ayeye11/se-thr/internal/services"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
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
	r.GET("/test",
		h.middls.IsAuth(true, true),
		h.middls.HasPermission("stock", "view"),
		h.testHandler)

	api := r.Group("/api")

	auth := api.Group("/auth")
	auth.POST("/register", h.registerHandler)
	auth.POST("/login", h.loginHandler)
	auth.POST("/logout", h.logoutHandler)
}

func (h *handler) testHandler(c *gin.Context) {
	val, ok := c.Get("claims")
	if !ok {
		h.res.SendMessage(c, 400, "fail get from ctx")
		return
	}

	claims, ok := val.(jwt.MapClaims)
	if !ok {
		h.res.SendMessage(c, 400, "fail convert claims")
		return
	}

	firstname, ok := claims["firstname"].(string)
	if !ok {
		h.res.SendMessage(c, 400, "fail get firstname from claims")
		return
	}

	lastname, ok := claims["lastname"].(string)
	if !ok {
		h.res.SendMessage(c, 400, "fail get lastname from claims")
		return
	}

	msg := fmt.Sprintf("hello %s %s!, all ok", firstname, lastname)

	h.res.SendMessage(c, 200, msg)
}
