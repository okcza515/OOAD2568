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

// package cli

// import (
// 	"ModEd/eval/controller"
// 	"fmt"
// 	"time"
// )

// func RunEvaluationCLI(evalController *controller.EvaluationController) {
// 	for {
// 		fmt.Println("\nEvaluation Menu")
// 		fmt.Println("1. Evaluate Assignment")
// 		fmt.Println("2. Comment on Assignment")
// 		fmt.Println("3. Evaluate Quiz")
// 		fmt.Println("4. Display Scores")
// 		fmt.Println("0. Exit")

// 		var choice string
// 		fmt.Print("Enter choice: ")
// 		fmt.Scanln(&choice)

// 		switch choice {
// 		case "1":
// 			evaluateAssignment(evalController)
// 		/*case "2":
// 		commentAssignment(evalController)*/
// 		case "3":
// 			evaluateQuiz(evalController)
// 		case "4":
// 			displayScores(evalController)
// 		case "0":
// 			return
// 		default:
// 			fmt.Println("Invalid choice")
// 		}
// 	}
// }

// func evaluateAssignment(ec *controller.EvaluationController) {
// 	var studentCode, instructorCode, comment string
// 	var assignmentId, score uint

// 	fmt.Print("Enter Student Code: ")
// 	fmt.Scanln(&studentCode)
// 	fmt.Print("Enter Instructor Code: ")
// 	fmt.Scanln(&instructorCode)
// 	fmt.Print("Enter Assignment ID: ")
// 	fmt.Scanln(&assignmentId)
// 	fmt.Print("Enter Score: ")
// 	fmt.Scanln(&score)

// 	err := ec.EvaluateAssignment(studentCode, instructorCode, assignmentId, score, comment)
// 	if err != nil {
// 		fmt.Println("Error:", err)
// 	} else {
// 		fmt.Println("Evaluation saved.")
// 	}
// }

/*func commentAssignment(ec *controller.EvaluationController) {
	var studentCode string
	var assignmentId uint
	var comment string

	fmt.Print("Enter Student Code: ")
	fmt.Scanln(&studentCode)
	fmt.Print("Enter Assignment ID: ")
	fmt.Scanln(&assignmentId)
	fmt.Print("Enter Comment: ")
	fmt.Scanln(&comment)

	err := ec.CommentAssignment(studentCode, assignmentId, comment)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Comment saved.")
	}
}*/

// func evaluateQuiz(ec *controller.EvaluationController) {
// 	var studentCode, instructorCode, comment string
// 	var quizId, score uint

// 	fmt.Print("Enter Student Code: ")
// 	fmt.Scanln(&studentCode)
// 	fmt.Print("Enter Instructor Code: ")
// 	fmt.Scanln(&instructorCode)
// 	fmt.Print("Enter Quiz ID: ")
// 	fmt.Scanln(&quizId)
// 	fmt.Print("Enter Score: ")
// 	fmt.Scanln(&score)

// 	err := ec.EvaluateQuiz(studentCode, instructorCode, quizId, score, comment)
// 	if err != nil {
// 		fmt.Println("Error:", err)
// 	} else {
// 		fmt.Println("Evaluation saved.")
// 	}
// }

// func displayScores(ec *controller.EvaluationController) {
// 	var studentCode string
// 	evals, err := ec.ListEvaluations(studentCode)
// 	if err != nil {
// 		fmt.Println("Error retrieving evaluations:", err)
// 		return
// 	}

// 	for _, e := range evals {
// 		kind := "Assignment"
// 		id := e.AssignmentId
// 		if e.QuizId != nil {
// 			kind = "Quiz"
// 			id = e.QuizId
// 		}
// 		fmt.Printf("Student: %s | %s ID: %d | Score: %d | Comment: %s | EvaluatedAt: %s\n",
// 			e.StudentCode,
// 			kind,
// 			id,
// 			e.Score,
// 			e.Comment,
// 			e.EvaluatedAt.Format(time.RFC1123),
// 		)
// 	}
// }

// package cli

// import (
// 	commonModel "ModEd/common/model"
// 	evalModel "ModEd/eval/model"
// 	"fmt"
// 	"time"
// 	//"gorm.io/gorm"
// )

// func PrintEvaluation() {
// 	fmt.Println("Hello from evaluation")
// }

// type Evaluation struct {
// 	StudentCode    commonModel.Student
// 	InstructorCode commonModel.Instructor
// 	AssignmentId   evalModel.Assignment
// 	QuizId         evalModel.Quiz
// 	Score          uint
// 	Comment        string
// 	EvaluatedAt    time.Time
// }

// func RecordEvaluation() {
// 	var studentCode, instructorCode, comment string
// 	var score, assignmentId, quizId uint

// 	fmt.Print("Enter Student Code: ")
// 	fmt.Scanln(&studentCode)

// 	fmt.Print("Enter Instructor Code: ")
// 	fmt.Scanln(&instructorCode)

// 	fmt.Print("Enter Assignment ID: ")
// 	fmt.Scanln(&assignmentId)

// 	fmt.Print("Enter Quiz ID: ")
// 	fmt.Scanln(&quizId)

// 	fmt.Print("Enter Score: ")
// 	fmt.Scanln(&score)

// 	fmt.Print("Enter Comment: ")
// 	fmt.Scanln(&comment)

// 	evaluation := Evaluation{
// 		StudentCode:    commonModel.Student{StudentCode: studentCode},
// 		InstructorCode: commonModel.Instructor{InstructorCode: instructorCode},
// 		AssignmentId:   evalModel.Assignment{AssignmentId: assignmentId},
// 		QuizId:         evalModel.Quiz{QuizId: quizId},
// 		Score:          score,
// 		Comment:        comment,
// 		EvaluatedAt:    time.Now(),
// 	}

// 	fmt.Println("=== evaluation Info ===")
// 	fmt.Println("Student Code:", evaluation.StudentCode.StudentCode)
// 	fmt.Println("Instructor Code:", evaluation.InstructorCode.InstructorCode)
// 	fmt.Println("Assignment ID:", evaluation.AssignmentId.AssignmentId)
// 	fmt.Println("Quiz ID:", evaluation.QuizId.QuizId)
// 	fmt.Println("Score", evaluation.Score)
// 	fmt.Println("Comment", evaluation.Comment)
// 	fmt.Println("Evaluated At:", evaluation.EvaluatedAt.Format("2006-01-02 15:04:05"))
// }
