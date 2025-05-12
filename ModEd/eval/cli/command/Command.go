package command

import (
	"ModEd/core"
	"ModEd/core/handler"
	"ModEd/core/migration"
	"ModEd/eval/cli/evaluation/menu"
	"ModEd/eval/controller"
	"fmt"

	"gorm.io/gorm"
)

const defaultDBPath = "../../data/ModEd.bin"

// Command interface defines the contract for all commands
type Command interface {
	Execute() error
}

// CommandExecutor handles the registration and execution of commands
type CommandExecutor struct {
	commands map[string]Command
}

// NewCommandExecutor creates a new CommandExecutor instance
func NewCommandExecutor() *CommandExecutor {
	return &CommandExecutor{commands: make(map[string]Command)}
}

// RegisterCommand registers a new command with the executor
func (ce *CommandExecutor) RegisterCommand(name string, command Command) {
	ce.commands[name] = command
}

// ExecuteCommand executes a command by its name
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

// EvaluationCommand handles the evaluation module functionality
type EvaluationCommand struct {
	DB                   *gorm.DB
	EvaluationController *controller.EvaluationController
	ProgressController   *controller.ProgressController
	AssignmentController *controller.AssignmentController
}

// Execute runs the evaluation module CLI
func (e *EvaluationCommand) Execute() error {
	menu.RunEvalModuleCLI(e.DB, e.EvaluationController, e.ProgressController, e.AssignmentController)
	return nil
}

// ResetDBCommand handles resetting the database
type ResetDBCommand struct{}

// Execute resets the database
func (r *ResetDBCommand) Execute() error {
	return resetDB()
}

// DisplayMainMenu shows the main menu options
func DisplayMainMenu() {
	menuHandler := handler.NewHandlerContext()
	menuHandler.SetMenuTitle("\nEvaluation Module Menu:")
	menuHandler.AddHandler("1", "Evaluation Assignment & Quiz", handler.FuncStrategy{})
	// menuHandler.AddHandler("2", "Evaluation Examination", handler.FuncStrategy{})
	menuHandler.AddHandler("0", "Exit", handler.FuncStrategy{})
	menuHandler.AddHandler("resetdb", "Re-initialize the database", handler.FuncStrategy{})
	menuHandler.ShowMenu()
}

// GetUserChoice gets the user's menu choice
func GetUserChoice() string {
	var choice string
	fmt.Print("Enter your choice: ")
	fmt.Scanln(&choice)
	return choice
}

// resetDB resets the database to its initial state
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

// InitializeDB initializes the database with required modules
func InitializeDB() (*gorm.DB, error) {
	return migration.
		GetInstance().
		SetPathDB(defaultDBPath).
		MigrateModule(core.MODULE_QUIZ).
		MigrateModule(core.MODULE_COMMON).
		BuildDB()
}

// InitializeControllers creates and returns the necessary controllers
func InitializeControllers(db *gorm.DB) (
	*controller.EvaluationController,
	*controller.ProgressController,
	*controller.AssignmentController,
) {
	return controller.NewEvaluationController(db),
		controller.NewProgressController(db),
		controller.NewAssignmentController(db)
}
