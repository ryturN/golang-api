package helper

import (
	"time"

	"github.com/golang-api/entity"
	"github.com/golang-jwt/jwt/v5"
)

var myKey = []byte("secretkey")

type getJWT struct {
	UsersId string `json:"users_id"`
	jwt.RegisteredClaims
}

func NewGetJWT(user *entity.Users) (string, error) {
	claims := getJWT{
		user.UsersId,
		jwt.RegisteredClaims{
			Issuer:    "golang-api",
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(60 * time.Minute)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(myKey)
	return ss, err
}
