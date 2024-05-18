package models

import "fmt"

type Event interface {
	Display()
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

func (s SystemLog) Display() {
	fmt.Println("ID: ", s.Log.ID)
	fmt.Println("Source: ", s.Log.Source)
	fmt.Println("Body: ", s.Log.Body)
	fmt.Println("Severity: ", s.Severity)
}
