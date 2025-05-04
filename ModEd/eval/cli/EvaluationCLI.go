//Chanawat Limpanatewin 65070503445
//MEP-1006

// //1. evaluate in assignment
// //2. comment in assignment
// //3. evaluate in quiz
// //4. diplay score assignment/quiz

package cli

import (
	"ModEd/eval/controller"
	"fmt"
	"time"
)

func RunEvaluationCLI(ec *controller.EvaluationController) {
	for {
		fmt.Println("\nEvaluation Menu")
		fmt.Println("1. Evaluate Assignment")
		fmt.Println("2. Comment on Assignment")
		fmt.Println("3. Evaluate Quiz")
		fmt.Println("4. Display Evaluations")
		fmt.Println("0. Exit")

		var choice string
		fmt.Print("Enter choice: ")
		fmt.Scanln(&choice)

		switch choice {
		case "1":
			evaluateAssignment(ec)
		case "2":
			commentAssignment(ec)
		case "3":
			evaluateQuiz(ec)
		case "4":
			displayEvaluations(ec)
		case "0":
			return
		default:
			fmt.Println("Invalid choice")
		}
	}
}

func evaluateAssignment(ec *controller.EvaluationController) {
	var studentCode, instructorCode string
	var assignmentID, score uint

	fmt.Print("Student Code: ")
	fmt.Scanln(&studentCode)
	fmt.Print("Instructor Code: ")
	fmt.Scanln(&instructorCode)
	fmt.Print("Assignment ID: ")
	fmt.Scanln(&assignmentID)
	fmt.Print("Score: ")
	fmt.Scanln(&score)

	ec.EvaluateAssignment(studentCode, instructorCode, assignmentID, score)
	fmt.Println("Assignment evaluated successfully.")
}

func commentAssignment(ec *controller.EvaluationController) {
	var studentCode string
	var assignmentID uint
	var comment string

	fmt.Print("Student Code: ")
	fmt.Scanln(&studentCode)
	fmt.Print("Assignment ID: ")
	fmt.Scanln(&assignmentID)
	fmt.Print("Comment: ")
	fmt.Scanln(&comment)

	ec.CommentAssignment(studentCode, assignmentID, comment)
	fmt.Println("Comment added successfully.")
}

func evaluateQuiz(ec *controller.EvaluationController) {
	var studentCode, instructorCode string
	var quizID, score uint

	fmt.Print("Student Code: ")
	fmt.Scanln(&studentCode)
	fmt.Print("Instructor Code: ")
	fmt.Scanln(&instructorCode)
	fmt.Print("Quiz ID: ")
	fmt.Scanln(&quizID)
	fmt.Print("Score: ")
	fmt.Scanln(&score)

	ec.EvaluateQuiz(studentCode, instructorCode, quizID, score)
	fmt.Println("Quiz evaluated successfully.")
}

func displayEvaluations(ec *controller.EvaluationController) {
	for _, e := range ec.ListEvaluations() {
		if e.AssignmentID != nil {
			fmt.Printf("Assignment Evaluation: Student %s Assignment ID %d Score %d Comment %s At %s\n",
				e.StudentCode, *e.AssignmentID, e.Score, e.Comment, e.EvaluatedAt.Format(time.RFC1123))
		}
		if e.QuizID != nil {
			fmt.Printf("Quiz Evaluation: Student %s Quiz ID %d Score %d Comment %s At %s\n",
				e.StudentCode, *e.QuizID, e.Score, e.Comment, e.EvaluatedAt.Format(time.RFC1123))
		}
	}
}
