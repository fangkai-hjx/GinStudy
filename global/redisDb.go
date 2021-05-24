package global

import (
	"context"
	"fmt"
	"github.com/go-redis/redis"
)

var (
	RedisDb *redis.Client
	Ctx     = context.Background()
)

//初始化Redis连接池
func SetupRedisDb() error {
	RedisDb = redis.NewClient(&redis.Options{
		Addr:     RedisSetting.Addr,
		Password: RedisSetting.Password,
		DB:       0,
	})
	_, err := RedisDb.Ping(Ctx).Result()
	if err != nil {
		fmt.Println("RedisDb.ping err：", err)
		return err
	}
	fmt.Println("Redis pool init success")
	return nil
}
