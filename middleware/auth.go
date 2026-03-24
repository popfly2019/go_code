package middleware

import (
	"go_demo/pkg/response"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var jwtKey = []byte("my-secret-key-123")

type UserClaims struct {
	UserID   uint   `json:"user_id"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func GenerateToken(userID uint, username string) (string, error) {
	expireTime := time.Now().Add(2 * time.Hour)
	claims := &UserClaims{
		UserID:   userID,
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expireTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "gin-demo",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}

func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取Token
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			response.Error(c, 401, "请先登录")
			c.Abort()
			return
		}

		// 解析Token（Bearer xxxx）
		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			response.Error(c, 401, "Token格式错误")
			c.Abort()
			return
		}

		// 验证Token
		tokenStr := parts[1]
		claims := &UserClaims{}
		token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})

		if err != nil || !token.Valid {
			response.Error(c, 401, "Token失效，请重新登录")
			c.Abort()
			return
		}

		// 将用户信息存入上下文，供后续接口使用
		c.Set("userID", claims.UserID)
		c.Set("username", claims.Username)
		c.Next()
	}
}
