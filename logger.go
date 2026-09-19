package core

type ILogger interface {
	Log(level LogLevel, format string, args ...any)

	LogDebug(format string, args ...any)
	LogInformation(format string, args ...any)
	LogWarning(format string, args ...any)
	LogError(format string, args ...any)
	LogCritical(format string, args ...any)
}
