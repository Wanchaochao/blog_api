package index

import (
	"blog/core"
	"blog/models"
	"database/sql"
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

	// 文章点赞点踩数
	praiseNum, err := gosql.Model(&models.Evaluate{}).Where("type = 1 and praise = 1 and foreign_key = ?", id).Count()
	if err != nil {
		return c.Fail(204, err)
	}
	againstNum, err := gosql.Model(&models.Evaluate{}).Where("type = 1 and praise = 0 and foreign_key = ?", id).Count()
	if err != nil {
		return c.Fail(204, err)
	}
	resp.ArticleEvaluate.PraiseNum = int(praiseNum)
	resp.ArticleEvaluate.AgainstNum = int(againstNum)

	return c.Success(resp)
}

var ArticlesPrevNext core.HandlerFunc = func(c *core.Context) core.Response {
	articleId := c.DefaultQuery("article_id", "")
	var resp = struct {
		Prev models.Articles `json:"prev"`
		Next models.Articles `json:"next"`
	}{}
	// 上一篇下一篇
	if err := gosql.Model(&resp.Prev).Where("id < ?", articleId).Get(); err != nil && err != sql.ErrNoRows {
		return c.Fail(205, err)
	}
	if err := gosql.Model(&resp.Next).Where("id > ?", articleId).Get(); err != nil && err != sql.ErrNoRows {
		return c.Fail(205, err)
	}
	return c.Success(resp)
}

// 文章的全部评论
var Comments core.HandlerFunc = func(c *core.Context) core.Response {
	articleId := c.DefaultQuery("article_id", "")
	comments := make([]*models.CommentsInfo, 0)
	gosql.Model(&comments).Where("article_id = ?", articleId).All()
	// 评论的点赞数点踩数
	for _, comment := range comments {
		if p, err := gosql.Model(&models.Evaluate{}).Where("type = 2 and praise = 1 and foreign_key = ?", comment.Id).Count(); err == nil {
			comment.PraiseNum = int(p)
		}
		if a, err := gosql.Model(&models.Evaluate{}).Where("type = 2 and praise = 0 and foreign_key = ?", comment.Id).Count(); err == nil {
			comment.AgainstNum = int(a)
		}
	}
	return c.Success(comments)
}
