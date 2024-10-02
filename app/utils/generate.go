package utils

import (
	"CWall/app/models"
	"CWall/app/services"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"math/rand"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func GenerateKey(Account string, Email string, Password string) string {
	features := Account + Email + Password
	rand.Seed(time.Now().UnixNano())
	features += fmt.Sprintf("%d", rand.Intn(10000))

	hasher := sha256.New()
	hasher.Write([]byte(features))
	hashedFeatures := hasher.Sum(nil)

	// 将哈希值转换为十六进制字符串
	secretKey := hex.EncodeToString(hashedFeatures)

	return secretKey
}
func GenerateToken(Account string) (string, error) {
	user, _ := services.GetUserByAccount(Account)
	var jwtKey = []byte(user.TKey)
	expirationTime := time.Now().Add(60 * time.Minute)
	claims := &models.Claims{
		Account: Account,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	return tokenString, err
}
