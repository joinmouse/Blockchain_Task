package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// ErrorHandler 中间件用于统一处理错误
func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next() // 继续处理请求

		// 检查是否有错误
		if len(c.Errors) > 0 {
			// 获取最后一个错误
			err := c.Errors.Last()
			var statusCode int

			// 根据错误类型设置状态码
			switch err.Type {
			case gin.ErrorTypePrivate:
				statusCode = http.StatusInternalServerError
			case gin.ErrorTypePublic:
				statusCode = http.StatusBadRequest
			default:
				statusCode = http.StatusInternalServerError
			}

			// 返回错误信息
			c.JSON(statusCode, gin.H{"error": err.Error()})
		}
	}
}
