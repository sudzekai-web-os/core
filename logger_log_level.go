package core

type LogLevel int

const (
	LogLevel_DEBUG LogLevel = iota
	LogLevel_INFORMATION
	LogLevel_WARNING
	LogLevel_ERROR
	LogLevel_CRITICAL
	LogLevel_NONE
)
