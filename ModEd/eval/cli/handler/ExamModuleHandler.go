package handler

import (
	"ModEd/asset/util"
	"ModEd/core"
	"ModEd/core/cli"
	"ModEd/core/handler"
	"ModEd/core/migration"
	"ModEd/eval/controller"
	"fmt"
)

const (
	defaultDBPath = "../../data/ModEd.bin"
)

type ExamModuleMenuStateHandler struct {
	Manager *cli.CLIMenuStateManager
	wrapper *controller.ExamModuleWrapper

	ExamMenuStateHandler	   *ExamMenuStateHandler
	QuestionMenuStateHandler   *QuestionMenuStateHandler
	SubmissionMenuStateHandler *SubmissionMenuStateHandler
	handler                    *handler.HandlerContext
}

type LoadExamCommand struct {
	manager  *cli.CLIMenuStateManager
	mainMenu cli.MenuState
}

func NewLoadExamCommand(manager *cli.CLIMenuStateManager, mainMenu cli.MenuState) *LoadExamCommand {
	return &LoadExamCommand{
		manager:  manager,
		mainMenu: mainMenu,
	}
}

func (l *LoadExamCommand) Execute() error {
	mgr := migration.GetInstance()
	db := mgr.DB

	if db == nil {
		var err error
		db, err = mgr.SetPathDB(defaultDBPath).
			MigrateModule(core.MODULE_COMMON).
			MigrateModule(core.MODULE_CURRICULUM).
			BuildDB()
		if err != nil {
			return fmt.Errorf("failed to initialize database: %v", err)
		}
	}
	// fmt.Println("Deleting existing progress records...")
	// if err := db.Exec("DELETE FROM progresses").Error; err != nil {
	// 	return fmt.Errorf("failed to delete existing progress records: %v", err)
	// }

	// fmt.Println("Loading seed data from path: ../../data/quiz/Progress.csv")
	// var progresses []evalModel.Progress

	// mgr.AddSeedData("../../data/quiz/Progress.csv", &progresses)

	err := mgr.LoadSeedData()
	if err != nil {
		return err
	}

	// var count int64
	// db.Model(&evalModel.Progress{}).Count(&count)
	// fmt.Printf("Successfully loaded %d progress records\n", count)

	fmt.Println("Seed data loaded successfully.")
	return nil
}

// Render implements cli.MenuState interface
func (l *LoadExamCommand) Render() {
	fmt.Println("\nLoading Seed Data:")
	fmt.Println("This will load progress seed data from CSV files.")
	fmt.Println("Press Enter to return to the main menu.")
}

// HandleUserInput implements cli.MenuState interface
func (l *LoadExamCommand) HandleUserInput(input string) error {
	err := l.Execute()
	if err != nil {
		fmt.Printf("Error loading seed data: %v\n", err)
	}
	util.PressEnterToContinue()
	l.manager.SetState(l.mainMenu)
	return nil
}

func NewExamModuleHandler(manager *cli.CLIMenuStateManager, wrapper *controller.ExamModuleWrapper) *ExamModuleMenuStateHandler {
	examModuleHandler := &ExamModuleMenuStateHandler{
		Manager: manager,
		wrapper: wrapper,
		handler: handler.NewHandlerContext(),
	}

	examModuleHandler.ExamMenuStateHandler = NewExamMenuStateHandler(manager, wrapper, examModuleHandler)
	examModuleHandler.QuestionMenuStateHandler = NewQuestionMenuStateHandler(manager, wrapper, examModuleHandler)
	examModuleHandler.SubmissionMenuStateHandler = NewSubmissionMenuStateHandler(manager, wrapper, examModuleHandler)

	examModuleHandler.Manager.AddMenu("1", examModuleHandler.ExamMenuStateHandler)
	examModuleHandler.Manager.AddMenu("2", examModuleHandler.QuestionMenuStateHandler)
	examModuleHandler.Manager.AddMenu("3", examModuleHandler.SubmissionMenuStateHandler)
	loadCmd := NewLoadExamCommand(manager, examModuleHandler)
	examModuleHandler.Manager.AddMenu("Load", loadCmd)
	examModuleHandler.Manager.AddMenu("Exit", nil)

	return examModuleHandler
}

func (examHandler *ExamModuleMenuStateHandler) Render() {
	examHandler.handler.SetMenuTitle("\nExamination Module Menu:")
	examHandler.handler.AddHandler("1", "Exam", handler.FuncStrategy{})
	examHandler.handler.AddHandler("2", "Question", handler.FuncStrategy{})
	examHandler.handler.AddHandler("3", "Submission", handler.FuncStrategy{})
	examHandler.handler.AddHandler("Load", "Load Seed Data", handler.FuncStrategy{})
	examHandler.handler.AddHandler("Exit", "Exit the Evaluation Module", handler.FuncStrategy{})
	examHandler.handler.ShowMenu()
}

func (handler *ExamModuleMenuStateHandler) HandleUserInput(input string) error {
	err := handler.Manager.GoToMenu(input)

	if err != nil {
		fmt.Println("Invalid input. Please try again.")
		util.PressEnterToContinue()
	}

	return nil
}
