package main

import (
	"ModEd/eval/cli/command"
	"fmt"
)

func main() {
	db, err := command.InitializeDB()
	if err != nil {
		panic(err)
	}

	evaluationController, progressController, assignmentController := command.InitializeControllers(db)

	commandExecutor := command.NewCommandExecutor()
	commandExecutor.RegisterCommand("1", &command.EvaluationCommand{
		DB:                   db,
		EvaluationController: evaluationController,
		ProgressController:   progressController,
		AssignmentController: assignmentController,
	})
	commandExecutor.RegisterCommand("resetdb", &command.ResetDBCommand{})

	for {
		command.DisplayMainMenu()
		choice := command.GetUserChoice()

		if choice == "0" {
			fmt.Println("Exiting...")
			return
		}

		if err := commandExecutor.ExecuteCommand(choice); err != nil {
			fmt.Println("Error executing command:", err)
		}
	}
}
