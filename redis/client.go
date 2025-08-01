package redis

import (
	"context"
	"github.com/go-redis/redis/v8"
	"time"
	"wallet/config"
)

var RDB *DB

type DB struct {
	client *redis.Client
}

var ctx = context.Background()

func NewRedisDB(conf config.RedisConfig) *DB {
	rdb := &DB{
		client: redis.NewClient(&redis.Options{
			Addr:     conf.Addr,
			Password: conf.Password, // no password set
			DB:       conf.DB,       // use default DB
		}),
	}
	RDB = rdb
	return rdb
}

func (d *DB) Set(key string, value interface{}, expiration time.Duration) error {
	return d.client.Set(ctx, key, value, expiration).Err()
}

func (d *DB) Get(key string) (string, error) {
	return d.client.Get(ctx, key).Result()
}

func (d *DB) Del(key string) error {
	return d.client.Del(ctx, key).Err()
}
