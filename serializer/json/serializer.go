package json

import (
	"encoding/json"

	"github.com/Gealber/construct_demo/accounts"
)

//Decode a slice of bytes into an User
func Decode(input []byte) (*accounts.User, error) {
	user := &accounts.User{}
	if err := json.Unmarshal(input, user); err != nil {
		return nil, err
	}
	return user, nil
}

//Encode an user into a slice of bytes
func Encode(user *accounts.User) ([]byte, error) {
	rawUser, err := json.Marshal(user)
	if err != nil {
		return nil, err
	}
	return rawUser, nil
}
