package redis

import (
	"fmt"
	"short-url/config"

	"github.com/go-redis/redis"
)

var RedisClient *redis.Client

func InitRedis() {
	RedisClient := redis.NewClient(&redis.Options{
		Addr:     config.Cfg.RedisCfg.Address,
		Password: config.Cfg.RedisCfg.Password,
		DB:       config.Cfg.RedisCfg.Db,
	})
	result, err := RedisClient.Ping().Result()
	if err != nil {
		panic(fmt.Sprintf("创建redis客户端失败: %s,%s", result, err))
		return
	}
}
