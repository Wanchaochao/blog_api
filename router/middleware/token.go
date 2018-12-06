package middleware

import (
	"blog/config"
	"blog/core"
	"blog/models"
)

var Token core.HandlerFunc = func(c *core.Context) core.Response {
	token := c.Request.Header.Get("Access-Token")
	user := &models.Users{
		Token: token,
	}
	if token != config.App.Common.Token {
		return c.Fail(201, "token error!")
	}
	c.Next()

	return nil
}
