package initialize

import (
	"context"
	"github.com/go-redis/redis/v8"
	"go-blog/global"
	"go-blog/global/settings"
)

func InitRedis() {
	global.RedisDB = redis.NewClient(&redis.Options{
		Addr:     settings.RedisHost,
		Password: settings.RedisPassword,
		DB:       0,
	})

	if _, err := global.RedisDB.Ping(context.Background()).Result(); err != nil {
		panic(err)
	}
}
