package model

import "github.com/dgrijalva/jwt-go"

type User struct {
	Id        int    `json:"id,omitempty" bson:"id"`
	FirstName string `json:"first_name,omitempty" bsn:"first_name"`
	LastName  string `json:"last_name,omitempty" bson:"las_name"`
	Phone     string `json:"phone,omitempty" bson:"phone"`
	Email     string `json:"email,omitempty" bson:"email"`
	Password  string `json:"password,omitempty" bson:"password"`
}

// CustomClaims is our custom metadata, which will be hashed
// and sent as the second segment in our JWT
type CustomClaims struct {
	User *User
	jwt.StandardClaims
}

type UserRequest struct {
}

type UserResponse struct {
	User   *User    `json"user,omitempty"`
	Users  []*User  `json:"users,omitempty"`
	Errors []*Error `json:"errors,omitempty"`
}

type Token struct {
	Token  string   `jon:"token,omitempty"`
	Valid  bool     `json:"valid,omitempty"`
	Errors []*Error `json:"errors,omitempty"`
}
