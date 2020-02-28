package transcoder

import (
	"os"
	"path"
	"strings"

	"github.com/xfrr/goffmpeg/models"
)

// Options defines options for transcoding a video.
type Options struct {
	InputPath  string
	OutputPath string
	OnProgress func(models.Progress) error
}

// OptionsFrom returns a set of standard options given a file info.
func OptionsFrom(file os.FileInfo, directory string) Options {
	return Options{
		InputPath:  directory + "/" + file.Name(),
		OutputPath: directory + "/" + strings.TrimSuffix(file.Name(), path.Ext(file.Name())) + ".mp4",
	}
}
