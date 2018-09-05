package router

import (
	"blog/core"
	"blog/router/middleware"
	"blog/debug"
	"github.com/gin-gonic/gin"
	"blog/handler/admin"
)

func Route(router *gin.Engine) {
	//设置模板
	//core.SetTemplate(router)

	//中间件token验证
	router.Use(middleware.Ginrus())

	//登录
	router.POST("/login",core.Handle(admin.LoginPost))

	//后台
	blogAdmin := router.Group("/adm")
	blogAdmin.Use(core.Middware(middleware.Token))
	{
		// 用户信息
		blogAdmin.GET("/getUserInfo",core.Handle(admin.GetUserInfo))

		// 文章管理
		blogAdmin.POST("/articleList",core.Handle(admin.ArticleList))
		blogAdmin.POST("/storeArticle",core.Handle(admin.CreateArticle))
		blogAdmin.GET("/deleteArticle",core.Handle(admin.DeleteArticle))

		// 文章分类
		blogAdmin.GET("/categories",core.Handle(admin.Categories))
		//blogAdmin.GET("/deleteCategory",core.Handle(admin.DeleteCategory)) // 删
		blogAdmin.GET("/createCategory",core.Handle(admin.CreateCategory)) // 增

	}

	//debug handler
	debug.Route(router)
}
