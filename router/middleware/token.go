package middleware

import (
	"blog/config"
	"blog/core"
)

var Token core.HandlerFunc = func(c *core.Context) core.Response {
	token := c.PostForm("token")
	if token == "" {
		token = c.Query("token")
	}
	if token != config.App.Common.Token {
		return c.Fail(201,"token error!")
	}
	c.Next()

	return nil
}



