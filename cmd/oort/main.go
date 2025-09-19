package main

import (
	"context"
	"log"
	"time"

	"github.com/ostackr/oort/pkg/core"
)

func main() {
	core.NewPipeline("test_pipeline", "Finn", []string{"playground"}).AddTask("transform", func(ctx context.Context) error {
		log.Println("transforming...")
		time.Sleep(500 * time.Millisecond)
		return nil
	})
}
