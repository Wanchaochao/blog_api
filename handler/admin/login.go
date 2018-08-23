package admin

import "app/core"

var LoginPost core.HandlerFunc = func(c *core.Context) core.Response {

	// 用户验证


	return c.Success("验证成功!")

	return nil
}
