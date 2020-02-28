package job

import (
	"os"

	"github.com/mapaiva/encoder/transcoder"
	"github.com/xfrr/goffmpeg/models"
)

// Job statuses.
const (
	StatusRunning Status = iota
	StatusFailed
	StatusDone
)

// Status represents a status of a job.
type Status int

// Job represents a running transcoding task.
type Job struct {
	ID         string
	File       os.FileInfo
	Transcoder transcoder.Transcoder
	Progress   models.Progress
	Status     Status
}
