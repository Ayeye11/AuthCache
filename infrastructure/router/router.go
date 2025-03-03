package router

import "github.com/gin-gonic/gin"

type router struct {
	handler *gin.Engine
}

func NewRouter() *router {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	return &router{r}
}

func (r *router) RegisterRoutes() *gin.Engine {
	// Dependecies:
	// asd

	// Return: 'Handler'
	return r.handler
}
