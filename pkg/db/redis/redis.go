package redis

import (
	"context"
	"os"
	"strconv"
	"time"

	"github.com/Cenzios/pf-backend/pkg/db/dbiface"
	"github.com/Cenzios/pf-backend/pkg/logger"
	"github.com/redis/go-redis/v9"
)

type redisImpl struct {
	client *redis.Client
}

var ctx = context.Background()

// New returns a new redisImpl as dbiface.Cache
func New() dbiface.Cache {
	addr := os.Getenv("REDIS_ADDR")     // e.g., "localhost:6379"
	pass := os.Getenv("REDIS_PASSWORD") // optional
	dbNum, _ := strconv.Atoi(os.Getenv("REDIS_DB"))

	rdb := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: pass,
		DB:       dbNum,
	})

	if err := rdb.Ping(ctx).Err(); err != nil {
		logger.Error.Fatalln("❌ Redis ping failed:", err)
	}

	logger.Info.Println("✅ Redis connected")
	return &redisImpl{client: rdb}
}

func (r *redisImpl) Get(key string) (string, error) {
	return r.client.Get(ctx, key).Result()
}

func (r *redisImpl) Set(key string, value string, ttlSeconds int) error {
	return r.client.Set(ctx, key, value, time.Duration(ttlSeconds)*time.Second).Err()
}

func (r *redisImpl) Delete(key string) error {
	return r.client.Del(ctx, key).Err()
}

func (r *redisImpl) Exists(key string) (bool, error) {
	result, err := r.client.Exists(ctx, key).Result()
	return result > 0, err
}
