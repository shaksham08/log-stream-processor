package handler

import (
	"fmt"

	"github.com/shaksham08/log-stream-processor/pkg/models"
)

func Init(ch chan models.Event) {
	for event := range ch {
		fmt.Println("Event recieved : ", event)
	}
}
