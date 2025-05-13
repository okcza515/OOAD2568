package criteriaManager

import (
	"ModEd/project/controller"
	"ModEd/project/utils"
	"fmt"
	"strconv"
)

type DeleteCriteriaHandler struct {
	controller *controller.AssessmentCriteriaController
}

func (h *DeleteCriteriaHandler) GetInputSequence() []utils.InputPrompt {
	return []utils.InputPrompt{
		{
			Label: "Criteria ID to delete (-1 to cancel)",
			Validate: func(input string) error {
				if input == "-1" {
					return fmt.Errorf("operation cancelled")
				}
				_, err := strconv.ParseUint(input, 10, 32)
				return err
			},
		},
	}
}

func (h *DeleteCriteriaHandler) ExecuteMenuComponent(inputs []string) {
	id, _ := strconv.ParseUint(inputs[0], 10, 32)

	err := h.controller.DeleteAssessmentCriteria(uint(id))
	if err != nil {
		fmt.Printf("Error deleting criteria: %v\n", err)
	} else {
		fmt.Println("Criteria deleted successfully!")
	}
}
