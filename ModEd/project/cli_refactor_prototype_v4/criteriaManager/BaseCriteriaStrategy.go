package criteriaManager

import (
	"ModEd/project/controller"
	"fmt"
	"strconv"
)

type BaseCriteriaStrategy struct {
	controller *controller.AssessmentCriteriaController
}

func NewBaseCriteriaStrategy(ctrl *controller.AssessmentCriteriaController) *BaseCriteriaStrategy {
	return &BaseCriteriaStrategy{controller: ctrl}
}

func (b *BaseCriteriaStrategy) getIDInput(prompt string) (uint, error) {
	fmt.Print(prompt)
	var input string
	fmt.Scanln(&input)

	if input == "-1" {
		return 0, fmt.Errorf("operation cancelled")
	}

	id, err := strconv.ParseUint(input, 10, 32)
	if err != nil {
		return 0, fmt.Errorf("invalid ID format")
	}
	return uint(id), nil
}

func (b *BaseCriteriaStrategy) getNameInput(prompt string) (string, error) {
	fmt.Print(prompt)
	var input string
	fmt.Scanln(&input)

	if input == "-1" {
		return "", fmt.Errorf("operation cancelled")
	}

	if len(input) < 3 {
		return "", fmt.Errorf("name must be at least 3 characters")
	}
	return input, nil
}
