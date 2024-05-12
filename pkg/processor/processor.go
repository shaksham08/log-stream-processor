package processor

import "github.com/shaksham08/log-stream-processor/pkg/models"

func ProcessStringToEvent(str string) models.Event {
	return models.SystemLog{
		Log:      models.Log{ID: 1, Source: "App", Body: str},
		Severity: "INFO",
	}
}
