package job

// File gathers file informations about a job.
type File struct {
	Name string `json:"name"`
	Size int64  `json:"size"`
}
