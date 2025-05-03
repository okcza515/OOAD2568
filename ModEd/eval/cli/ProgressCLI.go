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
		fmt.Println("1. Assignment Progress")
		fmt.Println("2. Quiz Progress")
		fmt.Println("3. Exit")
		fmt.Print("Enter your choice: ")

		var mainChoice int
		fmt.Scan(&mainChoice)

		switch mainChoice {
		case 1:
			handleAssignmentProgress(ProgressController)
		case 2:
			handleQuizProgress(ProgressController)
		case 3:
			fmt.Println("Exit")
			return
		default:
			fmt.Println("Invalid choice")
		}
	}
}

func handleAssignmentProgress(controller *controllerProgress.ProgressController) {
	for {
		fmt.Println("\nAssignment Progress Menu")
		fmt.Println("1. Get All Assignment Progress")
		fmt.Println("2. Get Assignment Progress By Student Code")
		fmt.Println("3. Get Assignment Progress By Status")
		fmt.Println("4. Get Assignment Submit Count")
		fmt.Println("5. Back")
		fmt.Print("Enter your choice: ")

		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:
			DisplayAllProgress(controller, true)
		case 2:
			DisplayProgressByStudentCode(controller, true)
		case 3:
			DisplayProgressByStatus(controller, true)
		case 4:
			DisplayAssignmentSubmitCount(controller)
		case 5:
			return
		default:
			fmt.Println("Invalid choice")
		}
	}
}

func handleQuizProgress(controller *controllerProgress.ProgressController) {
	for {
		fmt.Println("\nQuiz Progress Menu")
		fmt.Println("1. Get All Quiz Progress")
		fmt.Println("2. Get Quiz Progress By Student Code")
		fmt.Println("3. Get Quiz Progress By Status")
		fmt.Println("4. Get Quiz Submit Count")
		fmt.Println("5. Back")
		fmt.Print("Enter your choice: ")

		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:
			DisplayAllProgress(controller, false)
		case 2:
			DisplayProgressByStudentCode(controller, false)
		case 3:
			DisplayProgressByStatus(controller, false)
		case 4:
			DisplayQuizSubmitCount(controller)
		case 5:
			return
		default:
			fmt.Println("Invalid choice")
		}
	}
}

func DisplayAllProgress(controller *controllerProgress.ProgressController, isAssignment bool) {
	var id uint64
	var err error

	if isAssignment {
		id, err = util.PromptUint("Enter Assignment ID: ")
		if err != nil {
			fmt.Println("Invalid Assignment ID:", err)
			return
		}
		progressList, err := controller.GetAllProgress(uint(id), 0)
		if err != nil {
			fmt.Println("Error fetching data:", err)
			return
		}
		displayProgressList(progressList, true)
	} else {
		id, err = util.PromptUint("Enter Quiz ID: ")
		if err != nil {
			fmt.Println("Invalid Quiz ID:", err)
			return
		}
		progressList, err := controller.GetAllProgress(0, uint(id))
		if err != nil {
			fmt.Println("Error fetching data:", err)
			return
		}
		displayProgressList(progressList, false)
	}
}

func DisplayProgressByStudentCode(controller *controllerProgress.ProgressController, isAssignment bool) {
	var id uint64
	var err error

	if isAssignment {
		id, err = util.PromptUint("Enter Assignment ID: ")
		if err != nil {
			fmt.Println("Invalid Assignment ID:", err)
			return
		}
	} else {
		id, err = util.PromptUint("Enter Quiz ID: ")
		if err != nil {
			fmt.Println("Invalid Quiz ID:", err)
			return
		}
	}

	StudentCode := util.PromptString("Enter Student Code: ")

	var progressList []controllerProgress.Progress
	if isAssignment {
		progressList, err = controller.GetProgressByStudentCode(uint(id), 0, StudentCode)
	} else {
		progressList, err = controller.GetProgressByStudentCode(0, uint(id), StudentCode)
	}

	if err != nil {
		fmt.Println("Error fetching data:", err)
		return
	}

	if len(progressList) == 0 {
		fmt.Println("No progress found for student:", StudentCode)
		return
	}

	displayProgressList(progressList, isAssignment)
}

func DisplayProgressByStatus(controller *controllerProgress.ProgressController, isAssignment bool) {
	var id uint64
	var err error

	if isAssignment {
		id, err = util.PromptUint("Enter Assignment ID: ")
		if err != nil {
			fmt.Println("Invalid Assignment ID:", err)
			return
		}
	} else {
		id, err = util.PromptUint("Enter Quiz ID: ")
		if err != nil {
			fmt.Println("Invalid Quiz ID:", err)
			return
		}
	}

	fmt.Println("Status options: completed, in_progress, not_started")
	Status := util.PromptString("Enter status: ")

	var progressList []controllerProgress.Progress
	if isAssignment {
		progressList, err = controller.GetProgressByStatus(uint(id), 0, Status)
	} else {
		progressList, err = controller.GetProgressByStatus(0, uint(id), Status)
	}

	if err != nil {
		fmt.Println("Error fetching data:", err)
		return
	}

	if len(progressList) == 0 {
		fmt.Println("No progress found with status:", Status)
		return
	}

	displayProgressList(progressList, isAssignment)
}

func displayProgressList(progressList []controllerProgress.Progress, isAssignment bool) {
	fmt.Println("\nProgress List:")
	fmt.Println("-----------------------------------------------------")
	for _, p := range progressList {
		fmt.Printf("Student Code: %s\n", p.StudentCode.StudentCode)
		if isAssignment {
			fmt.Printf("Assignment ID: %d | Status: %s\n", p.AssignmentId, p.AssignmentStatus)
		} else {
			fmt.Printf("Quiz ID: %d | Status: %s\n", p.QuizId, p.QuizStatus)
		}
		fmt.Printf("Last Update: %v | Total Submit: %d\n\n",
			p.LastUpdate.Format("2006-01-02 15:04:05"), p.TotalSubmit)
	}
}

func DisplayAssignmentSubmitCount(controller *controllerProgress.ProgressController) {
	id, err := util.PromptUint("Enter Assignment ID: ")
	if err != nil {
		fmt.Println("Invalid Assignment ID:", err)
		return
	}

	count, err := controller.GetSubmitCountByAssignmentID(uint(id))
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Printf("Total submissions for Assignment %d: %d\n", id, count)
}

func DisplayQuizSubmitCount(controller *controllerProgress.ProgressController) {
	id, err := util.PromptUint("Enter Quiz ID: ")
	if err != nil {
		fmt.Println("Invalid Quiz ID:", err)
		return
	}

	count, err := controller.GetSubmitCountByQuizID(uint(id))
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Printf("Total submissions for Quiz %d: %d\n", id, count)
}
