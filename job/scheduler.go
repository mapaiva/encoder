package job

import "github.com/xfrr/goffmpeg/models"

// NewScheduler creates a new schedule instance.
func NewScheduler(jobs Datastore) *Scheduler {
	return &Scheduler{jobs: jobs}
}

// Scheduler coordinates all transcoding jobs executions.
type Scheduler struct {
	jobs Datastore
}

// Schedule handles a transcoding job.
func (s *Scheduler) Schedule(job Job) error {
	if err := s.jobs.Upsert(job); err != nil {
		return err
	}
	if err := job.Transcoder.Start(); err != nil {
		return err
	}
	for progress := range job.Transcoder.Progress() {
		s.onProgress(progress, job)
	}
	if err := <-job.Transcoder.Done(); err != nil {
		return err
	}
	return s.finish(job)
}

func (s *Scheduler) onProgress(progress models.Progress, job Job) {
	job.Progress = progress
	s.jobs.Upsert(job)
}

func (s *Scheduler) finish(job Job) error {
	job.Status = StatusDone
	return s.jobs.Upsert(job)
}
