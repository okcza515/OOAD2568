package examination

import (
	"ModEd/core/handler"
	"ModEd/eval/controller"
	"fmt"
)

type ExaminationCLIParams struct {
	ExaminationController *controller.ExaminationController
	SectionController     *controller.ExamSectionController
}

func RunExaminationCLI(params *ExaminationCLIParams) {
	ctx := handler.NewHandlerContext()
	ctx.SetMenuTitle("Examination Management")

	// Examination menu items
	ctx.AddHandler("1", "List Examinations", handler.FuncStrategy{
		Action: func() error {
			fmt.Println("\n===== Examinations =====")
			// Implementation would call params.ExaminationController methods
			fmt.Println("Examination listing will be implemented here")
			return nil
		},
	})

	ctx.AddHandler("2", "Create Examination", handler.FuncStrategy{
		Action: func() error {
			fmt.Println("\n===== Create Examination =====")
			// Implementation would call params.ExaminationController methods
			fmt.Println("Examination creation will be implemented here")
			return nil
		},
	})

	ctx.AddHandler("3", "Update Examination", handler.FuncStrategy{
		Action: func() error {
			fmt.Println("\n===== Update Examination =====")
			// Implementation would call params.ExaminationController methods
			fmt.Println("Examination update will be implemented here")
			return nil
		},
	})

	ctx.AddHandler("4", "Delete Examination", handler.FuncStrategy{
		Action: func() error {
			fmt.Println("\n===== Delete Examination =====")
			// Implementation would call params.ExaminationController methods
			fmt.Println("Examination deletion will be implemented here")
			return nil
		},
	})

	// Sections menu items
	ctx.AddHandler("5", "Manage Exam Sections", handler.FuncStrategy{
		Action: func() error {
			manageExamSections(params)
			return nil
		},
	})

	ctx.AddBackHandler(handler.FuncStrategy{
		Action: func() error {
			return nil
		},
	})

	for {
		ctx.ShowMenu()
		var userInput string
		fmt.Print("Enter your choice: ")
		fmt.Scanln(&userInput)
		if err := ctx.HandleInput(userInput); err != nil {
			fmt.Println("Error:", err)
		}
		if userInput == "back" {
			break
		}
	}
}

func manageExamSections(params *ExaminationCLIParams) {
	ctx := handler.NewHandlerContext()
	ctx.SetMenuTitle("Exam Sections Management")

	ctx.AddHandler("1", "List Sections", handler.FuncStrategy{
		Action: func() error {
			fmt.Println("\n===== Exam Sections =====")
			// Implementation would call params.SectionController methods
			fmt.Println("Section listing will be implemented here")
			return nil
		},
	})

	ctx.AddHandler("2", "Add Section", handler.FuncStrategy{
		Action: func() error {
			fmt.Println("\n===== Add Section =====")
			// Implementation would call params.SectionController methods
			fmt.Println("Section addition will be implemented here")
			return nil
		},
	})

	ctx.AddHandler("3", "Update Section", handler.FuncStrategy{
		Action: func() error {
			fmt.Println("\n===== Update Section =====")
			// Implementation would call params.SectionController methods
			fmt.Println("Section update will be implemented here")
			return nil
		},
	})

	ctx.AddHandler("4", "Delete Section", handler.FuncStrategy{
		Action: func() error {
			fmt.Println("\n===== Delete Section =====")
			// Implementation would call params.SectionController methods
			fmt.Println("Section deletion will be implemented here")
			return nil
		},
	})

	ctx.AddBackHandler(handler.FuncStrategy{
		Action: func() error {
			return nil
		},
	})

	for {
		ctx.ShowMenu()
		var userInput string
		fmt.Print("Enter your choice: ")
		fmt.Scanln(&userInput)
		if err := ctx.HandleInput(userInput); err != nil {
			fmt.Println("Error:", err)
		}
		if userInput == "back" {
			break
		}
	}
}
