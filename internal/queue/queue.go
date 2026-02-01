package queue

import (
	"sync"

	"github.com/VaLTrexx/go-job-queue/internal/job"
)

type Queue struct {
	mu   sync.Mutex
	jobs []job.Job
}

func New() *Queue {
	return &Queue{
		jobs: make([]job.Job, 0),
	}
}

func (q *Queue) Enqueue(j job.Job) {
	q.mu.Lock()
	defer q.mu.Unlock()

	q.jobs = append(q.jobs, j)
}

func (q *Queue) Dequeue() (job.Job, bool) {
	q.mu.Lock()
	defer q.mu.Unlock()

	if len(q.jobs) == 0 {
		return job.Job{}, false
	}

	j := q.jobs[0]
	q.jobs = q.jobs[1:]

	return j, true
}
