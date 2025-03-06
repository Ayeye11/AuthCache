package http

import "github.com/gin-gonic/gin"

type PackageHTTP struct {
	Req RequestHTTP
	Res ResponseHTTP
}

func LoadPkgHTTP() *PackageHTTP {
	return &PackageHTTP{
		Req: &request{},
		Res: &response{},
	}
}

type RequestHTTP interface {
	GetBodyRequest(c *gin.Context, p interface{ Validate(bool, ...string) error }, whitelist bool, targets ...string) error
}

type ResponseHTTP interface {
	SendMessage(c *gin.Context, status int, data any, message ...string)
	SendError(c *gin.Context, err error)
	SetCookie(c *gin.Context, token string)
}
