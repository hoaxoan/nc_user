package service

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/hoaxoan/nc_user/user"
	"github.com/hoaxoan/nc_user/model"
	"github.com/hoaxoan/nc_user/config"
	"time"
)

var (
	// Define a secure key string used
	// as a salt when hashing our tokens.
	// Please make your own way more secure than this,
	// use a randomly generated md5 hash or something.
	key = []byte("mySuperSecretKeyLol")
)

type Authable interface {
	Decode(token string) (*model.CustomClaims, error)
	Encode(user *model.User) (string, error)
}

type tokenService struct {
	Repo        user.Repository
}

func NewTokeService(repo user.Repository) Authable {
	return &tokenService{Repo:    repo}
}

// Decode a token string into a token object
func (srv *tokenService) Decode(token string) (*model.CustomClaims, error) {

	// Parse the token
	signKey := []byte(config.Config.JWTSecret.JWTKey)
	tokenType, err := jwt.ParseWithClaims(token, &model.CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return signKey, nil
	})

	// Validate the token and return the custom claims
	if claims, ok := tokenType.Claims.(*model.CustomClaims); ok && tokenType.Valid {
		return claims, nil
	} else {
		return nil, err
	}
}

// Encode a claim into a JWT
func (srv *tokenService) Encode(user *model.User) (string, error) {
	// Create the Claims
	claims := model.CustomClaims{
		user,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(48 * time.Hour).Unix(), //unix timestamp
			Issuer:    "go.echo.srv.user",
		},
	}

	// Create token
	signKey := []byte(config.Config.JWTSecret.JWTKey)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign token and return
	return token.SignedString(signKey)
}
