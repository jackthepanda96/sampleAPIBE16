package helper

import (
	"log"

	"github.com/golang-jwt/jwt"
)

func GenerateJWT(id uint, role string) string {
	var informasi = jwt.MapClaims{}
	informasi["id"] = id
	informasi["role"] = "admin"

	rawToken := jwt.NewWithClaims(jwt.SigningMethodHS256, informasi)

	resultToken, err := rawToken.SignedString([]byte("S3cr3t!!"))
	if err != nil {
		log.Println("generate jwt error ", err.Error())
		return ""
	}

	return resultToken
}

func DecodeJWT(token *jwt.Token) (uint, string) {
	if token.Valid {
		data := token.Claims.(jwt.MapClaims)
		user_id := data["id"].(float64)
		role := data["role"].(string)

		return uint(user_id), role
	}

	return 0, ""
}
