package admin

import (
	"blog/core"
	"blog/models"
	"github.com/ilibs/gosql"
	"strings"
)

var Categories core.HandlerFunc = func(c *core.Context) core.Response {
	categories := make([]*models.Category,0)
	rows,err := gosql.Queryx("select * from category")
	if err != nil {
		return c.Fail(301,err)
	}
	for rows.Next()  {
		v := &models.Category{}
		err := rows.StructScan(v)
		if err != nil {
			return c.Fail(302,err)
		}
		categories = append(categories,v)
	}
	return c.Success(categories)
}

var CreateCategory core.HandlerFunc = func(c *core.Context) core.Response {
	name := strings.TrimSpace(c.DefaultQuery("name",""))
	if name == "" {
		return c.Fail(201,"name can not be empty")
	}
	cate := &models.Category{}
	cate.Name = name
	_,err := gosql.Model(cate).Create()
	if err != nil {
		return c.Fail(301,err)
	}
	return c.Success("创建成功!")
}

//var DeleteCategory core.HandlerFunc = func(c *core.Context) core.Response {
//
//	return nil
//}