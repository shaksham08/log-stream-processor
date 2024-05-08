package handler

import (
	"context"
	"fmt"
	"sync"

	"github.com/shaksham08/log-stream-processor/config"
	"github.com/shaksham08/log-stream-processor/pkg/models"
)

func processEvent(ch chan models.Event, wg *sync.WaitGroup, i int, ctx context.Context) {

	defer wg.Done()

	for {
		select {
		case <-ctx.Done():
			fmt.Println("Closing handler ", i)
			return
		case event, ok := <-ch:
			{
				if !ok {
					return
				}
				fmt.Println("Event recieved ", event, "by handler ", i)
			}
		}

	}

}

func Init(ch chan models.Event, wg *sync.WaitGroup, ctx context.Context) {
	for i := 0; i < config.MAX_HANDLERS; i++ {
		wg.Add(1)
		go processEvent(ch, wg, i+1, ctx)
	}
}
