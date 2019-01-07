package router

import (
	"blog/core"
	"blog/debug"
	"blog/handler/admin"
	"blog/handler/index"
	"blog/router/middleware"
	"github.com/gin-gonic/gin"
)

//var allowOrigins = map[string]bool{
//	"http://admin.littlebug.vip":  true,
//	"https://admin.littlebug.vip": true,
//	"http://localhost:8080":       true,
//}

func Route(router *gin.Engine) {
	//中间件token验证
	router.Use(middleware.Ginrus())
	//登录
	router.POST("adm/login", core.Handle(admin.LoginPost))
	// 滑块验证码
	router.GET("adm/captcha", core.Handle(admin.Captcha))

	//后台
	blogAdmin := router.Group("/adm")
	blogAdmin.Use(core.Middleware(middleware.Token))
	{
		// 用户信息
		//blogAdmin.GET("/getUserInfo", core.Handle(admin.GetUserInfo))

		// 文章管理
		blogAdmin.POST("/articleList", core.Handle(admin.ArticleList))     // 文章列表
		blogAdmin.POST("/article", core.Handle(admin.Article))             // 单个文章
		blogAdmin.POST("/storeArticle", core.Handle(admin.CreateArticle))  // 创建文章
		blogAdmin.GET("/deleteArticle", core.Handle(admin.DeleteArticle))  // 删除文章
		blogAdmin.POST("/updateArticle", core.Handle(admin.UpdateArticle)) // 更新文章

		// 文章分类
		blogAdmin.GET("/categories", core.Handle(admin.Categories))
		//blogAdmin.GET("/deleteCategory",core.Handle(admin.DeleteCategory)) // 删
		blogAdmin.GET("/createCategory", core.Handle(admin.CreateCategory)) //
	}

	// 前台
	indexApi := router.Group("/api")
	{
		// 文章
		indexApi.POST("/articleList", core.Handle(index.ArticleList))      // 文章列表
		indexApi.POST("/articles", core.Handle(index.Articles))            // 某个分类下的所有文章
		indexApi.GET("/article", core.Handle(index.Article))               //单个文章
		indexApi.GET("/comments", core.Handle(index.Comments))             //文章的所有评论
		indexApi.POST("/createComment", core.Handle(index.CreateComments)) //评论
		indexApi.GET("/categories", core.Handle(index.Categories))         // 分类列表
		indexApi.GET("/evaluate", core.Handle(index.Evaluate))             // 点赞点踩

	}

	//debug handler
	debug.Route(router)
}
