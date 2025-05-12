package main

import (
	"ModEd/core"

	evalModel "ModEd/eval/model"

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

type LoadEvalCommand struct{}

func (l *LoadEvalCommand) Execute() error {
	mgr := migration.GetInstance()
	db := mgr.DB

	if db == nil {
		var err error
		db, err = mgr.SetPathDB(defaultDBPath).
			MigrateModule(core.MODULE_QUIZ).
			MigrateModule(core.MODULE_COMMON).
			BuildDB()
		if err != nil {
			return fmt.Errorf("failed to initialize database: %v", err)
		}
	}
	fmt.Println("Deleting existing progress records...")
	if err := db.Exec("DELETE FROM progresses").Error; err != nil {
		return fmt.Errorf("failed to delete existing progress records: %v", err)
	}

	fmt.Println("Loading seed data from path: ../../data/quiz/Progress.csv")
	var progresses []evalModel.Progress

	mgr.AddSeedData("../../data/quiz/Progress.csv", &progresses)

	err := mgr.LoadSeedData()
	if err != nil {
		return err
	}

	var count int64
	db.Model(&evalModel.Progress{}).Count(&count)
	fmt.Printf("Successfully loaded %d progress records\n", count)

	fmt.Println("Seed data loaded successfully.")
	return nil
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
		MigrateModule(core.MODULE_COMMON).
		BuildDB()

	if err != nil {
		panic(err)
	}

	evaluationController := controller.NewEvaluationController(db)
	progressController := controller.NewProgressController(db)
	assessmentController := controller.NewAssessmentController(db)

	CommandExecutor := NewCommandExecutor()
	CommandExecutor.RegisterCommand("1", &EvaluationCommand{db, evaluationController, progressController, assessmentController})
	CommandExecutor.RegisterCommand("resetdb", &ResetDBCommand{})
	CommandExecutor.RegisterCommand("loadeval", &LoadEvalCommand{})

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
	fmt.Println("0. Exit")
	fmt.Println("'resetdb' to re-initialize the database")
	fmt.Println("'loadeval' to load evaluation seed data")
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
		BuildDB()

	if err != nil {
		return err
	}

	return nil
}
