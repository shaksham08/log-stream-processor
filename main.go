package main

import (
	"fmt"

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

	fmt.Println(event)
	// handler.ProcessEvent()
}
