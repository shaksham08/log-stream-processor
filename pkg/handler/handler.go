package handler

import (
	"fmt"
	"sync"

	"github.com/shaksham08/log-stream-processor/config"
	"github.com/shaksham08/log-stream-processor/pkg/models"
)

func processEvent(ch chan models.Event, wg *sync.WaitGroup, idx int) {
	for event := range ch {
		fmt.Println("Event recieved : ", event, "by handler", idx)
	}
	defer wg.Done()
}

func Init(ch chan models.Event, wg *sync.WaitGroup) {
	for i := 0; i < config.MAX_HANDLERS; i++ {
		wg.Add(1)
		go processEvent(ch, wg, i)
	}
}
