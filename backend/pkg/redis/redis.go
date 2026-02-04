package redis

import (
	"backend/config"
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

// InitRedis 初始化Redis连接
func InitRedis() (*redis.Client, error) {
	db := redis.NewClient(&redis.Options{
		Addr:     config.GlobalConfig.Redis.Addr,
		Password: config.GlobalConfig.Redis.Password,
		DB:       config.GlobalConfig.Redis.DB,
	})

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// 测试连接
	_, err := db.Ping(ctx).Result()
	if err != nil {
		return nil, fmt.Errorf("连接Redis失败: %w", err)
	}

	return db, nil
}
