package middleware

import (
	"github.com/gin-gonic/gin"
)

func AuthDebug() gin.HandlerFunc {
	return func(c *gin.Context) {
		if pwd := c.Query("pwd"); pwd != "123456" {
			return
		}
		c.Next()
	}
}
