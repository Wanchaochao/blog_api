package admin

import (
	"blog/core"
	"blog/models"
	"github.com/ilibs/gosql"
)

type ArticleJson struct {
	Id int `json:"id"`
}

// 获取单个文章
var Article core.HandlerFunc = func(c *core.Context) core.Response {
	article := &models.Articles{}
	request := &ArticleJson{}
	if err := c.ShouldBindJSON(request); err != nil {
		return c.Fail(202, err)
	}
	if err := gosql.Model(article).Where("id = ?", request.Id).Get(); err != nil {
		return c.Fail(203, err)
	}
	return c.Success(article)
}

// 更新文章
var UpdateArticle core.HandlerFunc = func(c *core.Context) core.Response {
	article := &models.Articles{}
	if err := c.ShouldBindJSON(article); err != nil {
		return c.Fail(202, err)
	}
	if article.CategoryId == "" || article.Title == "" || article.Description == "" || article.Author == "" {
		return c.Fail(202, "params are not permitted")
	}
	if len(article.Content) < 50 {
		return c.Fail(204, "the content of article is too shot")
	}
	cate := &models.Category{}

	if err := gosql.Model(cate).Where("id = ?", article.CategoryId).Get(); err != nil {
		return c.Fail(203, "the article category is not existed!")
	}

	if _, err := gosql.Model(article).Update(); err != nil {
		return c.Fail(205, "update article failed!")
	}
	return c.Success("update successfully!")
}

type PageJson struct {
	Page int `json:"page"`
}

// 文章列表
var ArticleList core.HandlerFunc = func(c *core.Context) core.Response {
	article := &models.Articles{}
	page := &PageJson{}
	err := c.ShouldBindJSON(page)
	if err != nil {
		return c.Fail(202, err)
	}
	keyword := c.DefaultQuery("keyword", "")
	startTime := c.DefaultQuery("start_time", "")
	endTime := c.DefaultQuery("start_time", "")
	articleResp, err := models.GetArticleList(article, page.Page, 10, keyword, startTime, endTime)
	if err != nil {
		return c.Fail(203, err)
	}
	articleResp.Current = page.Page
	return c.Success(articleResp)
}

// 创建文章
var CreateArticle core.HandlerFunc = func(c *core.Context) core.Response {

	article := &models.Articles{}

	if err := c.ShouldBindJSON(article); err != nil {
		return c.Fail(301, err)
	}
	if article.CategoryId == "" || article.Title == "" || article.Description == "" || article.Author == "" {
		return c.Fail(202, "params are not permitted")
	}
	if len(article.Content) < 50 {
		return c.Fail(204, "the content of article is too shot")
	}
	cate := &models.Category{}
	err := gosql.Model(cate).Where("id = ?", article.CategoryId).Get()
	if err != nil {
		return c.Fail(203, "the article category is not existed!")
	}

	if _, err := gosql.Model(article).Create(); err != nil {
		return c.Fail(204, err)
	}
	return c.Success(nil)
}

// 删除文章
var DeleteArticle core.HandlerFunc = func(c *core.Context) core.Response {
	article := &models.Articles{}
	id := c.DefaultQuery("id", "")
	if id == "" {
		return c.Fail(203, "missing param article id")
	}
	if _, err := gosql.Model(article).Where("id = ?", id).Delete(); err != nil {
		return c.Fail(301, err)
	}
	return c.Success("删除成功")
}
