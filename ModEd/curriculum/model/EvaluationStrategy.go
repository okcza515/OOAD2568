// MEP-1008
package model

import (
	projectModel "ModEd/project/model"
	"bufio"
	"fmt"
	"os"
	"strings"
)

type EvaluationStrategy interface {
	Evaluate(criteria []projectModel.AssessmentCriteria) (float64, string)
}

type PresentationEvaluationStrategy struct{}
type AssignmentEvaluationStrategy struct{}
type ReportEvaluationStrategy struct{}

type MarkedCriteria struct {
	projectModel.AssessmentCriteria
	IsPass bool
}

type WeightedCriteria struct {
	projectModel.AssessmentCriteria
	Weight float64
}

func (p *PresentationEvaluationStrategy) Evaluate(criteria []projectModel.AssessmentCriteria) (float64, string) {
	fmt.Println("------------Evaluate Presentation------------")

	score := 0.0
	scale := map[string]float64{
		"poor":    5,
		"average": 10,
		"good":    15,
	}

	scanner := bufio.NewScanner(os.Stdin)

	for _, c := range criteria {
		for {
			fmt.Printf("Rate '%s' (poor/average/good): ", c.CriteriaName)
			scanner.Scan()
			input := strings.ToLower(strings.TrimSpace(scanner.Text()))

			if val, ok := scale[input]; ok {
				score += val
				break
			} else {
				fmt.Println("Invalid input. Please enter 'poor', 'average', or 'good'.")
			}
		}
	}

	fmt.Print("Enter comment: ")
	scanner.Scan()
	comment := scanner.Text()

	return score, comment
}

func (a *AssignmentEvaluationStrategy) Evaluate(criteria []projectModel.AssessmentCriteria) (float64, string) {
	fmt.Println("------------Evaluate Assignment------------")

	mockedCriteria := []MarkedCriteria{}
	scanner := bufio.NewScanner(os.Stdin)

	for _, c := range criteria {
		fmt.Printf("%s pass ? (yes/no): ", c.CriteriaName)
		scanner.Scan()
		input := strings.TrimSpace(scanner.Text())
		isPass := strings.ToLower(input) == "yes"
		mockedCriteria = append(mockedCriteria, MarkedCriteria{c, isPass})
	}

	score := 0.0
	for _, c := range mockedCriteria {
		if c.IsPass {
			score += 10
		}
	}

	fmt.Print("Enter comment: ")
	scanner.Scan()
	comment := scanner.Text()

	return score, comment
}

func (r *ReportEvaluationStrategy) Evaluate(criteria []projectModel.AssessmentCriteria) (float64, string) {

	fmt.Println("Evaluate Report")

	score := 0.0
	scale := map[string]float64{
		"poor":    5,
		"average": 10,
		"good":    15,
	}

	mockedCriteria := []MarkedCriteria{
		{criteria[0], true},
		{criteria[1], false},
		{criteria[2], true},
	}

	for _, c := range mockedCriteria {
		if c.IsPass {
			if val, ok := scale[c.CriteriaName]; ok {
				score += val
			} else {
				score += 10
			}
		}
	}
	return score, ""
}
