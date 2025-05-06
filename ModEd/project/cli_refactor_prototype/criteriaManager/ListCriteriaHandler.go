package criteriaManager

import (
	"ModEd/project/controller"
	"ModEd/project/utils"
	"fmt"
)

type ListCriteriaHandler struct {
	controller *controller.AssessmentCriteriaController
}

func (h *ListCriteriaHandler) GetInputSequence() []utils.InputPrompt {
	return nil // No input needed for listing
}

func (h *ListCriteriaHandler) ExecuteMenuComponent(_ []string) {
	criteriaList, err := h.controller.ListAllAssessmentCriterias()
	if err != nil {
		fmt.Printf("Error listing criteria: %v\n", err)
		return
	}

	if len(criteriaList) == 0 {
		fmt.Println("No criteria found")
		return
	}

	fmt.Println("\nAssessment Criteria:")
	for _, c := range criteriaList {
		fmt.Printf("ID: %d, Name: %s\n", c.ID, c.CriteriaName)
	}
}
