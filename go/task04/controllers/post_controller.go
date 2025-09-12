package controllers

import (
	"blog/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type PostController struct {
    DB *gorm.DB
}

// CreatePost 创建文章
func (pc *PostController) CreatePost(c *gin.Context) {
    var post models.Post
    if err := c.ShouldBindJSON(&post); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := pc.DB.Create(&post).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "文章创建失败"})
        return
    }

    c.JSON(http.StatusCreated, gin.H{"message": "文章创建成功", "post": post})
}

// GetPosts 获取文章列表
func (pc *PostController) GetPosts(c *gin.Context) {
    var posts []models.Post
    if err := pc.DB.Preload("User").Find(&posts).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "获取文章列表失败"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"posts": posts})
}

// GetPost 获取单篇文章
func (pc *PostController) GetPost(c *gin.Context) {
    id := c.Param("id")
    var post models.Post

    if err := pc.DB.Preload("User").Preload("Comments").First(&post, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "文章不存在"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"post": post})
}

// UpdatePost 更新文章
func (pc *PostController) UpdatePost(c *gin.Context) {
    id := c.Param("id")
    var post models.Post

    if err := pc.DB.First(&post, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "文章不存在"})
        return
    }

    if err := c.ShouldBindJSON(&post); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := pc.DB.Save(&post).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "文章更新失败"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "文章更新成功", "post": post})
}

// DeletePost 删除文章
func (pc *PostController) DeletePost(c *gin.Context) {
    id := c.Param("id")
    if err := pc.DB.Delete(&models.Post{}, id).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "文章删除失败"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "文章删除成功"})
} 
