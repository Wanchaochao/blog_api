package index

import (
	"blog_api/core"
	"blog_api/models"
	"github.com/ilibs/gosql"
)

type Resp struct {
	models.Category
	ArticleNum int
}

var Categories core.HandlerFunc = func(c *core.Context) core.Response {

	categories := make([]*Resp, 0)
	if err := gosql.Model(&categories).All(); err != nil {
		return c.Fail(301, err)
	}
	for _, category := range categories {
		if n, err := gosql.Model(&models.Articles{}).Where("category_id = ?", category.Id).Count(); err == nil {
			category.ArticleNum = int(n)
		}
	}
	return c.Success(categories, "ok")
}
