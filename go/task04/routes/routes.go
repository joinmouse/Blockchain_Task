package routes

import (
	"blog/controllers"
	"blog/middleware"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// SetupRouter 设置路由
func SetupRouter(db *gorm.DB, secretKey string) *gin.Engine {
	r := gin.Default()

    // 使用错误处理的中间件
    r.Use(middleware.ErrorHandler())

	// 初始化控制器
	userController := controllers.NewUserController(db, secretKey) // 传递配置
	postController := &controllers.PostController{DB: db}
    commentController := &controllers.CommentController{DB: db}

	// 用户相关路由
	r.POST("/register", userController.Register)
	r.POST("/login", userController.Login)

	// 文章相关路由，使用 JWT 认证中间件
	authorized := r.Group("/api")
	authorized.Use(middleware.AuthMiddleware(secretKey)) // 添加中间件
	{
		authorized.POST("/posts", postController.CreatePost)
		authorized.GET("/posts", postController.GetPosts)
		authorized.GET("/posts/:id", postController.GetPost)
		authorized.PUT("/posts/:id", postController.UpdatePost)
		authorized.DELETE("/posts/:id", postController.DeletePost)

		// 评论相关路由
		authorized.POST("/posts/:id/comments", commentController.CreateComment) // 创建评论
		authorized.GET("/posts/:id/comments", commentController.GetComments)    // 获取评论列表
	}
	return r
} 
