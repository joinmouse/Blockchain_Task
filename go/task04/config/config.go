package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

// Config 应用配置结构体
type Config struct {
    Database DatabaseConfig `yaml:"database"`
    Server   ServerConfig   `yaml:"server"`
	SecretKey string        `yaml:"secret_key"` // 密钥字段
}

// DatabaseConfig 数据库配置
type DatabaseConfig struct {
    Driver    string `yaml:"driver"`
    Host      string `yaml:"host"`
    Port      string `yaml:"port"`
    User      string `yaml:"user"`
    Password  string `yaml:"password"`
    DBName    string `yaml:"dbname"`
    Charset   string `yaml:"charset"`
    ParseTime bool   `yaml:"parseTime"`
    Loc       string `yaml:"loc"`
}

// ServerConfig 服务器配置
type ServerConfig struct {
    Port string `yaml:"port"`
    Mode string `yaml:"mode"`
}

// LoadConfig 从文件加载配置
func LoadConfig(filename string) (*Config, error) {
    buf, err := os.ReadFile(filename)
    if err != nil {
        return nil, fmt.Errorf("读取配置文件失败: %v", err)
    }

    config := &Config{}
    err = yaml.Unmarshal(buf, config)
    if err != nil {
        return nil, fmt.Errorf("解析配置文件失败: %v", err)
    }

    return config, nil
}

// GetDSN 获取数据库连接字符串
func (c *Config) GetDSN() string {
    return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=%v&loc=%s",
        c.Database.User,
        c.Database.Password,
        c.Database.Host,
        c.Database.Port,
        c.Database.DBName,
        c.Database.Charset,
        c.Database.ParseTime,
        c.Database.Loc,
    )
} 
