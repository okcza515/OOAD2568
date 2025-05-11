package main

import (
	"ModEd/core"

	"ModEd/core/migration"

	evaluation "ModEd/eval/cli/evaluation"

	"fmt"

	controller "ModEd/eval/controller"

	"gorm.io/gorm"
)

const (
	defaultDBPath = "../../data/ModEd.bin"
)

type Command interface {
	Execute() error
}

type EvaluationCommand struct {
	db                   *gorm.DB
	evaluationController *controller.EvaluationController
	progressController   *controller.ProgressController
	assessmentController *controller.AssessmentController
}

func (e *EvaluationCommand) Execute() error {
	evaluation.RunEvalModuleCLI(e.db, e.evaluationController, e.progressController, e.assessmentController)
	return nil
}

type ResetDBCommand struct{}

func (r *ResetDBCommand) Execute() error {
	return resetDB()
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
		return command.Execute()
	}
	return fmt.Errorf("Command not found: %s", name)
}

func main() {
	db, err := migration.
		GetInstance().
		SetPathDB(defaultDBPath).
		MigrateModule(core.MODULE_QUIZ).
		BuildDB()

	if err != nil {
		panic(err)
	}

	evaluationController := controller.NewEvaluationController(db)
	progressController := controller.NewProgressController(db)
	assessmentController := controller.NewAssessmentController(db)

	CommandExecutor := NewCommandExecutor()
	CommandExecutor.RegisterCommand("1", &EvaluationCommand{db, evaluationController, progressController, assessmentController})

	for {
		DisplayMainMenu()
		choice := GetUserChoice()

		if choice == "0" {
			fmt.Println("Exiting...")
			return
		}

		if err := CommandExecutor.ExecuteCommand(choice); err != nil {
			fmt.Println("Error executing command:", err)
		}
	}
}

func DisplayMainMenu() {
	fmt.Println("\nEvaluation Module Menu:")
	fmt.Println("1. Evaluation Assignment & Quiz")
	fmt.Println("2. Evaluation Examination")
	fmt.Println("'resetdb' to re-initialize the database")
}

func GetUserChoice() string {
	var choice string
	fmt.Print("Enter your choice: ")
	fmt.Scanln(&choice)
	return choice
}

func resetDB() error {
	err := migration.GetInstance().DropAllTables()
	if err != nil {
		return err
	}

	_, err = migration.GetInstance().
		SetPathDB(defaultDBPath).
		MigrateModule(core.MODULE_QUIZ).
		MigrateModule(core.MODULE_COMMON).
		MigrateModule(core.MODULE_INSTRUCTOR).
		BuildDB()

	if err != nil {
		return err
	}

	return nil
}
