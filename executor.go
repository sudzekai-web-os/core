package core

type IExecutor interface {
	Execute(command string, args ...string) CommandResult
	ExecuteInDirectory(directoryPath string, command string, args ...string) CommandResult
	ExecuteWithInput(input, command string, args ...string) CommandResult
}
