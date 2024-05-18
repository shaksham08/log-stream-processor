package main

import (
	"sync"

	"github.com/shaksham08/log-stream-processor/pkg/handler"
	"github.com/shaksham08/log-stream-processor/pkg/models"
)

func main() {
	event := models.SystemLog{
		Log: models.Log{
			ID:     1,
			Source: "App",
			Body:   "Something here",
		},
		Severity: "Info",
	}
	var wg sync.WaitGroup

	handler.Init(event, &wg)
	wg.Wait()
}
