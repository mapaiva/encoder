package transcoder

import (
	"os"
	"path"
	"strings"
)

// Options defines options for transcoding a video.
type Options struct {
	InputPath  string `json:"input_path"`
	OutputPath string `json:"output_path"`
}

// OptionsFrom returns a set of standard options given a file info.
func OptionsFrom(file os.FileInfo, directory string) Options {
	return Options{
		InputPath:  directory + "/" + file.Name(),
		OutputPath: directory + "/" + strings.TrimSuffix(file.Name(), path.Ext(file.Name())) + ".mp4",
	}
}
