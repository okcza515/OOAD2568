//Chanawat Limpanatewin 65070503445
//MEP-1006

package cli

import (
	"ModEd/eval/controller"
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

type EvaluationStrategy interface {
	Evaluate(studentCode, instructorCode string, id uint, score uint, comment string)
}

type AssignmentEvaluation struct {
	Controller *controller.EvaluationController
}

func (a *AssignmentEvaluation) Evaluate(studentCode, instructorCode string, id uint, score uint, comment string) {
	a.Controller.EvaluateAssignment(studentCode, instructorCode, id, score, comment)
}

type QuizEvaluation struct {
	Controller *controller.EvaluationController
}

func (q *QuizEvaluation) Evaluate(studentCode, instructorCode string, id uint, score uint, comment string) {
	q.Controller.EvaluateQuiz(studentCode, instructorCode, id, score, comment)
}

type EvaluationContext struct {
	Strategy EvaluationStrategy
}

func (c *EvaluationContext) SetStrategy(s EvaluationStrategy) {
	c.Strategy = s
}

func (c *EvaluationContext) Evaluate(studentCode, instructorCode string, id uint, score uint, comment string) {
	c.Strategy.Evaluate(studentCode, instructorCode, id, score, comment)
}

func evaluateGeneral(ec *controller.EvaluationController, evalType string) {
	var studentCode, instructorCode string
	var id, score uint

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("StudentID: ")
	fmt.Scanln(&studentCode)
	studentCode = strings.TrimSpace(studentCode)
	if studentCode == "" {
		fmt.Println("StudentID cannot be blank.")
		return
	}

	fmt.Print("InstructorID: ")
	fmt.Scanln(&instructorCode)
	instructorCode = strings.TrimSpace(instructorCode)
	if instructorCode == "" {
		fmt.Println("InstructorID cannot be blank.")
		return
	}

	if evalType == "assignment" {
		fmt.Print("AssignmentID: ")
	} else if evalType == "quiz" {
		fmt.Print("QuizID: ")
	}
	fmt.Scanln(&id)
	if id == 0 {
		fmt.Printf("%s ID cannot be zero.\n", strings.Title(evalType))
		return
	}

	fmt.Print("Score: ")
	fmt.Scanln(&score)

	fmt.Print("Comment: ")
	comment, _ := reader.ReadString('\n')
	comment = strings.TrimSpace(comment)

	ctx := &EvaluationContext{}
	if evalType == "assignment" {
		ctx.SetStrategy(&AssignmentEvaluation{Controller: ec})
	} else if evalType == "quiz" {
		ctx.SetStrategy(&QuizEvaluation{Controller: ec})
	}
	ctx.Evaluate(studentCode, instructorCode, id, score, comment)
	fmt.Printf("%s evaluated successfully.\n", strings.Title(evalType))
}

func RunEvaluationCLI(ec *controller.EvaluationController) {
	for {
		fmt.Println("\nEvaluation Menu")
		fmt.Println("1. Evaluate Assignment")
		fmt.Println("2. Evaluate Quiz")
		fmt.Println("3. Display Evaluations")
		fmt.Println("0. Exit")

		var choice string
		fmt.Print("Enter choice: ")
		fmt.Scanln(&choice)

		switch choice {
		case "1":
			evaluateGeneral(ec, "assignment")
		case "2":
			evaluateGeneral(ec, "quiz")
		case "3":
			displayEvaluations(ec)
		case "0":
			return
		default:
			fmt.Println("Invalid choice")
		}
	}
}

func displayEvaluations(ec *controller.EvaluationController) {
	for _, e := range ec.ListEvaluations() {
		if e.AssignmentID != nil {
			fmt.Printf("Assignment Evaluation || StudentID: %s | AssignmentID: %d | Score: %d | Comment: %s | Evaluated At: %s\n",
				e.StudentCode, *e.AssignmentID, e.Score, e.Comment, e.EvaluatedAt.Format(time.RFC1123))
		}
		if e.QuizID != nil {
			fmt.Printf("Quiz Evaluation || StudentID: %s | QuizID: %d | Score: %d | Comment: %s | Evaluated At: %s\n",
				e.StudentCode, *e.QuizID, e.Score, e.Comment, e.EvaluatedAt.Format(time.RFC1123))
		}
	}
}
