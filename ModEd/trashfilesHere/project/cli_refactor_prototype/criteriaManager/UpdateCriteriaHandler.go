package criteriaManager

import (
	"ModEd/project/controller"
	"ModEd/project/utils"
	"fmt"
	"strconv"
)

type UpdateCriteriaHandler struct {
	controller *controller.AssessmentCriteriaController
}

func (h *UpdateCriteriaHandler) GetInputSequence() []utils.InputPrompt {
	return []utils.InputPrompt{
		{
			Label: "Criteria ID to update (-1 to cancel)",
			Validate: func(input string) error {
				if input == "-1" {
					return fmt.Errorf("operation cancelled")
				}
				_, err := strconv.ParseUint(input, 10, 32)
				return err
			},
		},
		{
			Label: "New Criteria Name",
			Validate: func(input string) error {
				if len(input) < 3 {
					return fmt.Errorf("name must be at least 3 characters")
				}
				return nil
			},
		},
	}
}

func (h *UpdateCriteriaHandler) ExecuteMenuComponent(inputs []string) {
	id, _ := strconv.ParseUint(inputs[0], 10, 32)
	newName := inputs[1]

	criteria, err := h.controller.RetrieveByID(uint(id))
	if err != nil || criteria == nil {
		fmt.Println("Criteria not found")
		return
	}

	criteria.CriteriaName = newName
	err = h.controller.UpdateByID(criteria)
	if err != nil {
		fmt.Printf("Error updating criteria: %v\n", err)
	} else {
		fmt.Println("Criteria updated successfully!")
	}
}
