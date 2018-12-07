package admin

import (
	"blog/core"
	"blog/models"
	"crypto/md5"
	"encoding/hex"
	"github.com/ilibs/gosql"
)

// 用户登录验证
type LoginUserParams struct {
	Name     string `form:"name" json:"name" xml:"name"  binding:"required"`
	Password string `form:"password" json:"password" xml:"password"  binding:"required"`
}

var LoginPost core.HandlerFunc = func(c *core.Context) core.Response {
	params := &LoginUserParams{}
	if err := c.ShouldBindJSON(params); err != nil {
		return c.Fail(202, err)
	}

	// 查用户
	user := &models.Users{}
	err := gosql.Model(user).Where("name = ?", params.Name).Get()
	if err != nil {
		return c.Fail(204, err)
	}
	// 密码验证
	md5 := md5.New()
	md5.Write([]byte(params.Password))
	if md5str := hex.EncodeToString(md5.Sum(nil)); md5str != user.Password {
		return c.Fail(401, "密码错误!")
	}

	return c.Success(user)

	return nil
}

// 获取用户信息
type GetUserParams struct {
	Name  string `form:"userName"`
	Token string `form:"token"`
}

//var GetUserInfo core.HandlerFunc = func(c *core.Context) core.Response {
//	params := &GetUserParams{}
//	if err := c.Bind(params); err != nil {
//		return c.Fail(202, err)
//	}
//
//	if params.Name == "" || params.Token == "" {
//		return c.Fail(203, "missing required params token or Name")
//	}
//
//	gosql.Connect(config.App.Db)
//	user := &models.Users{}
//	err := gosql.Model(user).Where("name = ? and token = ?",params.Name, params.Token).Get()
//	if err != nil {
//		return c.Fail(204,err)
//	}
//	return c.Success(user)
//}
