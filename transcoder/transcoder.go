package transcoder

import (
	"github.com/xfrr/goffmpeg/models"
)

// Transcoder transcodes a video file into another.
type Transcoder interface {
	Start() error
	Done() <-chan error
	Progress() <-chan models.Progress
}
