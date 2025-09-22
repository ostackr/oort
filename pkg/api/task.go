package api

import "context"

type Task struct {
	ID  string
	Run func(ctx context.Context) error
}

func (pipeline *Pipeline) AddTask(id string, run func(ctx context.Context) error) {
	task := &Task{
		ID:  id,
		Run: run,
	}
	pipeline.Tasks = append(pipeline.Tasks, task)
}
