package main

import (
	"flag"
	"fmt"
	"os"

	"ModEd/hr/cli/commands"
	"ModEd/hr/util"
)

// Command interface can be defined here or in a shared package.
type Command interface {
	Run(args []string)
}

// Command registry using implementations from the commands package.
var commandsRegistry = map[string]Command{
	"list": &commands.ListStudentsCommand{},
	"add":  &commands.AddStudentCommand{},
	"updateStatus": &commands.UpdateStudentStatusCommand{},
	"import": &commands.ImportStudentsCommand{},
	// ... other command registrations ...
}

var (
	databasePath = flag.String("database", "data/ModEd.bin", "Path of SQLite Database")
)

func main() {
	flag.Parse()
	args := flag.Args()

	util.DatabasePath = databasePath

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
