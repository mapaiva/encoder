package job

import (
	"sync"
)

// NewDatastore creates a new in-memory job datastore.
func NewDatastore() Datastore {
	return &memory{jobs: make(map[string]Job)}
}

// Datastore represents a persistency implementation for jobs.
type Datastore interface {
	Upsert(Job) error
	All() ([]Job, error)
}

type memory struct {
	sync.Mutex

	jobs map[string]Job
}

func (m *memory) Upsert(job Job) error {
	m.Lock()
	defer m.Unlock()
	m.jobs[job.ID] = job
	return nil
}

func (m *memory) All() ([]Job, error) {
	m.Lock()
	defer m.Unlock()
	jobs := make([]Job, 0, len(m.jobs))
	for _, job := range m.jobs {
		jobs = append(jobs, job)
	}
	return jobs, nil
}

func (m *memory) GetBy(ID string) (Job, error) {
	m.Lock()
	defer m.Unlock()
	return m.jobs[ID], nil
}
