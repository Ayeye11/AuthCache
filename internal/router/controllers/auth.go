package controllers

import (
	"github.com/Ayeye11/se-thr/internal/common/errs"
	"github.com/Ayeye11/se-thr/internal/common/types"
	"github.com/gin-gonic/gin"
)

func (h *handler) registerHandler(c *gin.Context) {
	// Request
	p := &types.User{}
	if err := h.req.GetBodyRequest(c, p, false); err != nil {
		h.res.SendError(c, err)
		return
	}

	// Services
	hash, err := h.svc.Hash.HashPassword(p.Password)
	if err != nil {
		h.res.SendError(c, errs.InternalX(err))
		return
	}
	p.Password = hash

	role, err := h.svc.Auth.GetRole("client")
	if err != nil {
		h.res.SendError(c, errs.InternalX(err))
		return
	}
	p.Role = &types.Role{ID: role.ID}

	if err := h.svc.User.RegisterUser(p); err != nil {

		if errs.ErrIs(err, errs.ErrSvcUser_ConflictEmail) {
			h.res.SendError(c, errs.ErrHttpAlreadyExistEmail)
			return
		}

		h.res.SendError(c, errs.InternalX(err))
		return
	}

	h.res.SendMessage(c, 201, "register successfully")
}

func (h *handler) loginHandler(c *gin.Context) {
	// Request
	p := &types.User{}
	if err := h.req.GetBodyRequest(c, p, true, types.UserEmail, types.UserPassword); err != nil {
		h.res.SendError(c, err)
		return
	}

	// Service: Login
	u, err := h.svc.User.GetUser(p.Email)
	if err != nil {
		switch {

		case errs.ErrIs(err, errs.ErrSvcUser_NotFoundUser):
			h.res.SendError(c, errs.ErrHttpInvalidLogin)
			return

		default:
			h.res.SendError(c, errs.InternalX(err))
			return
		}
	}

	if !h.svc.Hash.ComparePasswords(u.Password, p.Password) {
		h.res.SendError(c, errs.ErrHttpInvalidLogin)
		return
	}

	// => Get token
	token, err := h.svc.Auth.CreateToken(u)
	if err != nil {
		h.res.SendError(c, errs.ErrHttpInvalidLogin)
		return
	}

	// Response
	h.res.SetCookie(c, token)
	h.res.SendMessage(c, 200, "login successfully")
}
