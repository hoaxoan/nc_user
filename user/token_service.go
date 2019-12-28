package user

import (
	"github.com/dgrijalva/jwt-go"
	md "github.com/hoaxoan/nc_course/nc_user/model"
)

var (
	// Define a secure key string used
	// as a salt when hashing our tokens.
	// Please make your own way more secure than this,
	// use a randomly generated md5 hash or something.
	key = []byte("mySuperSecretKeyLol")
)

type Authable interface {
	Decode(token string) (*md.CustomClaims, error)
	Encode(user *md.User) (string, error)
}

type TokenService struct {
	Repo *UserRepository
}

// Decode a token string into a token object
func (srv *TokenService) Decode(token string) (*md.CustomClaims, error) {

	// Parse the token
	tokenType, err := jwt.ParseWithClaims(string(key), &md.CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return key, nil
	})

	// Validate the token and return the custom claims
	if claims, ok := tokenType.Claims.(*md.CustomClaims); ok && tokenType.Valid {
		return claims, nil
	} else {
		return nil, err
	}
}

// Encode a claim into a JWT
func (srv *TokenService) Encode(user *md.User) (string, error) {
	// Create the Claims
	claims := md.CustomClaims{
		user,
		jwt.StandardClaims{
			ExpiresAt: 15000,
			Issuer:    "go.echo.srv.user",
		},
	}

	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign token and return
	return token.SignedString(key)
}
