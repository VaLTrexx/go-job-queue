package worker

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/VaLTrexx/go-job-queue/internal/core"
	"github.com/VaLTrexx/go-job-queue/internal/job"
)

type Worker struct {
	ID    int
	Queue core.Queue
	Store core.Store
}

func (w *Worker) Start() {
	for {

		j, ok, err := w.Queue.Dequeue()
		if err != nil {
			fmt.Println("queue error:", err)
			time.Sleep(1 * time.Second)
			continue
		}
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
