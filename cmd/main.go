package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/VaLTrexx/go-job-queue/internal/job"
	"github.com/VaLTrexx/go-job-queue/internal/queue"
	"github.com/VaLTrexx/go-job-queue/internal/store"
	"github.com/VaLTrexx/go-job-queue/internal/worker"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	q := queue.New()
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
	for _, j := range s.All() {
		fmt.Printf("Job %s -> %s (tries: %d)\n", j.ID, j.Status, j.Tries)
	}
}
