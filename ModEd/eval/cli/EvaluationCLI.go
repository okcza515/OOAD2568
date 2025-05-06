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
	Evaluate(studentCode, instructorCode string, assessmentID uint, assessmentType string, score uint, comment string)
}

type AssessmentEvaluation struct {
	Controller *controller.EvaluationController
}

func (a *AssessmentEvaluation) Evaluate(studentCode, instructorCode string, assessmentID uint, assessmentType string, score uint, comment string) {
	a.Controller.EvaluateAssessment(studentCode, instructorCode, assessmentID, assessmentType, score, comment)
}

type EvaluationContext struct {
	Strategy EvaluationStrategy
}

func (c *EvaluationContext) SetStrategy(s EvaluationStrategy) {
	c.Strategy = s
}

func (c *EvaluationContext) Evaluate(studentCode, instructorCode string, assessmentID uint, assessmentType string, score uint, comment string) {
	c.Strategy.Evaluate(studentCode, instructorCode, assessmentID, assessmentType, score, comment)
}

func evaluateGeneral(ec *controller.EvaluationController, assessmentType string) {
	var studentCode, instructorCode string
	var assessmentID, score uint

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

	fmt.Printf("%sID: ", strings.Title(assessmentType))
	fmt.Scanln(&assessmentID)
	if assessmentID == 0 {
		fmt.Printf("%s ID cannot be zero.\n", strings.Title(assessmentType))
		return
	}

	fmt.Print("Score: ")
	fmt.Scanln(&score)

	fmt.Print("Comment: ")
	comment, _ := reader.ReadString('\n')
	comment = strings.TrimSpace(comment)

	ctx := &EvaluationContext{}
	ctx.SetStrategy(&AssessmentEvaluation{Controller: ec})
	ctx.Evaluate(studentCode, instructorCode, assessmentID, assessmentType, score, comment)
	fmt.Printf("%s evaluated successfully.\n", strings.Title(assessmentType))
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
		fmt.Printf("%s Evaluation || StudentID: %s | AssessmentID: %d | Score: %d | Comment: %s | Evaluated At: %s\n",
			strings.Title(e.AssessmentType), e.StudentCode, e.AssessmentID, e.Score, e.Comment, e.EvaluatedAt.Format(time.RFC1123))
	}
}
