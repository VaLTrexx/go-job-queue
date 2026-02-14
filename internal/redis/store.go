package redis

import (
	"context"
	"encoding/json"

	r "github.com/redis/go-redis/v9"

	"github.com/VaLTrexx/go-job-queue/internal/core"
	"github.com/VaLTrexx/go-job-queue/internal/job"
)

type RedisStore struct {
	red_client *r.Client
}

func NewRedisStore(client *r.Client) core.Store {
	return &RedisStore{red_client: client}
}

func (s *RedisStore) Save(j job.Job) error {
	data, err := json.Marshal(j)
	if err != nil {
		return err
	}

	return s.red_client.HSet(context.Background(), "jobs", j.ID, data).Err()
}
