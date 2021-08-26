package redis

import (
	"cheese/setting"
	"fmt"
	"github.com/go-redis/redis"
)

var client *redis.Client

func Init(cfg *setting.RedisConfig) (err error) {
	client = redis.NewClient(&redis.Options{
		Addr:         fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
		Password:     cfg.Password, // no password set
		DB:           cfg.DB,       // use default DB
		PoolSize:     cfg.PoolSize,
		MinIdleConns: cfg.MinIdleConns,
	})

	_, err = client.Ping().Result()
	if err != nil {
		return err
	}
	return err
}

func Close() {
	_ = client.Close()
}
