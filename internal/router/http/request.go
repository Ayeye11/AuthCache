package http

import (
	"github.com/Ayeye11/AuthCache/internal/common/errs"
	"github.com/gin-gonic/gin"
)

type request struct{}

func (*request) GetBodyRequest(c *gin.Context, p interface{ Validate(bool, ...string) error }, whitelist bool, targets ...string) error {
	if err := c.BindJSON(p); err != nil {
		return errs.ErrHttpInvalidRequest
	}

	if err := p.Validate(whitelist, targets...); err != nil {
		return errs.NewErrorHTTP(400, err.Error())
	}

	return nil
}
