package core

import "github.com/VaLTrexx/go-job-queue/internal/job"

// Queue defines how jobs are enqueued and dequeued
type Queue interface {
	Enqueue(job.Job) error
	Dequeue() (job.Job, bool, error)
}

// Store defines how job state is persisted
type Store interface {
	Save(job.Job) error
	Get(string) (job.Job, bool, error)
	All() ([]job.Job, error)
}
