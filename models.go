package main

import "fmt"

type Event interface {
	display()
}

// Base struct
type Log struct {
	ID     int
	Source string
	Body   string
}

type SystemLog struct {
	Log
	Severity string
}

func (s SystemLog) display() {
	fmt.Printf("%+v", s)
}

func printSomething(e Event) {
	fmt.Printf("%+v", e)
}

func main() {
	var events []Event
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
		event.display()
		// Now since system log implements the function names display and interface Event also has function display
		// So SystemLog is also now part of Event struct and now System log can also access all the functions associated with Event
		printSomething(event)
	}
}
