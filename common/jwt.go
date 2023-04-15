package common

import (
	"HolaaPlanet/entity"
	"github.com/dgrijalva/jwt-go"
	"time"
)

// jwt加密秘钥
var jwtKey = []byte("This_is_HolaaPlanet_jwt_key")

// Claims
// Maintainers:贺胜 Times:2023-04-14
// Part 1:jwt的Claims结构体
type Claims struct {
	UserID int `json:"user_id"`
	jwt.StandardClaims
}

// ReleaseToken
// Maintainers:贺胜 Times:2023-04-14
// Part 1:生成token
// Part 2:使用jwt密钥生成token
// Part 3:返回token,错误
func ReleaseToken(user entity.User) (string, error) {
	// 设置token过期时间
	expirationTime := time.Now().Add(7 * 24 * time.Hour)

	// 设置Claims
	claims := &Claims{
		// 用户ID
		UserID: user.UserID,
		// 标准字段
		StandardClaims: jwt.StandardClaims{
			// 过期时间
			ExpiresAt: expirationTime.Unix(),
			// 发行人
			Issuer: "HolaaPlanet",
			// 颁发时间
			IssuedAt: time.Now().Unix(),
			// 主题
			Subject: "user token",
		},
	}

	//使用jwt密钥生成token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)

	if err != nil {
		return "", err
	}

	// 返回token
	return tokenString, nil
}

// ParseToken
// Maintainers:贺胜 Times:2023-04-14
// Part 1:解析token
// Part 2:使用jwt密钥解析token
// Part 3:返回token,Claims,错误
func ParseToken(tokenString string) (*jwt.Token, *Claims, error) {
	claims := &Claims{}

	// 解析token
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (i interface{}, e error) {
		return jwtKey, nil
	})

	return token, claims, err
}
