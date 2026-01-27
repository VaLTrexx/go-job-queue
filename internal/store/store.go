package store

import (
	"sync"

	"github.com/VaLTrexx/go-job-queue/internal/job"
)

type JobStore struct {
	mu   sync.Mutex
	jobs map[string]job.Job
}

func New() *JobStore {
	return &JobStore{
		jobs: make(map[string]job.Job),
	}
}

func (s *JobStore) Save(j job.Job) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.jobs[j.ID] = j
}

func (s *JobStore) Get(id string) (job.Job, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()
	j, ok := s.jobs[id]
	return j, ok
}
