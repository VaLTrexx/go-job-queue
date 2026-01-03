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
