package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

// AuthMiddleware JWT 认证中间件
func AuthMiddleware(secretKey string) gin.HandlerFunc {
    return func(c *gin.Context) {
        tokenString := c.Request.Header.Get("Authorization")
        if tokenString == "" {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "未提供 token"})
            c.Abort()
            return
        }

        // 解析 JWT
        token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
            // 验证签名方法
            if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
                return nil, http.ErrNotSupported
            }
            return []byte(secretKey), nil // 使用您的密钥
        })

        if err != nil || !token.Valid {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "无效的 token"})
            c.Abort()
            return
        }

        // 将用户ID存入上下文
        claims := token.Claims.(jwt.MapClaims)
        c.Set("userID", claims["id"])
        c.Next()
    }
}
