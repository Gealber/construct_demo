package jwtser

import (
	"crypto/rsa"
	"io/ioutil"
	"log"
	"os"
	"time"

	"github.com/Gealber/construct_demo/accounts"
	"github.com/dgrijalva/jwt-go"
)

// verify key and sign key
var (
	verifyKey *rsa.PublicKey
	signKey   *rsa.PrivateKey
)

var (
	privateKeyPath = os.Getenv("PRIVATE_KEY_PATH")
	publicKeyPath  = os.Getenv("PUBLIC_KEY_PATH")
)

//Encode de user into a jwt
func Encode(user *accounts.User) (string, error) {
	//create a signer for RSA256
	t := jwt.New(jwt.GetSigningMethod("RS256"))

	// set our claims
	user.Password = ""
	t.Claims = &accounts.CustomClaims{
		CustomUserInfo: user,
		TokenType:      "level11",
		StandardClaims: &jwt.StandardClaims{ExpiresAt: time.Now().Add(time.Minute * 20).Unix(), Issuer: "admin"},
	}
	return t.SignedString(signKey)
}

//Decode decode the jwt into a user struct
func Decode(tokenString string) (*accounts.User, error) {
	claims := &accounts.CustomClaims{}
	_, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return verifyKey, nil
	})
	if err != nil {
		log.Printf("Parsing token: %v", err)
		return nil, err
	}

	userInfo := claims.ExtractUserInfo()
	user := &accounts.User{
		Email: userInfo.Email,
	}
	return user, nil
}

// read the key files before starting http handlers
//work properly
func init() {
	var err error

	signBytes, err := ioutil.ReadFile(privateKeyPath)
	if err != nil {
		log.Fatalf("Err: %v", err)
	}

	signKey, err = jwt.ParseRSAPrivateKeyFromPEM(signBytes)
	if err != nil {
		log.Fatalf("Err: %v", err)
	}

	verifyBytes, err := ioutil.ReadFile(publicKeyPath)
	if err != nil {
		log.Fatalf("Err: %v", err)
	}

	verifyKey, err = jwt.ParseRSAPublicKeyFromPEM(verifyBytes)
	if err != nil {
		log.Fatalf("Err: %v", err)
	}
}
