package middlewares

import "github.com/Ayeye11/se-thr/internal/services"

type Middlewares interface {
}

func LoadMiddlewares(svc services.AuthService) Middlewares {
	return &middl{svc}
}

type middl struct {
	svc services.AuthService
}
