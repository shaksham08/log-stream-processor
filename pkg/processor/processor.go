package processor

import "github.com/shaksham08/log-stream-processor/pkg/models"

// ProcessEvent modifies or enriches an event.
func ProcessEvent(event models.Event) models.Event {
	// Example: Adding a timestamp or modifying the event's body.
	// This is a placeholder for your processing logic.
	switch e := event.(type) { // Type assertion
	case models.SystemLog:
		e.Body += " - Processed"
		return e
	default:
		// If the event type is unknown, return it unmodified.
		return event
	}
}

func ProcessStringToEvent(str string) models.Event {
	return models.SystemLog{
		Log:      models.Log{ID: 1, Source: "App", Body: str},
		Severity: "INFO",
	}
}
