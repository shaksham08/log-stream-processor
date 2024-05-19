package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
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

func listenForCancel(cancel context.CancelFunc, ch chan models.Event, wg *sync.WaitGroup) {
	defer wg.Done()
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	<-sigs
	fmt.Println("Received signal to cancel")
	close(ch)
	cancel()
}

func main() {
	var wg sync.WaitGroup
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	ch := make(chan models.Event, 100)
	handler.Init(&wg, ch, ctx)
	simulateIngress(ch)
	wg.Add(1)
	go listenForCancel(cancel, ch, &wg)
	wg.Wait()
}
