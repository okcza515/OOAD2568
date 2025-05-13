package assessmentManager

import (
	"ModEd/project/controller"
	"ModEd/project/utils"
	"fmt"
	"strconv"
)

type UpdateLinkHandler struct {
	criteriaCtrl   *controller.AssessmentCriteriaController
	assessmentCtrl *controller.AssessmentController
	linkCtrl       *controller.AssessmentCriteriaLinkController
}

func (h *UpdateLinkHandler) GetInputSequence() []utils.InputPrompt {
	return []utils.InputPrompt{
		{
			Label: "Senior Project ID (number, -1 to cancel)",
			Validate: func(input string) error {
				if input == "-1" {
					return fmt.Errorf("operation cancelled")
				}
				_, err := strconv.ParseUint(input, 10, 32)
				return err
			},
		},
		{
			Label: "Current Criteria ID (number)",
			Validate: func(input string) error {
				if input == "-1" {
					return fmt.Errorf("operation cancelled")
				}
				_, err := strconv.ParseUint(input, 10, 32)
				return err
			},
		},
		{
			Label: "New Criteria ID (number)",
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

func (h *UpdateLinkHandler) ExecuteMenuComponent(inputs []string) {
	projectID, _ := strconv.ParseUint(inputs[0], 10, 32)
	oldID, _ := strconv.ParseUint(inputs[1], 10, 32)
	newID, _ := strconv.ParseUint(inputs[2], 10, 32)

	// Get assessment
	assessment, err := h.assessmentCtrl.RetrieveByID(uint(projectID))
	if err != nil || assessment == nil {
		fmt.Println("Assessment not found")
		return
	}

	// Verify existing link
	existing, err := h.linkCtrl.RetrieveAssessmentCriteriaLink(assessment.ID, uint(oldID))
	if err != nil || existing == nil {
		fmt.Println("Existing link not found")
		return
	}

	// Update criteria ID
	existing.AssessmentCriteriaId = uint(newID)
	err = h.linkCtrl.UpdateByID(existing)
	if err != nil {
		fmt.Printf("Error updating link: %v\n", err)
		return
	}

	fmt.Println("Successfully updated criteria link")
}
