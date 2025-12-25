package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/yzucdh1/homework04/global"
	"github.com/yzucdh1/homework04/handler"
)

func main() {
	err := global.InitConfig()
	if err != nil {
		fmt.Println("读取配置文件失败", err)
		return
	}

	err = global.Connect()
	if err != nil {
		fmt.Println("数据库连接失败", err)
		return
	}

	//err = global.DB.AutoMigrate(&model.User{}, &model.Post{}, &model.Comment{})
	//if err != nil {
	//	fmt.Println("表迁移失败", err)
	//	return
	//}

	// 设置日志模式
	gin.SetMode(gin.DebugMode)
	// 创建一个默认的路由
	router := gin.Default()

	// 路由分组
	handler.UserGroup = router.Group("/api/user")
	handler.PostGroup = router.Group("/api/post")
	handler.CommentGroup = router.Group("/api/comment")
	handler.UserHandler()
	handler.PostHandler()
	handler.CommentHandler()

	// 启动监听，gin会把web服务运行在本机的0.0.0.0:8080端口上
	err = router.Run(fmt.Sprintf(":%d", global.Cfg.Server.Port))
	if err != nil {
		fmt.Println("服务启动失败", err)
	}
}
