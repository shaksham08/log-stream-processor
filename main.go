package main

import (
	"sync"
	"time"

	"github.com/shaksham08/log-stream-processor/pkg/handler"
	"github.com/shaksham08/log-stream-processor/pkg/models"
)

func simulateIngress(ch chan models.Event) {
	for i := 0; i < 5; i++ {
		ch <- models.SystemLog{
			Log:      models.Log{ID: i, Source: "App", Body: "System is running"},
			Severity: "INFO",
		}
		time.Sleep(1 * time.Second)
	}
}

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

	ch := make(chan models.Event, 100)
	simulateIngress(ch)
	handler.Init(event, &wg, ch)
	close(ch)
	wg.Wait()
}
