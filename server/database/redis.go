package database

import (
	redis "github.com/go-redis/redis"
)

var Redis *redis.Client = nil

func Init_redis(addr string, password string, db int) (*redis.Client, error) {
	Redis = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})
	_, err := Redis.Ping().Result()

	return Redis, err
}
