package job

type Status string

const (
	StatusPending Status = "PENDING"
	StatusRunning Status = "RUNNING"
	StatusDone    Status = "DONE"
	StatusFailed  Status = "FAILED"
)

type Job struct {
	ID     string
	Type   string
	Status Status
	Tries  int
}

func NewJob(jobType string) Job {
	return Job{
		ID:     NewID(),
		Type:   jobType,
		Status: StatusPending,
		Tries:  0,
	}
}
