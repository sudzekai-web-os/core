package core

type ILoggerWriter interface {
	Write(LogEntry)
	WriteBatch([]LogEntry)
}
