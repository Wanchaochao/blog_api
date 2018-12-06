package middleware

import (
	"blog/config"
	"blog/core"
	"blog/models"
	"github.com/ilibs/gosql"
)

var Token core.HandlerFunc = func(c *core.Context) core.Response {
	token := c.Request.Header.Get("Access-Token")
	if token != config.App.Common.Token {
		return c.Fail(201, "token error!")
	}
	user := &models.Users{}
	err := gosql.Model(user).Where("token = ?", token).Get()
	if err != nil {
		return c.Fail(204, "Access-Token 非法")
	}
	c.Next()

	return nil
}
