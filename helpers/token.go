package helpers

import (
	"errors"
	"time"

	"capstone/entities"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"github.com/spf13/viper"
)

type TokenHelper interface {
	GenerateAccessToken(user interface{}) (string, error)
	GenerateRefreshToken(user interface{}) (string, error)
}

type tokenHelper struct{}

func NewTokenHelper() TokenHelper {
	return &tokenHelper{}
}

type JWTClaims struct {
	Id       uuid.UUID
	Username string
	Role     string
	jwt.StandardClaims
}

func init() {
	viper.AutomaticEnv()
}

func (t *tokenHelper) GenerateAccessToken(user interface{}) (string, error) {
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
			ExpiresAt: time.Now().Add(20 * time.Minute).Unix(),
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

func (t *tokenHelper) GenerateRefreshToken(user interface{}) (string, error) {
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
			ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
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
