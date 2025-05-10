package main

import (
	"ModEd/core"
	"ModEd/core/migration"
	"ModEd/eval/cli/assessment"
	"ModEd/eval/cli/examination"
	"fmt"
	"os"

	"ModEd/eval/controller"

	"gorm.io/gorm"
)

const (
	defaultDBPath = "../../data/ModEd.bin"
)

type Command interface {
	Execute() error
}

// AssessmentCommand
type AssessmentCommand struct {
	db *gorm.DB
}

func (c *AssessmentCommand) Execute() error {
	assessment.RunAssessmentModuleCLI(c.db)
	return nil
}

// QuizCommand
type QuizCommand struct {
	db *gorm.DB
}

// ExaminationCommand
type ExaminationCommand struct {
	db *gorm.DB
}

func (e *ExaminationCommand) Execute() error {
	params := &examination.ExaminationCLIParams{
		ExaminationController: controller.NewExaminationController(e.db),
		SectionController:     controller.NewExamSectionController(e.db),
	}
	examination.RunExaminationCLI(params)
	return nil
}

// QuestionCommand
type QuestionCommand struct {
	db *gorm.DB
}

func (q *QuestionCommand) Execute() error {
	fmt.Println("Question module not implemented yet")
	return nil
}

// ResetDBCommand
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
	return fmt.Errorf("invalid command: %s", name)
}

func main() {
	db, err := migration.
		GetInstance().
		SetPathDB(defaultDBPath).
		MigrateModule(core.MODULE_EVAL).
		BuildDB()

	if err != nil {
		fmt.Printf("Error initializing database: %v\n", err)
		os.Exit(1)
	}

	commandExecutor := NewCommandExecutor()
	commandExecutor.RegisterCommand("1", &AssessmentCommand{db})
	//commandExecutor.RegisterCommand("2", &QuizCommand{db})
	commandExecutor.RegisterCommand("3", &ExaminationCommand{db})
	commandExecutor.RegisterCommand("4", &QuestionCommand{db})
	commandExecutor.RegisterCommand("resetdb", &ResetDBCommand{})

	for {
		displayMainMenu()
		choice := getUserChoice()

		if choice == "0" {
			fmt.Println("Exiting...")
			return
		}

		if err := commandExecutor.ExecuteCommand(choice); err != nil {
			fmt.Printf("Error: %v\n", err)
		}
	}
}

func displayMainMenu() {
	fmt.Println("\nEvaluation Module CLI Menu")
	fmt.Println("1. Assessment Management")
	//fmt.Println("2. Quiz Management")
	fmt.Println("3. Examination Management")
	fmt.Println("4. Question Management")
	fmt.Println("0. Exit")
	fmt.Println("'resetdb' to re-initialize database")
}

func getUserChoice() string {
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
		MigrateModule(core.MODULE_EVAL).
		BuildDB()

	if err != nil {
		return err
	}

	fmt.Println("Database reset successfully")
	return nil
}
