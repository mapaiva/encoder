package server

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/mapaiva/encoder/job"
)

// New returns a new server.
func New(jobs job.Datastore) *Server {
	return &Server{jobs: jobs}
}

// Server represents an HTTP server.
type Server struct {
	jobs job.Datastore
}

// Start starts a new server.
func (s *Server) Start() error {
	e := echo.New()
	e.GET("/jobs", s.Jobs)
	return e.Start(":1323")
}

// Jobs represents a handler that retrives a job list.
func (s *Server) Jobs(ctx echo.Context) error {
	jobs, err := s.jobs.All()
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, jobs)
}
