package criteriaManager

import (
	"ModEd/project/controller"
	"ModEd/project/utils"
	"fmt"
)

type DefineCriteriaHandler struct {
	controller *controller.AssessmentCriteriaController
}

func (h *DefineCriteriaHandler) GetInputSequence() []utils.InputPrompt {
	return []utils.InputPrompt{
		{
			Label: "Criteria Name (-1 to cancel)",
			Validate: func(input string) error {
				if input == "-1" {
					return fmt.Errorf("operation cancelled")
				}
				if len(input) < 3 {
					return fmt.Errorf("name must be at least 3 characters")
				}
				return nil
			},
		},
	}
}

func (h *DefineCriteriaHandler) ExecuteMenuComponent(inputs []string) {
	err := h.controller.InsertAssessmentCriteria(inputs[0])
	if err != nil {
		fmt.Printf("Error adding criteria: %v\n", err)
	} else {
		fmt.Println("Criteria added successfully!")
	}
}
