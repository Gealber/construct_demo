package jwtser

import (
	"testing"

	"github.com/Gealber/construct_demo/accounts"
)

func Test_JWTEncodeDecode(t *testing.T) {
	user := &accounts.User{
		Email:    "email@test.com",
		Username: "Lola",
		Password: "TrickyOnePass#",
	}

	tokenString, err := Encode(user)
	if err != nil {
		t.Fatalf("Unable to encode user: %v", err)
	}

	userDecode, err := Decode(tokenString)
	if err != nil {
		t.Fatalf("Unable to decode tokenString: %v", err)
	}

	if user.Email != userDecode.Email {
		t.Fatalf("Encode and decode emails user doesn't match: %s %s", user.Email, userDecode.Email)
	}
}
