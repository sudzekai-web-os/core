package core

import "time"

type LogEntry struct {
	TimeStamp time.Time

	PreCategory string
	Category    string
	SubCategory string

	LogLevel LogLevel

	Message string
}
