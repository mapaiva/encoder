package transcoder

import (
	"github.com/xfrr/goffmpeg/models"
	"github.com/xfrr/goffmpeg/transcoder"
)

// NewFfmpeg creates a new ffmpeg transcoder.
func NewFfmpeg(options Options) Transcoder {
	return &ffmpeg{options: options}
}

type ffmpeg struct {
	options  Options
	done     <-chan error
	progress <-chan models.Progress
}

func (f *ffmpeg) Start() error {
	trans := new(transcoder.Transcoder)
	if err := trans.Initialize(f.options.InputPath, f.options.OutputPath); err != nil {
		return err
	}
	f.done = trans.Run(true)
	f.progress = trans.Output()
	return nil
}

func (f *ffmpeg) Done() <-chan error {
	return f.done
}

func (f *ffmpeg) Progress() <-chan models.Progress {
	return f.progress
}
