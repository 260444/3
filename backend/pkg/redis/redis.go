package redis

import (
	"backend/config"
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

var Rdb *redis.Client

// InitRedis 初始化Redis连接
func InitRedis() error {
	Rdb = redis.NewClient(&redis.Options{
		Addr:     config.GlobalConfig.Redis.Addr,
		Password: config.GlobalConfig.Redis.Password,
		DB:       config.GlobalConfig.Redis.DB,
	})

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// 测试连接
	_, err := Rdb.Ping(ctx).Result()
	if err != nil {
		return fmt.Errorf("连接Redis失败: %w", err)
	}

	return nil
}