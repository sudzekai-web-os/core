package core

type ILoggerFactory interface {
	SetMinLevel(LogLevel)
	GetMinLevel() LogLevel

	AddWriter(ILoggerWriter) ILoggerFactory
	GetWriters() []ILoggerWriter

	SetPreCategory(string) ILoggerFactory
	GetPreCategory() string

	SetSubCategory(string) ILoggerFactory
	GetSubCategory() string

	Copy() ILoggerFactory

	NewLogger(string) ILogger
}
