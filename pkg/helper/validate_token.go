package helper

import (
	"fmt"
	"strings"

	"github.com/golang-jwt/jwt"
)

const SECRET_KEY = "1234567"

type Claims struct {
	jwt.StandardClaims
	UserId int
}

func GetUserId(tokenStr string) int {
	claims := &Claims{}
	fmt.Println(claims)

	data, err := jwt.ParseWithClaims(tokenStr, claims, func(t *jwt.Token) (interface{}, error) {
		return []byte(SECRET_KEY), nil
	})

	fmt.Println(data)

	if err != nil {
		return 0
	}

	return claims.UserId
}

func ExtractToken(bearToken string) string {
	// normally Authorization the_token_xxx
	strArr := strings.Split(bearToken, " ")
	if len(strArr) == 2 {
		return strArr[1]
	}
	return ""
}
