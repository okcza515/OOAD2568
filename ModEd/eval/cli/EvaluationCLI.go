//Chanawat Limpanatewin 65070503445
//MEP-1006

package cli

import (
	commonModel "ModEd/common/model"
	evalModel "ModEd/eval/model"
	"fmt"
	"time"
	//"gorm.io/gorm"
)

func PrintEvaluation() {
	fmt.Println("Hello from evaluation")
}

type Evaluation struct {
	StudentCode    commonModel.Student
	InstructorCode commonModel.Instructor
	AssignmentId   evalModel.Assignment
	QuizId         evalModel.Quiz
	Score          uint
	Comment        string
	EvaluatedAt    time.Time
}

func RecordEvaluation() {
	var studentCode, instructorCode, comment string
	var score, assignmentId, quizId uint

	fmt.Print("Enter Student Code: ")
	fmt.Scanln(&studentCode)

	fmt.Print("Enter Instructor Code: ")
	fmt.Scanln(&instructorCode)

	fmt.Print("Enter Assignment ID: ")
	fmt.Scanln(&assignmentId)

	fmt.Print("Enter Quiz ID: ")
	fmt.Scanln(&quizId)

	fmt.Print("Enter Score: ")
	fmt.Scanln(&score)

	fmt.Print("Enter Comment: ")
	fmt.Scanln(&comment)

	evaluation := Evaluation{
		StudentCode:    commonModel.Student{StudentCode: studentCode},
		InstructorCode: commonModel.Instructor{InstructorCode: instructorCode},
		AssignmentId:   evalModel.Assignment{AssignmentId: assignmentId},
		QuizId:         evalModel.Quiz{QuizId: quizId},
		Score:          score,
		Comment:        comment,
		EvaluatedAt:    time.Now(),
	}

	fmt.Println("=== evaluation Info ===")
	fmt.Println("Student Code:", evaluation.StudentCode.StudentCode)
	fmt.Println("Instructor Code:", evaluation.InstructorCode.InstructorCode)
	fmt.Println("Assignment ID:", evaluation.AssignmentId.AssignmentId)
	fmt.Println("Quiz ID:", evaluation.QuizId.QuizId)
	fmt.Println("Score", evaluation.Score)
	fmt.Println("Comment", evaluation.Comment)
	fmt.Println("Evaluated At:", evaluation.EvaluatedAt.Format("2006-01-02 15:04:05"))
}
