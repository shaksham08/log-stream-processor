package handler

import (
	"context"
	"fmt"
	"sync"

	"github.com/shaksham08/log-stream-processor/config"
	"github.com/shaksham08/log-stream-processor/pkg/models"
)

func ProcessEvent(wg *sync.WaitGroup, ch chan models.Event, ctx context.Context) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			return
		case event, ok := <-ch:
			if !ok { // we are adding this because on closing this channel we are getting false repeatedly
				return
			}
			fmt.Println("THe event recieved is ", event)
		}
	}
}

func Init(wg *sync.WaitGroup, ch chan models.Event, ctx context.Context) {
	for i := 0; i < config.MAX_HANDLER; i++ {
		wg.Add(1)
		go ProcessEvent(wg, ch, ctx)
	}
}
