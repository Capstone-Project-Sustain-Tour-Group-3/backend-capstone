package helpers

import (
	"capstone/entities"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"github.com/spf13/viper"
	"time"
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
