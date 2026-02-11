package utils

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var JwtKey = []byte("ethan972357234kdfdklsf")

type MyClaims struct {
	UserID   uint   `json:"user_id"`
	Username string `json:"username"`
	Age      int    `json:"age"`
	Email    string `json:"email"`
	jwt.RegisteredClaims
}

// 创建JWT
func GenToken(userID uint, username string, age int, email string) (string, error) {
	claims := MyClaims{
		UserID:   userID,
		Username: username,
		Age:      age,
		Email:    email,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "ethan",
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(JwtKey)
}

// 解析JWT
func ParseToken(tokenStr string) (*MyClaims, error) {
	token, err := jwt.ParseWithClaims(
		tokenStr,
		&MyClaims{},
		func(token *jwt.Token) (interface{}, error) {
			return JwtKey, nil
		},
	)
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*MyClaims)

	// 解析后果token
	fmt.Println("解析后的token：", claims)

	if !ok || !token.Valid {
		return nil, jwt.ErrInvalidKey
	}
	return claims, nil
}
