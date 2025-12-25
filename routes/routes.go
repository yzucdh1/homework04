package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/yzucdh1/homework04/controller"
	"github.com/yzucdh1/homework04/middleware"
)

func SetRoutes() *gin.Engine {
	router := gin.Default()

	router.Use(middleware.LoggerMiddleware())
	//router.Use(gin.Recovery())

	userController := controller.UserController{}
	postController := controller.PostController{}
	commentController := controller.CommentController{}

	// 路由分组
	api := router.Group("/api/v1")
	{
		// 不需要认证
		userGroup := api.Group("/user")
		{
			userGroup.POST("/register", userController.Register)
			userGroup.POST("/login", userController.Login)
		}

		// 需要认证
		authGroup := api.Group("")
		authGroup.Use(middleware.JWTTokenMiddleware())
		{
			authGroup.POST("/post/create", postController.PostCreate)
			authGroup.POST("/post/list", postController.PostList)
			authGroup.GET("/post/detail/:id", postController.PostDetail)
			authGroup.POST("/post/update", postController.PostUpdate)
			authGroup.GET("/post/delete/:id", postController.PostDelete)
			authGroup.POST("/comment/create", commentController.CommentCreate)
			authGroup.GET("/comment/list/:pId", commentController.CommentList)
		}
	}

	return router
}
