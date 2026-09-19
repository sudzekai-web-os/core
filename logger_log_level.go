package core

type LogLevel int

const (
	LogLevel_NONE LogLevel = iota
	LogLevel_DEBUG
	LogLevel_INFORMATION
	LogLevel_WARNING
	LogLevel_ERROR
	LogLevel_CRITICAL
)
