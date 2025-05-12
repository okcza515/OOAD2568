//MEP-1006 Quiz and Assignment

package handler

import (
	"ModEd/asset/util"
	"ModEd/core/cli"
	"ModEd/core/handler"
	"ModEd/eval/controller"
	"ModEd/eval/model"
	"fmt"
)

type EvaluationMenuStateHandler struct {
	Manager                    *cli.CLIMenuStateManager
	wrapper                    *controller.EvalModuleWrapper
	EvalModuleMenuStateHandler cli.MenuState
	handler                    *handler.HandlerContext
	backhandler                *handler.ChangeMenuHandlerStrategy
}

func NewEvaluationMenuStateHandler(manager *cli.CLIMenuStateManager, wrapper *controller.EvalModuleWrapper, evalModuleMenuStateHandler cli.MenuState) *EvaluationMenuStateHandler {
	return &EvaluationMenuStateHandler{
		Manager:                    manager,
		wrapper:                    wrapper,
		EvalModuleMenuStateHandler: evalModuleMenuStateHandler,
		handler:                    handler.NewHandlerContext(),
		backhandler:                handler.NewChangeMenuHandlerStrategy(manager, evalModuleMenuStateHandler),
	}
}

func (menu *EvaluationMenuStateHandler) Render() {
	util.ClearScreen()
	menu.handler.SetMenuTitle("\nEvaluation Menu")
	menu.handler.AddHandler("1", "Evaluation Assignment", handler.FuncStrategy{Action: menu.CreateEvaluation})
	menu.handler.AddHandler("2", "View All Evaluations", handler.FuncStrategy{Action: menu.ViewAllEvaluations})
	menu.handler.AddHandler("3", "View Evaluation by StudentID", handler.FuncStrategy{Action: menu.ViewEvaluationByID})
	menu.handler.AddHandler("4", "Update Evaluation", handler.FuncStrategy{Action: menu.UpdateEvaluation})
	menu.handler.AddBackHandler(menu.backhandler)
	menu.handler.ShowMenu()
}

func (menu *EvaluationMenuStateHandler) HandleUserInput(input string) error {
	return menu.handler.HandleInput(input)
}

func (menu *EvaluationMenuStateHandler) CreateEvaluation() error {
	var studentCode, instructorCode, comment string
	var assignmentId, score uint

	fmt.Print("Enter Student Code: ")
	fmt.Scanln(&studentCode)
	fmt.Print("Enter Instructor Code: ")
	fmt.Scanln(&instructorCode)
	fmt.Print("Enter Assignment ID: ")
	fmt.Scanln(&assignmentId)
	fmt.Print("Enter Score: ")
	fmt.Scanln(&score)
	fmt.Print("Enter Comment: ")
	fmt.Scanln(&comment)

	err := menu.wrapper.EvaluationController.CreateEvaluation(studentCode, instructorCode, assignmentId, score, comment)
	if err != nil {
		fmt.Printf("Error creating evaluation: %v\n", err)
		return err
	}

	fmt.Println("Evaluation created successfully!")
	util.PressEnterToContinue()
	util.ClearScreen()
	return nil
}

func (menu *EvaluationMenuStateHandler) ViewAllEvaluations() error {
	evaluations, err := menu.wrapper.EvaluationController.ViewAllEvaluations()
	if err != nil {
		fmt.Printf("Error viewing evaluations: %v\n", err)
		return err
	}

	if len(evaluations) == 0 {
		fmt.Println("No evaluations found.")
		util.PressEnterToContinue()
		util.ClearScreen()
		return nil
	}

	menu.displayEvaluationTableHeader()
	for _, eval := range evaluations {
		menu.displayEvaluation(eval)
	}
	util.PressEnterToContinue()
	util.ClearScreen()
	return nil
}

func (menu *EvaluationMenuStateHandler) ViewEvaluationByID() error {
	var studentCode string
	fmt.Print("Enter Student Code: ")
	fmt.Scanln(&studentCode)

	evaluations, err := menu.wrapper.EvaluationController.ViewEvaluationByID(studentCode)
	if err != nil {
		fmt.Printf("Error viewing evaluations: %v\n", err)
		return err
	}

	if len(evaluations) == 0 {
		fmt.Println("No evaluations found for this student.")
		util.PressEnterToContinue()
		util.ClearScreen()
		return nil
	}

	menu.displayEvaluationTableHeader()
	for _, eval := range evaluations {
		menu.displayEvaluation(eval)
	}
	util.PressEnterToContinue()
	util.ClearScreen()
	return nil
}

func (menu *EvaluationMenuStateHandler) UpdateEvaluation() error {
	var id uint
	var score uint
	var comment string

	fmt.Print("Enter Evaluation ID: ")
	fmt.Scanln(&id)
	fmt.Print("Enter new Score: ")
	fmt.Scanln(&score)
	fmt.Print("Enter new Comment: ")
	fmt.Scanln(&comment)

	err := menu.wrapper.EvaluationController.UpdateEvaluation(id, score, comment)
	if err != nil {
		fmt.Printf("Error updating evaluation: %v\n", err)
		return err
	}

	fmt.Println("Evaluation updated successfully!")
	util.PressEnterToContinue()
	util.ClearScreen()
	return nil
}

func (menu *EvaluationMenuStateHandler) displayEvaluationTableHeader() {
	fmt.Printf("\n%-5s %-15s %-15s %-10s %-10s %-20s %-20s",
		"ID", "Student Code", "Instructor Code", "Score", "Assignment", "Comment", "Evaluated At")
	fmt.Printf("\n%-5s %-15s %-15s %-10s %-10s %-20s %-20s",
		"---", "------------", "--------------", "-----", "----------", "-------", "------------")
}

func (menu *EvaluationMenuStateHandler) displayEvaluation(eval interface{}) {
	evaluation := eval.(model.Evaluation)
	fmt.Printf("\n%-5d %-15s %-15s %-10d %-10d %-20s %-20s",
		evaluation.ID,
		evaluation.StudentCode,
		evaluation.InstructorCode,
		evaluation.Score,
		evaluation.AssignmentId,
		evaluation.Comment,
		evaluation.EvaluatedAt.Format("2006-01-02 15:04:05"))
}
