package assessmentManager

import (
	"ModEd/project/controller"
	"ModEd/project/utils"
	"fmt"
	"strconv"
)

type DeleteLinkHandler struct {
	assessmentCtrl *controller.AssessmentController
	linkCtrl       *controller.AssessmentCriteriaLinkController
	criteriaCtrl   *controller.AssessmentCriteriaController
}

func (h *DeleteLinkHandler) GetInputSequence() []utils.InputPrompt {
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
			Label: "Criteria ID to unlink (number)",
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

func (h *DeleteLinkHandler) ExecuteMenuComponent(inputs []string) {
	projectID, _ := strconv.ParseUint(inputs[0], 10, 32)
	criteriaID, _ := strconv.ParseUint(inputs[1], 10, 32)

	// Get assessment
	assessment, err := h.assessmentCtrl.RetrieveAssessmentBySeniorProjectId(uint(projectID))
	if err != nil || assessment == nil {
		fmt.Println("The Senior Project's Assessment not found")
		return
	}

	// Delete link
	err = h.linkCtrl.DeleteAssessmentCriteriaLink(assessment.ID, uint(criteriaID))
	if err != nil {
		fmt.Printf("Error deleting link: %v\n", err)
		return
	}

	fmt.Println("Successfully unlinked criteria from assessment")
}
