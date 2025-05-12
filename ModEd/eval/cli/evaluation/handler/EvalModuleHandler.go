package handler

import (
	"ModEd/asset/util"

	"ModEd/core"

	"ModEd/core/cli"

	"ModEd/core/migration"

	"ModEd/eval/controller"

	evalModel "ModEd/eval/model"

	"ModEd/core/handler"

	"fmt"
)

const (
	defaultDBPath = "../../data/ModEd.bin"
)

type EvalModuleMenuStateHandler struct {
	Manager *cli.CLIMenuStateManager
	wrapper *controller.EvalModuleWrapper

	AssignmentMenuStateHandler *AssignmentMenuStateHandler
	ProgressMenuStateHandler   *ProgressMenuStateHandler
	EvaluationMenuStateHandler *EvaluationMenuStateHandler
	handler                    *handler.HandlerContext
}

type LoadEvalCommand struct {
	manager  *cli.CLIMenuStateManager
	mainMenu cli.MenuState
}

func NewLoadEvalCommand(manager *cli.CLIMenuStateManager, mainMenu cli.MenuState) *LoadEvalCommand {
	return &LoadEvalCommand{
		manager:  manager,
		mainMenu: mainMenu,
	}
}

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

// Render implements cli.MenuState interface
func (l *LoadEvalCommand) Render() {
	fmt.Println("\nLoading Seed Data:")
	fmt.Println("This will load progress seed data from CSV files.")
	fmt.Println("Press Enter to return to the main menu.")
}

// HandleUserInput implements cli.MenuState interface
func (l *LoadEvalCommand) HandleUserInput(input string) error {
	err := l.Execute()
	if err != nil {
		fmt.Printf("Error loading seed data: %v\n", err)
	}
	util.PressEnterToContinue()
	util.ClearScreen()
	l.manager.SetState(l.mainMenu)
	return nil
}

func NewEvalModuleHandler(manager *cli.CLIMenuStateManager, wrapper *controller.EvalModuleWrapper) *EvalModuleMenuStateHandler {
	evalModuleHandler := &EvalModuleMenuStateHandler{
		Manager: manager,
		wrapper: wrapper,
		handler: handler.NewHandlerContext(),
	}

	evalModuleHandler.ProgressMenuStateHandler = NewProgressMenuStateHandler(manager, wrapper, evalModuleHandler)
	evalModuleHandler.AssignmentMenuStateHandler = NewAssignmentMenuStateHandler(manager, wrapper, evalModuleHandler)
	evalModuleHandler.EvaluationMenuStateHandler = NewEvaluationMenuStateHandler(manager, wrapper, evalModuleHandler)

	evalModuleHandler.Manager.AddMenu("1", evalModuleHandler.AssignmentMenuStateHandler)
	evalModuleHandler.Manager.AddMenu("2", evalModuleHandler.AssignmentMenuStateHandler)
	evalModuleHandler.Manager.AddMenu("3", evalModuleHandler.ProgressMenuStateHandler)
	evalModuleHandler.Manager.AddMenu("4", evalModuleHandler.EvaluationMenuStateHandler)
	loadCmd := NewLoadEvalCommand(manager, evalModuleHandler)
	evalModuleHandler.Manager.AddMenu("Load", loadCmd)
	evalModuleHandler.Manager.AddMenu("Exit", nil)

	return evalModuleHandler
}

func (evalHandler *EvalModuleMenuStateHandler) Render() {
	util.ClearScreen()
	evalHandler.handler.SetMenuTitle("\nEvaluation Module Menu:")
	evalHandler.handler.AddHandler("1", "Assignment", handler.FuncStrategy{})
	evalHandler.handler.AddHandler("2", "Quiz", handler.FuncStrategy{})
	evalHandler.handler.AddHandler("3", "Progress", handler.FuncStrategy{})
	evalHandler.handler.AddHandler("4", "Evaluation", handler.FuncStrategy{})
	evalHandler.handler.AddHandler("Load", "Load Seed Data", handler.FuncStrategy{})
	evalHandler.handler.AddHandler("Exit", "Exit the Evaluation Module", handler.FuncStrategy{})
	evalHandler.handler.ShowMenu()
}

func (handler *EvalModuleMenuStateHandler) HandleUserInput(input string) error {
	util.ClearScreen()
	err := handler.Manager.GoToMenu(input)

	if err != nil {
		fmt.Println("Invalid input. Please try again.")
		util.PressEnterToContinue()
	}

	return nil
}
