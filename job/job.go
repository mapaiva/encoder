package job

import (
	"github.com/mapaiva/encoder/transcoder"
	"github.com/xfrr/goffmpeg/models"
)

// Job represents a running transcoding task.
type Job struct {
	ID         string                `json:"id"`
	File       File                  `json:"file"`
	Progress   models.Progress       `json:"progress"`
	Status     Status                `json:"status"`
	Options    transcoder.Options    `json:"options"`
	Transcoder transcoder.Transcoder `json:"-"`
}
