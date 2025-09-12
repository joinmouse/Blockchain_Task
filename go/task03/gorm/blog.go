package main

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// User 用户模型
type User struct {
	ID        uint   `gorm:"primaryKey"`
	Name      string `gorm:"size:100"`
	Posts     []Post  `gorm:"foreignKey:UserID"` // 一对多关系
	PostCount int    `gorm:"default:0"`          // 文章数量统计字段
}

// Post 文章模型
type Post struct {
	ID       uint      `gorm:"primaryKey"`
	Title    string    `gorm:"size:100"`
	Body     string    `gorm:"type:text"`
	UserID   uint      // 外键
	Comments []Comment  `gorm:"foreignKey:PostID"` // 一对多关系
	Status    string    `gorm:"default:'有评论'"`    // 评论状态
}

// Comment 评论模型
type Comment struct {
	ID     uint   `gorm:"primaryKey"`
	Body   string `gorm:"type:text"`
	PostID uint   // 外键
}

// 钩子函数：在 Post 创建时更新用户的文章数量
func (post *Post) AfterCreate(tx *gorm.DB) (err error) {
	var user User
	if err := tx.Model(&User{}).Where("id = ?", post.UserID).First(&user).Error; err != nil {
		return err
	}
	user.PostCount++
	return tx.Save(&user).Error
}

// 钩子函数：在 Comment 删除时检查文章的评论数量
func (comment *Comment) AfterDelete(tx *gorm.DB) (err error) {
	var post Post
	if err := tx.Model(&Post{}).Where("id = ?", comment.PostID).First(&post).Error; err != nil {
		return err
	}

	var count int64
	tx.Model(&Comment{}).Where("post_id = ?", comment.PostID).Count(&count)

	if count == 0 {
		post.Status = "无评论"
		return tx.Save(&post).Error
	}
	return nil
}

func main() {
	// 数据库连接
	dsn := "root:12345abc@tcp(127.0.0.1:3306)/gorm_blog?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln("连接数据库失败:", err)
	}

	// 创建表
	err = db.AutoMigrate(&User{}, &Post{}, &Comment{})
	if err != nil {
		log.Fatalln("创建表失败:", err)
	}

	// // 插入一些用户数据
	// user1 := User{Name: "Alice"}
	// user2 := User{Name: "Bob"}
	// db.Create(&user1)
	// db.Create(&user2)

	// 插入一些文章数据
	post1 := Post{Title: "Go 语言入门", Body: "这是一本关于 Go 语言的入门书籍", UserID: 1}
	db.Create(&post1)

	// 插入一些评论数据
	comment1 := Comment{Body: "这本书很不错！", PostID: post1.ID}
	db.Create(&comment1)

	fmt.Println("初始数据插入成功")

	// 查询用户的文章数量
	var updatedUser User
	db.First(&updatedUser, 1)
	fmt.Printf("用户 %s 的文章数量: %d\n", updatedUser.Name, updatedUser.PostCount)

	// 删除评论
	db.Delete(&comment1)

	// 查询文章状态
	var updatedPost Post
	db.First(&updatedPost, post1.ID)
	fmt.Printf("文章状态: %s\n", updatedPost.Status)
}
