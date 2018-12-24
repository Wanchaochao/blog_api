package index

import (
	"blog/core"
	"blog/models"
	"github.com/ilibs/gosql"
)

type Resp struct {
	models.Category
	Articles []*models.Articles `json:"articles" db:"-" relation:"id,category_id"`
}

var Categories core.HandlerFunc = func(c *core.Context) core.Response {

	categories := make([]*Resp, 0)

	if err := gosql.Model(&categories).All(); err != nil {
		return c.Fail(301, err)
	}
	return c.Success(categories)
}
