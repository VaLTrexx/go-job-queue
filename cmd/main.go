package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/VaLTrexx/go-job-queue/internal/job"
	"github.com/VaLTrexx/go-job-queue/internal/redis"
	"github.com/VaLTrexx/go-job-queue/internal/store"
	"github.com/VaLTrexx/go-job-queue/internal/worker"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	client := redis.NewClient()
	if err := redis.Ping(client); err != nil {
		panic(err)
	}

	q := redis.NewRedisQueue(client, "job_queue")

	s := store.New()

	for i := 0; i < 5; i++ {
		j := job.NewJob("email")
		s.Save(j)
		q.Enqueue(j)
	}

	w1 := worker.Worker{ID: 1, Queue: q, Store: s}
	w2 := worker.Worker{ID: 2, Queue: q, Store: s}

	go w1.Start()
	go w2.Start()

	time.Sleep(15 * time.Second)

	fmt.Println("\nFinal job states:")

	jobs, err := s.All()
	if err != nil {
		panic(err)
	}

	for _, j := range jobs {
		fmt.Printf("Job %s -> %s (tries: %d)\n", j.ID, j.Status, j.Tries)
	}
}
