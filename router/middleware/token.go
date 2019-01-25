package middleware

import (
	"blog_api/config"
	"blog_api/core"
	"blog_api/models"
	"github.com/ilibs/gosql"
)

var Token core.HandlerFunc = func(c *core.Context) core.Response {

	if config.App.Common.HttpTest {
		c.Next()
		return nil
	}

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
