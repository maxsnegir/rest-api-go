package db

import (
	"github.com/go-redis/redis"
)

type RedisStore struct {
	Client *redis.Client
	DbDsn  string
}

func (r *RedisStore) Connect() error {
	client := redis.NewClient(&redis.Options{
		Addr: r.DbDsn,
	})
	_, err := client.Ping().Result()
	if err != nil {
		return err
	}
	r.Client = client
	return nil
}

func (r *RedisStore) Close() error {
	if err := r.Client.Close(); err != nil {
		return err
	}
	return nil
}

func NewRedisStore(DbDsn string) *RedisStore {
	return &RedisStore{
		DbDsn: DbDsn,
	}
}
