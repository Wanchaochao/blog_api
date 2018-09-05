package middleware

import (
	"blog/config"
	"blog/core"
)

var Token core.HandlerFunc = func(c *core.Context) core.Response {
	token := c.Request.Header.Get("Access-Token")

	if token != config.App.Common.Token {
		return c.Fail(201,"token error!")
	}
	c.Next()

	return nil
}



