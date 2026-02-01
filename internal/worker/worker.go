package worker

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/VaLTrexx/go-job-queue/internal/job"
	"github.com/VaLTrexx/go-job-queue/internal/queue"
	"github.com/VaLTrexx/go-job-queue/internal/store"
)

type Worker struct {
	ID    int
	Queue *queue.Queue
	Store *store.JobStore
}

func (w *Worker) Start() {
	for {
		j, ok := w.Queue.Dequeue()
		if !ok {
			time.Sleep(1 * time.Second)
			continue
		}
		j.Status = job.StatusRunning
		w.Store.Save(j)

		fmt.Printf("Worker %d picked job %s (try %d)\n", w.ID, j.ID, j.Tries+1)
		time.Sleep(2 * time.Second)

		if rand.Intn(2) == 0 {
			j.Tries++

			if j.Tries >= 3 {
				j.Status = job.StatusFailed
				w.Store.Save(j)
				fmt.Printf("Worker %d FAILED job %s permanently\n", w.ID, j.ID)
			} else {
				j.Status = job.StatusPending
				w.Store.Save(j)
				w.Queue.Enqueue(j)
				fmt.Printf("Worker %d failed job %s, retrying\n", w.ID, j.ID)
			}

			continue
		}
		j.Status = job.StatusDone
		w.Store.Save(j)
		fmt.Printf("Worker %d completed job %s\n", w.ID, j.ID)
	}
}
