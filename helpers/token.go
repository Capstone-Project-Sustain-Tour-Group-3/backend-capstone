package helpers

import (
	"errors"
	"time"

	"capstone/entities"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"github.com/spf13/viper"
)

type JWTClaims struct {
	Id       uuid.UUID
	Username string
	Role     string
	jwt.StandardClaims
}

func init() {
	viper.AutomaticEnv()
}

func GenerateAccessToken(user interface{}) (string, error) {
	accessTokenSecret := []byte(viper.GetString("ACCESS_TOKEN_SECRET"))

	var role string
	var userID uuid.UUID
	var username string

	switch u := user.(type) {
	case *entities.User:
		role = "user"
		userID = u.Id
		username = u.Username
	case *entities.Admin:
		userID = u.Id
		username = u.Username
		if u.Role == "super admin" {
			role = "superadmin"
		} else {
			role = "admin"
		}
	default:
		return "", errors.New("invalid user type")
	}

	claims := JWTClaims{
		Id:       userID,
		Username: username,
		Role:     role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(5 * time.Minute).Unix(),
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

func GenerateRefreshToken(user interface{}) (string, error) {
	accessTokenSecret := []byte(viper.GetString("REFRESH_TOKEN_SECRET"))

	var role string
	var userID uuid.UUID
	var username string

	switch u := user.(type) {
	case *entities.User:
		role = "user"
		userID = u.Id
		username = u.Username
	case *entities.Admin:
		userID = u.Id
		username = u.Username
		if u.Role == "super admin" {
			role = "superadmin"
		} else {
			role = "admin"
		}
	default:
		return "", errors.New("invalid user type")
	}

	claims := JWTClaims{
		Id:       userID,
		Username: username,
		Role:     role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(10 * time.Minute).Unix(),
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
	token, err := jwt.ParseWithClaims(tokenStr, &JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		return accessTokenSecret, nil
	})

	if err != nil || !token.Valid {
		if errors.Is(err, jwt.ErrSignatureInvalid) {
			return nil, errors.New("JWT Token tidak valid")
		}
		return nil, errors.New("Token sudah kadaluwarsa")
	}

	claims, ok := token.Claims.(*JWTClaims)

	if !ok {
		return nil, errors.New("Token tidak valid")
	}

	if claims == nil {
		return nil, errors.New("Token sudah kadaluwarsa")
	}

	return claims, nil
}
