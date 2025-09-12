package controllers

import (
	"blog/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CommentController struct {
    DB *gorm.DB
}

// CreateComment 创建评论
func (pc *CommentController) CreateComment(c *gin.Context) {
    var comment models.Comment
    if err := c.ShouldBindJSON(&comment); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // 从上下文中获取用户ID
    userID, exists := c.Get("userID")
    if !exists {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "未认证用户"})
        return
    }
    comment.UserID = userID.(uint) // 设置评论者ID

    // 获取文章ID
    postID := c.Param("id")
	id, err := strconv.Atoi(postID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
	}
    comment.PostID = uint(id)

    // 创建评论
    if err := pc.DB.Create(&comment).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "评论创建失败"})
        return
    }

    c.JSON(http.StatusCreated, gin.H{"message": "评论创建成功", "comment": comment})
}

// GetComments 获取文章的所有评论
func (pc *CommentController) GetComments(c *gin.Context) {
    postID := c.Param("id")
    var comments []models.Comment

    if err := pc.DB.Where("post_id = ?", postID).Find(&comments).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "获取评论列表失败"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"comments": comments})
}
