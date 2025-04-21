package main

import (
	"flag"
	"fmt"
	"os"

	"ModEd/hr/cli/commands"
	"ModEd/hr/util"

	"gorm.io/gorm"
)

// Command interface representing an executable command.
type Command interface {
	// Execute performs the command and returns an error if something goes wrong.
	Execute(args []string, tx *gorm.DB) error
}

// Invoker holds a registry of commands.
type Invoker struct {
	commands map[string]Command
}

// NewInvoker creates a new invoker with registered commands.
func NewInvoker() *Invoker {
	return &Invoker{
		commands: map[string]Command{
			"list":              &commands.ListStudentsCommand{},
			"add":               &commands.AddStudentCommand{},
			"delete":            &commands.DeleteStudentCommand{},
			"update":            &commands.UpdateCommand{},
			"status":            &commands.UpdateStudentStatusCommand{},
			"updateStatus":      &commands.UpdateStudentStatusCommand{},
			"import":            &commands.ImportCommand{},
			"migrate":           &commands.MigrateStudentsCommand{},
			"export":            &commands.ExportStudentsCommand{},
			"request":           &commands.RequestCommand{},
			"answerResignation": &commands.AnswerResignationCommand{},
			// ... additional command registrations ...
		},
	}
}

// ExecuteCommand looks up and executes the command with the given arguments.
func (inv *Invoker) ExecuteCommand(name string, args []string, tx *gorm.DB) error {
	cmd, ok := inv.commands[name]
	if !ok {
		return fmt.Errorf("unknown command: %s", name)
	}
	return cmd.Execute(args, tx)
}

var (
	databasePath = flag.String("database", "data/ModEd.bin", "Path of SQLite Database")
)

func main() {
	flag.Parse()
	args := flag.Args()

	// Open the database
	util.DatabasePath = databasePath
	db := util.OpenDatabase(*databasePath)

	if len(args) < 1 {
		fmt.Println("Usage: go run humanresourcecli.go [-database=<path>] {list|...} [options]")
		os.Exit(1)
	}

	commandName := args[0]
	// Create an invoker and execute the command.
	invoker := NewInvoker()
	if err := invoker.ExecuteCommand(commandName, args[1:], db); err != nil {
		fmt.Printf("Error executing command: %v\n", err)
		os.Exit(1)
	}
}
