package cli

import (
	controllerProgress "ModEd/eval/controller"

	util "ModEd/eval/util"

	"fmt"

	"gorm.io/gorm"
)

func RunProgressCLI(db *gorm.DB) {
	ProgressController := controllerProgress.NewProgressController(db)

	for {
		fmt.Println("\nProgress CLI")
		fmt.Println("1. Get All Progress")
		fmt.Println("2. Get Progress By Student Code")
		fmt.Println("3. Get Progress By Status")
		fmt.Println("4. Get Submit Count By AssignmentID")
		fmt.Print("Enter your choice: ")

		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:
			DisplayAllProgress(ProgressController)
		case 2:
			DisplayProgressByStudentCode(ProgressController)
		case 3:
			DisplayProgressByStatus(ProgressController)
		case 4:
			DisplaySubmitCountByAsssignmentID(ProgressController)
		case 5:
			fmt.Println("Exit")
			return
		default:
			fmt.Println("This Number isn't included in choices.")
		}
	}
}

func DisplayAllProgress(controller *controllerProgress.ProgressController) {
	AssignmentId, err := util.PromptUint("Enter Assignment ID: ")
	if err != nil {
		fmt.Println("This Assignment ID doesn't exist.", err)
		return
	}

	progressList, err := controller.GetAllProgress(uint(AssignmentId))
	if err != nil {
		fmt.Println("Error fetching datas", err)
		return
	}

	fmt.Println("\nAll students assignment progress.")
	fmt.Println("-----------------------------------------------------")
	for _, p := range progressList {
		fmt.Printf("Student ID: %s | Assignment ID: %d | Assignment Title: %s | Status: %s | Last Update: %v",
			p.StudentCode.StudentCode, p.AssignmentId.AssignmentId, p.Title.Title, p.Status.Status, p.LastUpdate.Format("2006-01-02"))
	}
}

func DisplayProgressByStudentCode(controller *controllerProgress.ProgressController) {
	AssignmentId, err := util.PromptUint("Enter Assignment ID: ")
	StudentCode := util.PromptString("Enter Student ID: ")

	if err != nil {
		fmt.Println("Assignment ID doesn't exist.", err)
		return
	}

	progressList, err := controller.GetProgressByStudentCode(uint(AssignmentId), StudentCode)
	if err != nil {
		fmt.Println("Error fetching data.", err)
		return
	}

	if len(progressList) == 0 {
		fmt.Println("Student ID doesn't exist.:", StudentCode)
		return
	}

	fmt.Println("\nDisplaying progress of the selected student for the given Assignment ID:")
	fmt.Println("-----------------------------------------------------")
	for _, p := range progressList {
		fmt.Printf("Student ID: %s | Assignment ID: %d | Assignment Title: %s | Status: %s | Last Update: %v",
			p.StudentCode.StudentCode, p.AssignmentId.AssignmentId, p.Title.Title, p.Status.Status, p.LastUpdate.Format("2006-01-02"))
	}
}

func DisplayProgressByStatus(controller *controllerProgress.ProgressController) {
	AssignmentId, err := util.PromptUint("Enter Assignment ID: ")
	Status := util.PromptString("Enter student assignment Status: ")

	if err != nil {
		fmt.Println("Assignment ID doesn't exist.", err)
		return
	}

	progressList, err := controller.GetProgressByStatus(uint(AssignmentId), Status)
	if err != nil {
		fmt.Println("Error fetching data.", err)
		return
	}

	if len(progressList) == 0 {
		fmt.Println("No progress found for status.", Status)
		return
	}

	fmt.Printf("\nDisplaying progress for Assignment ID %d with status '%s':\n", AssignmentId, Status)
	fmt.Println("-----------------------------------------------------")
	for _, p := range progressList {
		fmt.Printf("Student ID: %s | Assignment ID: %d | Assignment Title: %s | Status: %s | Last Update: %v",
			p.StudentCode.StudentCode, p.AssignmentId.AssignmentId, p.Title.Title, p.Status.Status, p.LastUpdate.Format("2006-01-02"))
	}
}

func DisplaySubmitCountByAsssignmentID(controller *controllerProgress.ProgressController) {
	AssignmentId, err := util.PromptUint("Enter Assignment ID: ")

	if err != nil {
		fmt.Println("This Assignment ID doesn't exist.", err)
		return
	}

	count, err := controller.GetSubmitCountByAssignmentID(uint(AssignmentId))
	if err != nil {
		fmt.Println("Error", err)
		return
	}

	fmt.Printf("Total Submission for Assignment %d is %d.", AssignmentId, count)
}
