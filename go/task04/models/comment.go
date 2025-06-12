package models

import (
	"gorm.io/gorm"
)

// Comment 评论模型
type Comment struct {
	gorm.Model          // 嵌入 gorm.Model，包含 ID、CreatedAt、UpdatedAt 和 DeletedAt 字段
	Content string      `gorm:"not null"` // 评论内容，不能为空
	UserID  uint       // 评论者的用户ID，外键关联到 User 表
	User    User       `gorm:"foreignKey:UserID"` // 关联的用户模型
	PostID  uint       // 文章ID，外键关联到 Post 表
	Post    Post       `gorm:"foreignKey:PostID"` // 关联的文章模型
}
