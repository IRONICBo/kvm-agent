package conn

import (
	"context"
	"fmt"
	"kvm-agent/internal/config"
	"kvm-agent/internal/log"

	"github.com/go-redis/redis/v8"
)

var r *redis.Client

// InitRedisDB init redis client.
func InitRedisDB() {
	r = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", config.Config.Redis.Ip, config.Config.Redis.Port),
		Password: config.Config.Redis.Password, // no password set
		DB:       config.Config.Redis.Database, // use default DB
	})
	pong, err := r.Ping(context.Background()).Result()
	if err != nil {
		log.Panic("Redis", "connect ping failed, err:", err.Error())
	} else {
		log.Info("Redis", "connect ping response:", pong)
	}
}

// GetRedis get redis client instance.
func GetRedis() *redis.Client {
	return r
}

// CloseRedis close redis client instance.
func CloseRedis() {
	if r != nil {
		err := r.Close()
		if err != nil {
			log.Error("Redis", "close failed, err:", err.Error())
		}
	}
}
