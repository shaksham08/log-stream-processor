package handler

import (
	"fmt"
	"sync"

	"github.com/shaksham08/log-stream-processor/config"
	"github.com/shaksham08/log-stream-processor/pkg/models"
)

func ProcessEvent(wg *sync.WaitGroup, ch chan models.Event) {
	defer wg.Done()
	for chanEvent := range ch {
		fmt.Println(chanEvent)
	}

}

func Init(event models.SystemLog, wg *sync.WaitGroup, ch chan models.Event) {
	for i := 0; i < config.MAX_HANDLER; i++ {
		wg.Add(1)
		go ProcessEvent(wg, ch)
	}
}
