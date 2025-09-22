package api

import (
	"context"
	"fmt"
	"log"
	"sync/atomic"
)

type Pipeline struct {
	ID     string
	Author string
	Tags   []string
	Cron   string
	Tasks  []*Task
}

func NewPipeline(ID string, Author string, Tags []string, Cron string) *Pipeline {
	return &Pipeline{
		ID:     ID,
		Author: Author,
		Tags:   Tags,
		Cron:   Cron,
		Tasks:  []*Task{},
	}
}

func (s *Scheduler) RegisterPipeline(p *Pipeline) error {
	s.mu.Lock()
	if _, exists := s.pipelines[p.ID]; exists {
		s.mu.Unlock()
		return fmt.Errorf("pipeline %s already registered", p.ID)
	}
	rp := &registeredPipeline{p: p}
	s.pipelines[p.ID] = rp
	s.mu.Unlock()

	_, err := s.cron.AddFunc(p.Cron, func() {
		if !atomic.CompareAndSwapInt32(&rp.running, 0, 1) {
			log.Printf("pipeline %s is still running, skipping this schedule", p.ID)
			return
		}

		go func() {
			defer atomic.StoreInt32(&rp.running, 0)
			log.Printf("starting pipeline %s (author=%s)", rp.p.ID, rp.p.Author)
			for _, t := range rp.p.Tasks {
				ctx := context.Background()
				if err := t.Run(ctx); err != nil {
					log.Printf("task %s in pipeline %s failed: %v", t.ID, rp.p.ID, err)
				} else {
					log.Printf("task %s in pipeline %s succeeded", t.ID, rp.p.ID)
				}
			}
			log.Printf("pipeline %s finished", rp.p.ID)
		}()
	})
	if err != nil {
		return fmt.Errorf("failed to add cron func: %w", err)
	}

	return nil
}
