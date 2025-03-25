package common

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/redis/go-redis/v9"
)

var (
	once     sync.Once
	instance *RedisClient
)

// RedisClient 封装 Redis 客户端
type RedisClient struct {
	client *redis.Client
}

// GetRedisClient 获取 Redis 客户端实例（单例模式）
func GetRedisClient(addr, password string, db int) *RedisClient {
	once.Do(func() {
		client := redis.NewClient(&redis.Options{
			Addr:     addr,
			Password: password,
			DB:       db,
		})

		// 测试连接
		ctx := context.Background()
		_, err := client.Ping(ctx).Result()
		if err != nil {
			panic("Failed to connect to Redis: " + err.Error())
		}

		instance = &RedisClient{client: client}
	})
	return instance
}

// Set 设置键值对
func (r *RedisClient) Set(key string, value interface{}, expiration time.Duration) error {
	ctx := context.Background()
	return r.client.Set(ctx, key, value, expiration).Err()
}

// Get 获取键值对
func (r *RedisClient) Get(key string) (string, error) {
	ctx := context.Background()
	val, err := r.client.Get(ctx, key).Result()

	// 检查键是否存在
	if err == redis.Nil {
		return "", fmt.Errorf("key %s does not exist", key)
	}

	// 其他错误
	if err != nil {
		return "", err
	}

	// 返回键对应的值
	return val, nil
}

// Del 删除键
func (r *RedisClient) Del(key string) error {
	ctx := context.Background()
	return r.client.Del(ctx, key).Err()
}

// Close 关闭 Redis 连接
func (r *RedisClient) Close() error {
	return r.client.Close()
}
