package controller

import (
	"ModEd/curriculum/model"
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

type AssignmentEvaluation struct{}
type PresentationEvaluation struct{}
type ReportEvaluation struct{}

type ProjectEvaluationController struct{}

func (e *ProjectEvaluationController) EvaluateTask(evaluation *model.Evaluation) (float64, string, error) {
	var evaluator model.ProjectEvaluationStrategy
	switch evaluation.AssignmentType {
	case "assignment":
		evaluator = &AssignmentEvaluation{}
	case "presentation":
		evaluator = &PresentationEvaluation{}
	case "report":
		evaluator = &ReportEvaluation{}
	default:
		return 0, "", errors.New("Invalid assignment type")
	}

	evaluation.SetEvaluationStrategy(evaluator)

	return evaluation.ExecuteEvaluation()
}

func (a *AssignmentEvaluation) Evaluate(evaluation model.Evaluation) (float64, string, error) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("-----Assignment Evaluation------")

	fmt.Print("Enter score: ")
	var score float64
	_, err := fmt.Scanf("%f", &score)
	if err != nil {
		return 0, "", err
	}

	reader.ReadString('\n')

	fmt.Print("Enter comment: ")
	comment, err := reader.ReadString('\n')
	if err != nil {
		return 0, "", err
	}
	comment = strings.TrimSpace(comment)

	return score, comment, nil
}

func (a *PresentationEvaluation) Evaluate(evaluation model.Evaluation) (float64, string, error) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("-----Presentation Evaluation-----")

	fmt.Print("Enter score: ")
	var score float64
	_, err := fmt.Scanf("%f", &score)
	if err != nil {
		return 0, "", err
	}

	reader.ReadString('\n')

	fmt.Print("Enter comment: ")
	comment, err := reader.ReadString('\n')
	if err != nil {
		return 0, "", err
	}
	comment = strings.TrimSpace(comment)

	return score, comment, nil
}

func (a *ReportEvaluation) Evaluate(evaluation model.Evaluation) (float64, string, error) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("------Report Evaluation------")

	fmt.Print("Enter score: ")
	var score float64
	_, err := fmt.Scanf("%f", &score)
	if err != nil {
		return 0, "", err
	}

	reader.ReadString('\n')

	fmt.Print("Enter comment: ")
	comment, err := reader.ReadString('\n')
	if err != nil {
		return 0, "", err
	}
	comment = strings.TrimSpace(comment)

	return score, comment, nil
}
