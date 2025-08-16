package router

import (
	"mb-server/controller/admin"

	"github.com/gin-gonic/gin"
)

func StartRouter() *gin.Engine {
	router := gin.Default()
	api := router.Group("/api")
	{
		// 权限
		auth := api.Group("/auth")
		{
			auth.POST("/login", admin.AdminLogin)
			auth.POST("/register", admin.AdminRegister)
		}

		// 日志
		blog := api.Group("/blog")
		{
			blog.GET("/getBlogList", func(ctx *gin.Context) {})
			blog.GET("/getBlog/:id", func(ctx *gin.Context) {})
			blog.POST("/addBlog", func(ctx *gin.Context) {})
			blog.POST("/updateBlog/:id", func(ctx *gin.Context) {})
			blog.POST("/removeBlog/:id", func(ctx *gin.Context) {})
		}

		// 说说
		moment := api.Group("/moment")
		{
			moment.GET("/getMomentList", func(ctx *gin.Context) {})
			moment.POST("/addMoment", func(ctx *gin.Context) {})
			moment.POST("/updateMoment", func(ctx *gin.Context) {})
			moment.POST("/removeMomnet", func(ctx *gin.Context) {})
		}

		// 评论
		comment := api.Group("comment")
		{
			comment.GET("/getComment", func(ctx *gin.Context) {})
			comment.POST("/updateMoment", func(ctx *gin.Context) {})
			comment.POST("/removeMomnet", func(ctx *gin.Context) {})
		}

	}
	return router
}
