package helper

import (
	"log"
	"time"

	"github.com/golang-jwt/jwt"
)

func GenerateJWT(id uint) string {
	var informasi = jwt.MapClaims{}
	informasi["id"] = id
	informasi["valid"] = true
	informasi["exp"] = time.Now().Add(time.Hour + 1)
	jwt := jwt.NewWithClaims(jwt.SigningMethodHS256, informasi)
	token, err := jwt.SignedString("S3cr3t!!")
	if err != nil {
		log.Println("generate jwt error ", err.Error())
		return ""
	}

	return token
}
