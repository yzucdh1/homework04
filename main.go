package main

import (
	"fmt"
	"github.com/yzucdh1/homework04/global"
	"github.com/yzucdh1/homework04/routes"
	"log"
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

	router := routes.SetRoutes()

	// 启动监听，gin会把web服务运行在本机的0.0.0.0:8080端口上
	err = router.Run(fmt.Sprintf(":%d", global.Cfg.Server.Port))
	if err != nil {
		log.Fatal("服务启动失败", err)
	}
}
