package models

import (
	"gorm.io/gorm"
)

// User 用户模型
type User struct {
    gorm.Model
    Username string `gorm:"unique;not null"` // 用户名
    Password string `gorm:"not null"`        // 密码
    Email    string `gorm:"unique;not null"` // 邮箱
    Posts    []Post `gorm:"foreignKey:UserID"` // 用户发布的文章
} 
