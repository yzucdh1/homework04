package global

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
	"path/filepath"
)

var Cfg Config

// Config 总配置结构体（与 config.yaml 结构对应）
type Config struct {
	Database DatabaseConfig `yaml:"database"` // 数据库配置（对应 yaml 中的 database 节点）
	Server   ServerConfig   `yaml:"server"`   // 服务配置（可选）
}

// DatabaseConfig 数据库配置结构体
type DatabaseConfig struct {
	Driver          string `yaml:"driver"` // 数据库驱动
	DSN             string `yaml:"dsn"`    // 连接串
	MaxOpenConns    int    `yaml:"max_open_conns"`
	MaxIdleConns    int    `yaml:"max_idle_conns"`
	ConnMaxLifetime int    `yaml:"conn_max_lifetime"`
	ConnMaxIdleTime int    `yaml:"conn_max_idle_time"`
}

// ServerConfig 服务配置结构体（可选）
type ServerConfig struct {
	Port int `yaml:"port"`
}

// InitConfig 初始化配置：读取并解析 config.yaml
func InitConfig() error {
	// 获取配置文件路径（项目根目录下的 config.yaml）
	configPath := filepath.Join(".", "config.yaml")

	// 检查文件是否存在
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		return fmt.Errorf("配置文件不存在：%s", configPath)
	}

	// 读取文件内容
	fileContent, err := os.ReadFile(configPath)
	if err != nil {
		return fmt.Errorf("读取配置文件失败：%w", err)
	}

	// YAML 解析为结构体
	if err := yaml.Unmarshal(fileContent, &Cfg); err != nil {
		return fmt.Errorf("解析配置文件失败：%w", err)
	}

	return nil
}
