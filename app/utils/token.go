package utils

import (
	"CWall/app/models"
	"CWall/app/services"

	"github.com/dgrijalva/jwt-go"
)

func GetAccountByToken(tokenString string) string {
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
	Account := token.Claims.(*models.Claims).Account
	return Account
}
