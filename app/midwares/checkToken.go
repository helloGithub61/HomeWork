package middleware

import (
	"CWall/app/models"
	"CWall/app/services"
	"fmt"

	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func TokenMiddleware(c *gin.Context) {
	tokenString := c.GetHeader("Authorization")
	if tokenString == "" {
		c.Redirect(302, "/login")
		c.Abort()
		return
	}
	// 使用jwt.ParseWithClaims解析令牌字符串
	// 传入令牌字符串、声明的指针和用于获取签名的密钥的回调函数
	token, _ := jwt.ParseWithClaims(tokenString, &models.Claims{}, func(token *jwt.Token) (interface{}, error) {
		// 从令牌中提取Claims
		claims := token.Claims.(*models.Claims)
		// 根据令牌中的用户名获取对应的密钥
		user, err := services.GetUserByAccount(claims.Account)
		if err != nil {
			// 如果获取密钥失败，返回错误
			return nil, err
		}
		// 返回密钥的字节切片，用于验证令牌签名
		return []byte(user.TKey), nil
	})
	fmt.Println("====")
	fmt.Println(token.Claims.(*models.Claims).Account)
	fmt.Println("====")
	// 检查令牌是否有效
	if _, ok := token.Claims.(*models.Claims); ok {
		// 解析成功有效
		if token.Valid {
			//令牌有效
			c.Next()
		} else {
			c.Redirect(302, "/login")
			c.JSON(http.StatusUnauthorized, "解析失败")
			c.Abort()
		}
	} else {
		// 如果令牌无效或解析失败
		c.Redirect(302, "/login")
		c.JSON(http.StatusUnauthorized, "解析失败")
		c.Abort()
	}
}
