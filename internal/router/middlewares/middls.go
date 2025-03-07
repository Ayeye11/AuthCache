package middlewares

import (
	"strconv"

	"github.com/Ayeye11/se-thr/internal/common/errs"
	"github.com/Ayeye11/se-thr/internal/common/types"
	"github.com/Ayeye11/se-thr/internal/router/http"
	"github.com/Ayeye11/se-thr/internal/services"
	"github.com/gin-gonic/gin"
)

type Middlewares interface {
	IsAuth(sendClaims, sendPermissions bool) gin.HandlerFunc
	HasPermission(category, action string) gin.HandlerFunc
}

func LoadMiddlewares(pkgHTTP *http.PackageHTTP, svc services.AuthService) Middlewares {
	return &middl{pkgHTTP.Req, pkgHTTP.Res, svc}
}

type middl struct {
	req http.RequestHTTP
	res http.ResponseHTTP
	svc services.AuthService
}

func (m *middl) IsAuth(sendClaims, sendPermissions bool) gin.HandlerFunc {
	return func(c *gin.Context) {

		token, err := c.Cookie("token")
		if err != nil {
			m.res.SendError(c, errs.ErrHttpMissingToken)
			c.Abort()
			return
		}

		claims, err := m.svc.CheckToken(token)
		if err != nil {
			m.res.SendError(c, errs.ErrHttpInvalidToken)
			c.Abort()
			return
		}

		if sendClaims {
			c.Set("claims", claims)
		}

		if !sendPermissions {
			c.Next()
			return
		}

		val, ok := claims["role_id"].(string)
		if !ok {
			m.res.SendError(c, errs.InternalX(errs.BscError("fail to get role id")))
			c.Abort()
			return
		}
		roleID, err := strconv.Atoi(val)
		if err != nil {
			m.res.SendError(c, errs.InternalX(err))
			c.Abort()
			return
		}

		// TODO: Check role permissions on Redis...

		perms, err := m.svc.GetPermissions(roleID)
		if err != nil {
			m.res.SendError(c, errs.InternalX(err))
			c.Abort()
			return
		}

		c.Set("permissions", perms)
		c.Next()
	}
}

func (m *middl) HasPermission(category, action string) gin.HandlerFunc {
	return func(c *gin.Context) {

		val, ok := c.Get("permissions")
		if !ok {
			m.res.SendError(c, errs.InternalX(errs.BscError("fail to get permissions in context")))
			c.Abort()
			return
		}

		perms, ok := val.([]*types.Permission)
		if !ok {
			m.res.SendError(c, errs.InternalX(errs.BscError("fail to get role id")))
			c.Abort()
			return
		}

		for _, p := range perms {
			if p.Category == category && p.Action == action {
				c.Next()
				return
			}
		}

		m.res.SendError(c, errs.ErrHttpForbidden)
		c.Abort()
	}
}
