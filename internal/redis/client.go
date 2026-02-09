package redis

import (
	"context"

	r "github.com/redis/go-redis/v9"
)

func NewClient() *r.Client {
	return r.NewClient(&r.Options{
		Addr: "localhost:6379",
		DB:   0,
	})
}

func Ping(client *r.Client) error {
	return client.Ping(context.Background()).Err()
}
