package main

import "fmt"

type Event interface {
	display()
}

type Log struct {
	ID     int
	Source string
	Body   string
}

type SystemLog struct {
	Log
	Severity string
}

func main() {
	var events []SystemLog
	slog := SystemLog{
		Log: Log{
			ID:     1,
			Source: "System",
			Body:   "System is running",
		},
		Severity: "System",
	}

	events = append(events, slog)

	for _, event := range events {
		fmt.Printf("%+v", event)
	}
}
