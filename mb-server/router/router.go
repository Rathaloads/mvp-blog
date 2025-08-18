package router

import (
	"mb-server/controller/admin"
	"mb-server/middleware"

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
		blog := api.Group("/blog").Use(middleware.AuthMiddleware)
		{
			blog.GET("/getBlogList", func(ctx *gin.Context) {
				v, ok := ctx.Get("Email")
				if !ok {
					ctx.JSON(200, gin.H{"message": "no exist"})
				}
				ctx.JSON(200, gin.H{"message": v})
			})
			blog.GET("/getBlog/:id", func(ctx *gin.Context) { ctx.JSON(200, gin.H{"message": "111111"}) })
			blog.POST("/addBlog", func(ctx *gin.Context) { ctx.JSON(200, gin.H{"message": "111111"}) })
			blog.POST("/updateBlog/:id", func(ctx *gin.Context) { ctx.JSON(200, gin.H{"message": "111111"}) })
			blog.POST("/removeBlog/:id", func(ctx *gin.Context) { ctx.JSON(200, gin.H{"message": "111111"}) })
		}

		// 说说
		moment := api.Group("/moment").Use(middleware.AuthMiddleware)
		{
			moment.GET("/getMomentList", func(ctx *gin.Context) { ctx.JSON(200, gin.H{"message": "111111"}) })
			moment.POST("/addMoment", func(ctx *gin.Context) { ctx.JSON(200, gin.H{"message": "111111"}) })
			moment.POST("/updateMoment", func(ctx *gin.Context) { ctx.JSON(200, gin.H{"message": "111111"}) })
			moment.POST("/removeMomnet", func(ctx *gin.Context) { ctx.JSON(200, gin.H{"message": "111111"}) })
		}

		// 评论
		comment := api.Group("comment").Use(middleware.AuthMiddleware)
		{
			comment.GET("/getComment", func(ctx *gin.Context) { ctx.JSON(200, gin.H{"message": "111111"}) })
			comment.POST("/updateMoment", func(ctx *gin.Context) { ctx.JSON(200, gin.H{"message": "111111"}) })
			comment.POST("/removeMomnet", func(ctx *gin.Context) { ctx.JSON(200, gin.H{"message": "111111"}) })
		}

	}
	return router
}
