package handler

import "github.com/gin-gonic/gin"

var CommentGroup *gin.RouterGroup

func CommentHandler() {
	CommentGroup.Use(JWTTokenMiddleware)
}
