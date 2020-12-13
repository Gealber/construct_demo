package redis

import (
	"errors"
	"fmt"
	"log"
	"os"

	jwtser "github.com/Gealber/construct_demo/serializer/jwt"
	"github.com/go-redis/redis"
)

//RedisRepository represents the struct of a redis repo
type RedisRepository struct {
	URL    string
	Logger *log.Logger
	Client *redis.Client
}

//NewRedisRepo creates and returns a new redis repository
func NewRedisRepo() (*RedisRepository, error) {
	redisURL := os.Getenv("REDIS_URL")
	if redisURL == "" {
		redisURL = "redis://localhost:6379/0"
	}

	opts, err := redis.ParseURL(redisURL)
	if err != nil {
		return nil, err
	}
	client := redis.NewClient(opts)
	_, err = client.Ping().Result()
	if err != nil {
		return nil, err
	}

	logger := log.New(os.Stdout, "[REDIS]", 0)

	repo := &RedisRepository{
		URL:    redisURL,
		Logger: logger,
		Client: client,
	}
	return repo, nil
}

//Find find a key stored on the Redis DB
func (r *RedisRepository) Find(email string) (string, error) {
	data, err := r.Client.Get(email).Result()
	if err != nil {
		return "", errors.New(fmt.Sprintf("Unable to find key: %s, err: %v", email, err))
	}

	if len(data) == 0 {
		return "", errors.New(fmt.Sprintf("Unable to find key: %s", email))
	}

	return data, nil
}

//Store store a key on the Redis DB
func (r *RedisRepository) Store(token string) error {
	//I should parse the jwt and extract the email
	user, err := jwtser.Decode(token)
	if err != nil {
		return errors.New(fmt.Sprintf("Unable to parse token: %s", token))
	}

	email := user.Email
	_, err = r.Client.Set(email, token, 0).Result()
	if err != nil {
		return errors.New(fmt.Sprintf("Unable to store token: %s, err: %v", token, err))
	}
	return nil
}
