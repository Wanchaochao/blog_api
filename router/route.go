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

	//后台
	blogAdmin := router.Group("/admin")
	blogAdmin.Use(core.Middware(middleware.Token))
	{
		blogAdmin.POST("/login",core.Handle(admin.LoginPost))
	}

	//debug handler
	debug.Route(router)
}
