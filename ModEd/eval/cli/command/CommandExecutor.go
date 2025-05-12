package command

import (
	"fmt"
)

type Command interface {
	Execute() error
}

type CommandExecutor struct {
	commands map[string]Command
}

func NewCommandExecutor() *CommandExecutor {
	return &CommandExecutor{commands: make(map[string]Command)}
}

func (ce *CommandExecutor) RegisterCommand(name string, command Command) {
	ce.commands[name] = command
}

func (ce *CommandExecutor) ExecuteCommand(name string) error {
	if command, exists := ce.commands[name]; exists {
		ClearTerminal()
		return command.Execute()
	}
	return fmt.Errorf("Command not found: %s", name)
}

// ClearTerminal clears the terminal screen
func ClearTerminal() {
	fmt.Print("\033[H\033[2J")
}
