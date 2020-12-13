package accounts

import (
	"github.com/dgrijalva/jwt-go"
)

// User represents the structure of an user on the dashboard
type User struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}

// CustomClaims contains the neccessary info
type CustomClaims struct {
	CustomUserInfo *User
	TokenType      string
	*jwt.StandardClaims
}

//ExtractUserInfo ...
func (c *CustomClaims) ExtractUserInfo() *User {
	return c.CustomUserInfo
}

func (c *CustomClaims) Valid() error {
	return c.StandardClaims.Valid()
}

// Token contains the jwt-token
type Token struct {
	Token string `json:"token" bson:"token"`
}
