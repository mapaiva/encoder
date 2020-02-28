package job

import "encoding/json"

// Job statuses.
const (
	StatusRunning Status = iota
	StatusFailed
	StatusDone
)

var (
	statusNames = map[Status]string{
		StatusRunning: "running",
		StatusFailed:  "failed",
		StatusDone:    "done",
	}
)

// Status represents a status of a job.
type Status int

// MarshalJSON marshals status as a string.
func (s Status) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.String())
}

func (s Status) String() string {
	return statusNames[s]
}
