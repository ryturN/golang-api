package helper

import (
	"errors"
	"time"

	"github.com/golang-api/entity"
	"github.com/golang-jwt/jwt/v5"
)

var myKey = []byte("secretkey")

type GetJWT struct {
	UsersId string `json:"users_id"`
	jwt.RegisteredClaims
}

func NewGetJWT(user *entity.Users) (string, error) {
	claims := GetJWT{
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

func ValidateToken(accessToken string) (*string, error) {
	token, err := jwt.ParseWithClaims(accessToken, &GetJWT{}, func(token *jwt.Token) (interface{}, error) {
		return myKey, nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return nil, errors.New("invalid signature")
		}
		return nil, errors.New("your token was expired")
	}
	claims, ok := token.Claims.(*GetJWT)
	if !ok || !token.Valid {
		return nil, errors.New("unauthorized")
	}

	return &claims.UsersId, nil
}
