package cli

import (
	controllerProgress "ModEd/eval/controller"

	evalModel "ModEd/eval/model"

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
			DisplayAllProgress(controller)
		case 2:
			DisplayProgressByStudentCode(controller)
		case 3:
			DisplayProgressByStatus(controller)
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
			DisplayAllProgress(controller)
		case 2:
			DisplayProgressByStudentCode(controller)
		case 3:
			DisplayProgressByStatus(controller)
		case 4:
			DisplayQuizSubmitCount(controller)
		case 5:
			return
		default:
			fmt.Println("Invalid choice")
		}
	}
}

func getAssessmentTypeAndID() (evalModel.AssessmentType, uint, error) {
	fmt.Println("\nSelect Assessment Type:")
	fmt.Println("1. Assignment")
	fmt.Println("2. Quiz")
	fmt.Print("Enter your choice (1 or 2): ")

	var typeChoice int
	fmt.Scan(&typeChoice)

	var assessmentType evalModel.AssessmentType
	var promptMessage string

	switch typeChoice {
	case 1:
		assessmentType = evalModel.AssignmentType
		promptMessage = "Enter Assignment ID: "
	case 2:
		assessmentType = evalModel.QuizType
		promptMessage = "Enter Quiz ID: "
	default:
		return "", 0, fmt.Errorf("invalid choice")
	}

	id, err := util.PromptUint(promptMessage)
	if err != nil {
		return "", 0, fmt.Errorf("invalid ID: %v", err)
	}

	return assessmentType, uint(id), nil
}

func DisplayAllProgress(controller *controllerProgress.ProgressController) {
	assessmentType, id, err := getAssessmentTypeAndID()
	if err != nil {
		fmt.Println(err)
		return
	}

	progressList, err := controller.GetAllProgressByType(assessmentType, id)
	if err != nil {
		fmt.Println("Error fetching data:", err)
		return
	}
	DisplayProgressList(progressList)
}

func DisplayProgressByStudentCode(controller *controllerProgress.ProgressController) {
	assessmentType, id, err := getAssessmentTypeAndID()
	if err != nil {
		fmt.Println(err)
		return
	}

	studentCode := util.PromptString("Enter Student Code: ")

	progressList, err := controller.GetProgressByStudentCode(assessmentType, id, studentCode)
	if err != nil {
		fmt.Println("Error fetching data:", err)
		return
	}

	if len(progressList) == 0 {
		fmt.Println("No progress found for student:", studentCode)
		return
	}

	DisplayProgressList(progressList)
}

func DisplayProgressByStatus(controller *controllerProgress.ProgressController) {
	assessmentType, id, err := getAssessmentTypeAndID()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Status options: submitted, unsubmitted")
	status := util.PromptString("Enter status: ")

	validStatuses := map[string]bool{
		"submitted":   true,
		"unsubmitted": true,
	}

	if !validStatuses[status] {
		fmt.Println("Invalid status. Must be one of: submitted, unsubmitted")
		return
	}

	progressList, err := controller.GetProgressByStatus(assessmentType, id, evalModel.AssessmentStatus(status))
	if err != nil {
		fmt.Println("Error fetching data:", err)
		return
	}

	if len(progressList) == 0 {
		fmt.Println("No progress found with status:", status)
		return
	}

	DisplayProgressList(progressList)
}

func DisplayProgressList(progressList []controllerProgress.Progress) {
	if len(progressList) == 0 {
		fmt.Println("No progress data found")
		return
	}

	fmt.Println("\nProgress List:")
	fmt.Println("-----------------------------------------------------")
	for _, p := range progressList {
		fmt.Printf("Student Code: %s\n", p.StudentCode)
		fmt.Printf("Assessment ID: %d | Status: %s\n", p.AssessmentId, p.Status)
		fmt.Printf("Last Update: %v | Total Submit: %d\n",
			p.LastUpdate.Format("2006-01-02 15:04:05"), p.TotalSubmit)
		fmt.Println()
	}
}

func DisplayAssignmentSubmitCount(controller *controllerProgress.ProgressController) {
	id, err := util.PromptUint("Enter Assignment ID: ")
	if err != nil {
		fmt.Println("Invalid Assignment ID:", err)
		return
	}

	count, err := controller.GetSubmitCount(uint(id))
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

	count, err := controller.GetSubmitCount(uint(id))
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Printf("Total submissions for Quiz %d: %d\n", id, count)
}
