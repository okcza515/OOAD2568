//MEP-1006 Quiz and Assignment

package main

import (
	"ModEd/core"
	"ModEd/core/handler"
	"ModEd/core/migration"
	"ModEd/eval/cli/evaluation/command"
	examinationCommand "ModEd/eval/cli/exam/command"
	controller "ModEd/eval/controller"
	"fmt"
)

const (
	defaultDBPath = "../../data/ModEd.bin"
)

func main() {
	db, err := migration.
		GetInstance().
		SetPathDB(defaultDBPath).
		MigrateModule(core.MODULE_QUIZ).
		MigrateModule(core.MODULE_COMMON).
		MigrateModule(core.MODULE_CURRICULUM).
		MigrateModule(core.MODULE_EVAL).
		BuildDB()

	if err != nil {
		panic(err)
	}

	evaluationController := controller.NewEvaluationController(db)
	progressController := controller.NewProgressController(db)
	assignmentController := controller.NewAssignmentController(db)

	examController := controller.NewExamController(db)
	questionController := controller.NewQuestionController(db)
	submissionController := controller.NewSubmissionController(db)

	commandExecutor := command.NewCommandExecutor()
	commandExecutor.RegisterCommand("1", &command.EvaluationCommand{
		DB:                   db,
		EvaluationController: evaluationController,
		ProgressController:   progressController,
		AssignmentController: assignmentController,
	})
	commandExecutor.RegisterCommand("2", &examinationCommand.ExaminationCommand{
		DB:                   db,
		ExamController:       examController,
		QuestionController:   questionController,
		SubmissionController: submissionController,
	})
	commandExecutor.RegisterCommand("resetdb", &command.ResetDBCommand{})

	for {
		DisplayMainMenu()
		choice := GetUserChoice()

		if choice == "0" {
			fmt.Println("Exiting...")
			return
		}

		if err := commandExecutor.ExecuteCommand(choice); err != nil {
			fmt.Println("Error executing command:", err)
		}
	}
}

func DisplayMainMenu() {
	menuHandler := handler.NewHandlerContext()
	menuHandler.SetMenuTitle("\nEvaluation Module Menu:")
	menuHandler.AddHandler("1", "Evaluation Assignment & Quiz", handler.FuncStrategy{})
	menuHandler.AddHandler("2", "Exam Question & Submission", handler.FuncStrategy{})
	menuHandler.AddHandler("0", "Exit", handler.FuncStrategy{})
	menuHandler.AddHandler("resetdb", "Re-initialize the database", handler.FuncStrategy{})
	menuHandler.ShowMenu()
}

func GetUserChoice() string {
	var choice string
	fmt.Print("Enter your choice: ")
	fmt.Scanln(&choice)
	return choice
}