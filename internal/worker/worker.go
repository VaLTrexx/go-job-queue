package worker

import (
	"fmt"
	"time"

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
		job, ok := w.Queue.Dequeue()
		if !ok {
			time.Sleep(1 * time.Second)
			continue
		}

		fmt.Printf("Worker %d processing job %s\n", w.ID, job.ID)
	}
}
