package redis

import (
	"github.com/Gealber/construct_demo/accounts"
	"github.com/Gealber/construct_demo/serializer/jwt"
	"testing"
)

func Test_StoreFindRedis(t *testing.T) {
	redisRepo, err := NewRedisRepo()
	if err != nil {
		t.Fatalf("Unable to init repo")
	}

	user := &accounts.User{
		Email: "email@test.cu",
	}
	tokenString, err := jwtser.Encode(user)
	if err != nil {
		t.Fatalf("Unable to encode token: %v", err)
	}

	err = redisRepo.Store(tokenString)
	if err != nil {
		t.Fatalf("%v", err)
	}

	tokenFound, err := redisRepo.Find(user.Email)
	if err != nil {
		t.Fatalf("%v", err)
	}

	if tokenString != tokenFound {
		t.Fatalf("Tokens differ")
	}

}
