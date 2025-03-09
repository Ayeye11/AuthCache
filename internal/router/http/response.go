package http

import (
	"github.com/Ayeye11/AuthCache/internal/common/errs"
	"github.com/gin-gonic/gin"
)

type response struct{}

// Function behavior:
// - If no data is provided, it returns a JSON response with `nil`.
// - If the data is a string and no message is provided, it wraps the string in a "message" field.
// - If data is provided but no message, it wraps the data in a "data" field.
// - If both data and a message are provided, it returns both in a JSON response.
func (*response) SendMessage(c *gin.Context, status int, data any, message ...string) {
	hasMsg := len(message) != 0

	if data == nil {
		c.JSON(status, nil)
		return
	}

	if str, ok := data.(string); ok && !hasMsg {
		c.JSON(status, gin.H{"message": str})
		return
	}

	if !hasMsg {
		c.JSON(status, gin.H{"data": data})
		return
	}

	c.JSON(status, gin.H{"data": data, "message": message[0]})
}

func (*response) SendError(c *gin.Context, err error) {
	errHTTP := errs.ToHTTP(err)

	c.JSON(errHTTP.Status(), gin.H{"error": errHTTP.SafeMessage()})
}

func (*response) SetCookie(c *gin.Context, token string) {
	c.SetCookie("token", token, 0, "/", "", false, true)
}
