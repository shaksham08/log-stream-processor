package main

import (
	"sync"
	"time"

	"github.com/shaksham08/log-stream-processor/pkg/handler"
	"github.com/shaksham08/log-stream-processor/pkg/models"
)

func simulateIngress(ch chan models.Event, wg *sync.WaitGroup) {
	for i := 0; i < 10; i++ {
		ch <- models.SystemLog{
			Log: models.Log{
				ID:     i,
				Source: "System",
				Body:   "System is running",
			},
			Severity: "Info",
		}
		time.Sleep(1 * time.Second)
	}
	defer wg.Done()
}

func main() {
	ch := make(chan models.Event)
	var wg sync.WaitGroup
	go handler.Init(ch, &wg)
	wg.Add(1)
	simulateIngress(ch, &wg)
	close(ch)
	defer wg.Wait()
}
