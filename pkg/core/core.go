package core

import "context"

type Task struct {
	ID  string
	Run func(ctx context.Context) error
}

type Pipeline struct {
	ID     string
	Author string
	Tags   []string
	Task   []*Task
}

func NewPipeline(ID string, Author string, Tags []string) *Pipeline {
	return &Pipeline{
		ID:     ID,
		Author: Author,
		Tags:   Tags,
		Task:   []*Task{},
	}
}

func (pipeline *Pipeline) AddTask(id string, run func(ctx context.Context) error) {
	task := &Task{
		ID:  id,
		Run: run,
	}
	pipeline.Task = append(pipeline.Task, task)
}
