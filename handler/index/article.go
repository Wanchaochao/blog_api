package index

import (
	"blog/core"
	"blog/models"
	"github.com/ilibs/gosql"
)

type ListRequest struct {
	Page      int    `json:"page"`
	Keyword   string `json:"keyword"`
	StartTime string `json:"start_time"`
	EndTime   string `json:"end_time"`
}

// 文章列表
var ArticleList core.HandlerFunc = func(c *core.Context) core.Response {
	article := &models.Articles{}
	request := &ListRequest{}
	err := c.ShouldBindJSON(request)
	if err != nil {
		return c.Fail(202, err)
	}
	articleResp, err := models.GetArticleList(article, request.Page, 10, request.Keyword, request.StartTime, request.EndTime)
	if err != nil {
		return c.Fail(203, err)
	}
	return c.Success(articleResp)
}

// 获取单个文章
var Article core.HandlerFunc = func(c *core.Context) core.Response {
	id := c.DefaultQuery("id", "")
	if id == "" {
		return c.Fail(202, "缺少参数")
	}

	resp := &models.ArticleInfo{}

	if err := gosql.Model(resp).Where("id = ?", id).Get(); err != nil {
		return c.Fail(203, err)
	}
	praiseNum, err := gosql.Model(&models.Evaluate{}).Where("type = 1 and praise = 1 and foreign_key = ?", id).Count()
	if err != nil {
		return c.Fail(204, err)
	}
	againstNum, err := gosql.Model(&models.Evaluate{}).Where("type = 1 and praise = 0 and foreign_key = ?", id).Count()
	if err != nil {
		return c.Fail(204, err)
	}
	resp.ArticleEvaluate.Praise = int(praiseNum)
	resp.ArticleEvaluate.Against = int(againstNum)

	return c.Success(resp)
}
