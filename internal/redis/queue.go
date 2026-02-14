package redis

import (
	"context"
	"encoding/json"

	r "github.com/redis/go-redis/v9"

	"github.com/VaLTrexx/go-job-queue/internal/core"
	"github.com/VaLTrexx/go-job-queue/internal/job"
)

type RedisQueue struct {
	client *r.Client
	key    string
}

func NewRedisQueue(client *r.Client, key string) core.Queue {
	return &RedisQueue{
		client: client,
		key:    key,
	}
}

func (q *RedisQueue) Enqueue(j job.Job) error {
	data, err := json.Marshal(j)
	if err != nil {
		return err
	}

	return q.client.LPush(context.Background(), q.key, data).Err()
}

func (q *RedisQueue) Dequeue() (job.Job, bool, error) {
	result, err := q.client.BRPop(context.Background(), 0, q.key).Result()
	if err != nil {
		return job.Job{}, false, err
	}

	var j job.Job
	if err := json.Unmarshal([]byte(result[1]), &j); err != nil {
		return job.Job{}, false, err
	}

	return j, true, nil
}
