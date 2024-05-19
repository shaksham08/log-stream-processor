package handler

import (
	"context"
	"fmt"
	"sync"

	"github.com/shaksham08/log-stream-processor/config"
	"github.com/shaksham08/log-stream-processor/pkg/filter"
	"github.com/shaksham08/log-stream-processor/pkg/models"
	"github.com/shaksham08/log-stream-processor/pkg/processor"
)

func ProcessEvent(wg *sync.WaitGroup, ch chan models.Event, ctx context.Context) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			return
		case event, ok := <-ch:
			if !ok {
				return
			}
			processedEvent := processor.ProcessEvent(event)
			// Example filter: Only process INFO severity events.
			filter := filter.SeverityFilter("INFO")
			if filter(processedEvent) {
				fmt.Printf("Handler: %v\n", processedEvent)
				// Here we would forward the event to the egress layer.
			}
		}
	}
}

func Init(wg *sync.WaitGroup, ch chan models.Event, ctx context.Context) {
	for i := 0; i < config.MAX_HANDLER; i++ {
		wg.Add(1)
		go ProcessEvent(wg, ch, ctx)
	}
}
