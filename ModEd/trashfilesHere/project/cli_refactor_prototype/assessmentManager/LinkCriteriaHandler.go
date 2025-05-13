package assessmentManager

import (
	"ModEd/project/controller"
	"ModEd/project/utils"
	"fmt"
	"strconv"
)

type LinkCriteriaHandler struct {
	criteriaCtrl   *controller.AssessmentCriteriaController
	assessmentCtrl *controller.AssessmentController
	linkCtrl       *controller.AssessmentCriteriaLinkController
}

func (h *LinkCriteriaHandler) GetInputSequence() []utils.InputPrompt {
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
			Label: "Criteria ID to link (number, -1 to cancel)",
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

func (h *LinkCriteriaHandler) ExecuteMenuComponent(inputs []string) {
	projectID, _ := strconv.ParseUint(inputs[0], 10, 32)
	criteriaID, _ := strconv.ParseUint(inputs[1], 10, 32)

	// Check existing links
	existing, err := h.linkCtrl.ListProjectAssessmentCriteriaLinks(uint(projectID))
	if err != nil {
		fmt.Printf("Error checking existing links: %v\n", err)
		return
	}

	for _, link := range existing {
		if link.AssessmentCriteriaId == uint(criteriaID) {
			fmt.Println("This criteria is already linked")
			return
		}
	}

	// Get assessment ID
	assessment, err := h.assessmentCtrl.RetrieveAssessmentBySeniorProjectId(uint(projectID))
	if err != nil {
		fmt.Printf("Error retrieving assessment: %v\n", err)
		return
	}

	// Create new link
	_, err = h.linkCtrl.InsertAssessmentCriteriaLink(assessment.ID, uint(criteriaID))
	if err != nil {
		fmt.Printf("Error creating link: %v\n", err)
		return
	}

	fmt.Println("Successfully linked criteria to assessment")
}
