package utils

import (
	redis "github.com/coroo/go-pawoon/config/redis"
	"fmt"
	"encoding/base64"
	"os"
	jwt "github.com/dgrijalva/jwt-go"
)

func CreateAuth() string {
	auth := os.Getenv("BASIC_AUTH_USERNAME") + ":" + os.Getenv("BASIC_AUTH_PASSWORD")
	return "Basic " + base64.StdEncoding.EncodeToString([]byte(auth))
}

func ExtractClaims(tokenStr string) (jwt.MapClaims, bool) {
	hmacSecret := []byte("secret")
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		 // check token signing method etc
		 return hmacSecret, nil
	})

	if err != nil {
		return nil, false
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, true
	} else {
		fmt.Println("Invalid JWT Token")
		return nil, false
	}
}

func IsInBlacklist(token string) bool {
    redisConn := redis.Connect()
    redisToken, _ := redisConn.GetValue(token)
    
    if redisToken == nil {
        return false
    }
    
    return true
}
