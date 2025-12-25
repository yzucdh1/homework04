package main

import (
	"github.com/sirupsen/logrus"
	"github.com/yzucdh1/homework04/global"
	"github.com/yzucdh1/homework04/model"
	"log"
)

func main() {
	// 初始化日志
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.Info("开始数据库迁移...")

	// 初始化数据库连接
	err := global.InitConfig()
	if err != nil {
		log.Fatal("读取配置文件失败", err)
		return
	}

	err = global.Connect()
	if err != nil {
		log.Fatal("数据库连接失败", err)
		return
	}

	// 执行数据库迁移
	logrus.Info("运行数据库迁移...")

	err = global.DB.AutoMigrate(
		&model.User{},
		&model.Post{},
		&model.Comment{},
	)

	if err != nil {
		log.Fatal("数据库迁移失败", err)
	}

	logrus.Info("数据库迁移成功")

	// 显示创建的表信息
	logrus.Info("Created tables:")
	logrus.Info("- users (用户表)")
	logrus.Info("- posts (文章表)")
	logrus.Info("- comments (评论表)")

	// 检查表是否存在
	if global.DB.Migrator().HasTable(&model.User{}) {
		logrus.Info("✓ users table created successfully")
	}
	if global.DB.Migrator().HasTable(&model.Post{}) {
		logrus.Info("✓ posts table created successfully")
	}
	if global.DB.Migrator().HasTable(&model.Comment{}) {
		logrus.Info("✓ comments table created successfully")
	}
}
