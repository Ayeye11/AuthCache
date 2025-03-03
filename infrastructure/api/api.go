package api

import (
	"github.com/Ayeye11/se-thr/internal/api/controller"
	"github.com/gin-gonic/gin"
)

type api struct {
	handler *gin.Engine
}

func NewRouter() *api {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	return &api{r}
}

func (r *api) RegisterRoutes() *gin.Engine {
	// Dependencies:
	controller := controller.InitController()
	controller.RegisterRoutes(r.handler)

	// Return: 'Handler'
	return r.handler
}
