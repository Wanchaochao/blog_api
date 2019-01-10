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
	return c.Success(articleResp, "ok")
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

	if err := gosql.Model(&resp.Prev).Where("id < ?", id).Get(); err != nil && err != sql.ErrNoRows {
		return c.Fail(205, err)
	}
	if err := gosql.Model(&resp.Next).Where("id > ?", id).Get(); err != nil && err != sql.ErrNoRows {
		return c.Fail(205, err)
	}

	return c.Success(resp, "ok")
}

// 文章的全部评论
var Comments core.HandlerFunc = func(c *core.Context) core.Response {
	articleId := c.DefaultQuery("article_id", "")
	comments := make([]*models.CommentsInfo, 0)
	gosql.Model(&comments).Where("article_id = ?", articleId).OrderBy("created_at desc").All()
	// 评论的点赞数点踩数
	for _, comment := range comments {
		if p, err := gosql.Model(&models.Evaluate{}).Where("type = 2 and praise = 1 and foreign_key = ?", comment.Id).Count(); err == nil {
			comment.PraiseNum = int(p)
		}
		if a, err := gosql.Model(&models.Evaluate{}).Where("type = 2 and praise = 0 and foreign_key = ?", comment.Id).Count(); err == nil {
			comment.AgainstNum = int(a)
		}
	}
	return c.Success(comments, "ok")
}

// 评论文章
var CreateComments core.HandlerFunc = func(c *core.Context) core.Response {

	comments := &models.Comments{}
	if err := c.ShouldBindJSON(comments); err != nil {
		return c.Fail(201, err)
	}
	pk, err := gosql.Model(comments).Create()
	if err != nil {
		return c.Fail(202, err)
	}
	comment := &models.Comments{}
	if err := gosql.Model(comment).Where("id = ?", pk).Get(); err != nil && nil != sql.ErrNoRows {
		return c.Fail(203, err)
	}
	return c.Success(comment, "ok")
}

// 根据分类获取分类下所有文章
var Articles core.HandlerFunc = func(c *core.Context) core.Response {

	req := struct {
		CategoryId int `json:"category_id"`
	}{}

	if err := c.ShouldBindJSON(&req); err != nil {
		return c.Fail(201, err)
	}

	type resp = struct {
		models.Articles
		Like        int `json:"like"`
		Dislike     int `json:"dislike"`
		CommentsNum int `json:"comments_num"`
	}

	articles := make([]*resp, 0)

	if err := gosql.Model(&articles).Where("category_id = ?", req.CategoryId).All(); err != nil {
		return c.Fail(202, err)
	}

	// 文章的点踩点赞数
	for _, value := range articles {
		rows, err := gosql.Queryx("select count(*) num, praise from evaluate where foreign_key = ? and type = 1 group by praise", value.Articles.Id)
		if err != nil {
			return c.Fail(203, err)
		}

		for rows.Next() {
			var num int
			var praise int
			rows.Scan(&num, &praise)
			if praise == 0 {
				value.Dislike = num
			} else {
				value.Like = num
			}
		}
		// 文章的评论数
		num, err := gosql.Model(&models.Comments{}).Where("article_id = ?", value.Articles.Id).Count()
		if err != nil {
			return c.Fail(204, err)
		}
		value.CommentsNum = int(num)
	}
	return c.Success(articles, "ok")
}
