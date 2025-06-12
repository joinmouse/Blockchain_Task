package models

import (
	"gorm.io/gorm"
)

// Post 文章模型
type Post struct {
    gorm.Model
    Title    string    `gorm:"not null"` // 文章标题
    Content  string    `gorm:"type:text"` // 文章内容
    UserID   uint      `gorm:"not null"` // 作者ID
    User     User      `gorm:"foreignKey:UserID"` // 关联用户
    Comments []Comment `gorm:"foreignKey:PostID"` // 文章评论
} 
