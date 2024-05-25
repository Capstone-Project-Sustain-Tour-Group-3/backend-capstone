package helpers

import (
	"errors"
	"fmt"
	"time"

	"capstone/entities"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"github.com/spf13/viper"
)

type JWTClaims struct {
	Id       uuid.UUID
	Username string
	jwt.StandardClaims
}

func init() {
	viper.AutomaticEnv()
}

func GenerateAccessToken(user *entities.User) (string, error) {
	accessTokenSecret := []byte(viper.GetString("ACCESS_TOKEN_SECRET"))

	claims := JWTClaims{
		Id:       user.Id,
		Username: user.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(60 * time.Minute).Unix(),
			NotBefore: time.Now().Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedString, err := token.SignedString(accessTokenSecret)
	if err != nil {
		return "", err
	}
	return signedString, nil
}

func ParseJWT(tokenStr string) (*JWTClaims, error) {
	accessTokenSecret := []byte(viper.GetString("ACCESS_TOKEN_SECRET"))
	fmt.Println(accessTokenSecret)
	token, err := jwt.ParseWithClaims(tokenStr, &JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		return accessTokenSecret, nil
	})

	if err != nil || !token.Valid {
		if err == jwt.ErrSignatureInvalid {
			return nil, errors.New("JWT Token tidak valid")
		}
		return nil, errors.New("Token sudah kadaluwarsa")
	}

	claims := token.Claims.(*JWTClaims)
	if claims == nil {
		return nil, errors.New("Token sudah kadaluwarsa")
	}

	return claims, nil
}
