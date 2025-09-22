package api

import (
	"sync"

	"github.com/robfig/cron/v3"
)

type Scheduler struct {
	cron      *cron.Cron
	pipelines map[string]*registeredPipeline
	mu        sync.Mutex
}

type registeredPipeline struct {
	p       *Pipeline
	running int32 // idle = 0, busy = 1
}

func NewScheduler() *Scheduler {
	return &Scheduler{
		cron:      cron.New(cron.WithSeconds()),
		pipelines: make(map[string]*registeredPipeline),
	}
}

func (s *Scheduler) Start() {
	s.cron.Start()
}

func (s *Scheduler) Stop() {
	s.cron.Stop()
}
