package redis

import (
	"log"
)

//RedisRepository represents the struct of a redis repo
type RedisRepository struct {
	URL    string
	Logger *log.Logger
}
