package global

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"time"
)

var DB *gorm.DB

func Connect() error {
	var err error
	cfg := Cfg.Database // 获取解析后的数据库配置

	// 根据驱动类型选择对应的 Gorm 驱动
	switch cfg.Driver {
	case "mysql":
		// 连接 MySQL
		DB, err = gorm.Open(mysql.Open(cfg.DSN), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Info), // 打印SQL日志
		})
	default:
		return fmt.Errorf("不支持的数据库驱动：%s", cfg.Driver)
	}

	if err != nil {
		return fmt.Errorf("数据库连接失败：%w", err)
	}

	// 配置连接池（从配置文件读取参数）
	sqlDB, err := DB.DB()
	if err != nil {
		return fmt.Errorf("获取 SQL DB 实例失败：%w", err)
	}

	// 设置连接池参数
	sqlDB.SetMaxOpenConns(cfg.MaxOpenConns)                                    // 最大打开连接数
	sqlDB.SetMaxIdleConns(cfg.MaxIdleConns)                                    // 最大空闲连接数
	sqlDB.SetConnMaxLifetime(time.Duration(cfg.ConnMaxLifetime) * time.Second) // 连接最大生命周期
	sqlDB.SetConnMaxIdleTime(time.Duration(cfg.ConnMaxIdleTime) * time.Second) // 连接最大空闲时间

	return nil
}
