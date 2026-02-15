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

func (s *RedisStore) Get(id string) (job.Job, bool, error) {
	val, err := s.red_client.HGet(context.Background(), "jobs", id).Result()

	if err == r.Nil {
		return job.Job{}, false, nil
	}
	if err != nil {
		return job.Job{}, false, err
	}
	var j job.Job
	if err := json.Unmarshal([]byte(val), &j); err != nil {
		return job.Job{}, false, err
	}

	return j, true, nil
}

func (s *RedisStore) All() ([]job.Job, error) {
	values, err := s.red_client.HGetAll(context.Background(), "jobs").Result()
	if err != nil {
		return nil, err
	}

	jobs := make([]job.Job, 0, len(values))
	for _, val := range values {
		var j job.Job
		if err := json.Unmarshal([]byte(val), &j); err != nil {
			return nil, err
		}
		jobs = append(jobs, j)
	}

	return jobs, nil
}
