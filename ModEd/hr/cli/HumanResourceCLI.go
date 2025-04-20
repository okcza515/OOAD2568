package main

import (
	"flag"
	"fmt"
	"os"

	"ModEd/hr/cli/commands"
	"ModEd/hr/controller"
	"ModEd/hr/util"
)

// Command interface can be defined here or in a shared package.
type Command interface {
	Run(args []string)
}

// Command registry using implementations from the commands package.
var commandsRegistry = map[string]Command{
	"list":         &commands.ListStudentsCommand{},
	"add":          &commands.AddStudentCommand{},
	"delete":       &commands.DeleteStudentCommand{},
	"update":       &commands.UpdateStudentCommand{},
	"status":       &commands.UpdateStudentStatusCommand{},
	"updateStatus": &commands.UpdateStudentStatusCommand{},
	"import":       &commands.ImportStudentsCommand{},
	"migrate":      &commands.MigrateStudentsCommand{},
	"export":       &commands.ExportStudentsCommand{},
	// ... other command registrations ...
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

	// Auto-call migration before executing any command
	if err := controller.MigrateStudentsToHR(db); err != nil {
		fmt.Printf("Migration failed: %v\n", err)
		os.Exit(1)
	}

	if len(args) < 1 {
		fmt.Println("Usage: go run humanresourcecli.go [-database=<path>] {list|...} [options]")
		os.Exit(1)
	}

	commandName := args[0]
	cmd, ok := commandsRegistry[commandName]
	if !ok {
		fmt.Printf("Unknown command: %s\n", commandName)
		os.Exit(1)
	}

	cmd.Run(args[1:])
}
