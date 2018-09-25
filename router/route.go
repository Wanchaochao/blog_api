package router

import (
	"blog/core"
	"blog/router/middleware"
	"blog/debug"
	"github.com/gin-gonic/gin"
	"blog/handler/admin"
	"github.com/gin-contrib/cors"
	"time"
)

var allowOrigins = map[string]bool{
	"http://admin.littlebug.vip":  true,
	"https://admin.littlebug.vip": true,
	"http://localhost:8080":       true,
}

func Route(router *gin.Engine) {
	//设置模板
	//core.SetTemplate(router)

	// 跨域
	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowOriginFunc: func(origin string) bool {
			return allowOrigins[origin]
		},
		AllowMethods: []string{"*"},
		AllowHeaders: []string{
			"Origin",
			"Content-Length",
			"Content-Type",
			"Access-Token",
			"Access-Control-Allow-Origin",
		},
		AllowCredentials: false,
		MaxAge:           12 * time.Hour,
	}))

	//中间件token验证
	router.Use(middleware.Ginrus())

	//登录
	router.POST("/login", core.Handle(admin.LoginPost))

	//后台
	blogAdmin := router.Group("/adm")
	blogAdmin.Use(core.Middware(middleware.Token))
	{
		// 用户信息
		blogAdmin.GET("/getUserInfo", core.Handle(admin.GetUserInfo))

		// 文章管理
		blogAdmin.POST("/articleList", core.Handle(admin.ArticleList))
		blogAdmin.POST("/storeArticle", core.Handle(admin.CreateArticle))
		blogAdmin.GET("/deleteArticle", core.Handle(admin.DeleteArticle))

		// 文章分类
		blogAdmin.GET("/categories", core.Handle(admin.Categories))
		//blogAdmin.GET("/deleteCategory",core.Handle(admin.DeleteCategory)) // 删
		blogAdmin.GET("/createCategory", core.Handle(admin.CreateCategory)) // 增

	}

	//debug handler
	debug.Route(router)
}
