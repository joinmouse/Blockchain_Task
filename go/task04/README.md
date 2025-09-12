# 个人博客系统

这是一个使用 Go 语言、Gin 框架和 GORM 库开发的个人博客系统后端，支持用户认证、文章管理和评论功能。

## 运行环境

- Go 版本：1.21.0 或更高
- 数据库：MySQL 5.7 或更高
- Gin 框架
- GORM 库

## 依赖安装步骤


1. **初始化 Go 模块**（如果尚未初始化）：
   ```bash
   go mod init <your-module-name>
   ```

2. **安装依赖**：
   ```bash
   go get -u github.com/gin-gonic/gin
   go get -u gorm.io/gorm
   go get -u gorm.io/driver/mysql
   go get -u gopkg.in/yaml.v2
   ```

## 启动方式

1. **配置数据库**：
   在 `config/config.yaml` 文件中配置数据库连接信息，例如：
   ```yaml
   database:
     driver: mysql
     host: 127.0.0.1
     port: 3306
     user: root
     password: 12345abc
     dbname: blog
     charset: utf8mb4
     parseTime: true
     loc: Local

   server:
     port: 8080
     mode: debug  # debug or release

   secret_key: "your_generated_secret_key"  # 在这里放置您的密钥
   ```

2. **运行项目**：
   在项目根目录下运行以下命令：
   ```bash
   go run main.go
   ```

## 功能

- 用户注册和登录
- 文章的创建、读取、更新和删除（CRUD）
- 评论的创建和读取

## 贡献

欢迎任何形式的贡献！请提交问题或拉取请求。

## 许可证

本项目采用 MIT 许可证，详细信息请参见 LICENSE 文件。
