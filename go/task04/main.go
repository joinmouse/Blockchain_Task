package main

import (
	"blog/config"
	"blog/models"
	"blog/routes"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
    // 加载配置文件
    cfg, err := config.LoadConfig("config/config.yaml")
    if err != nil {
        log.Fatal("加载配置文件失败:", err)
    }

    // 连接数据库
    db, err := gorm.Open(mysql.Open(cfg.GetDSN()), &gorm.Config{})
    if err != nil {
        log.Fatal("数据库连接失败:", err)
    }

    // 自动迁移数据库表
    db.AutoMigrate(&models.User{}, &models.Post{}, &models.Comment{})

    // 设置路由
    r := routes.SetupRouter(db, cfg.SecretKey)

    // 启动服务器
    addr := ":" + cfg.Server.Port
    if err := r.Run(addr); err != nil {
        log.Fatal("服务器启动失败:", err)
    }
}
