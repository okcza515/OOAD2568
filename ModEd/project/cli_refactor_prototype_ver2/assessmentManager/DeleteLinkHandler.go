package assessmentManager

import (
	"fmt"
	"strconv"
)

type DeleteLinkHandler struct {
	*BaseAssessmentStrategy
}

func (h *DeleteLinkHandler) Execute() error {
	inputs := make([]string, 2)

	fmt.Print("Senior Project ID (number, -1 to cancel): ")
	fmt.Scanln(&inputs[0])
	if inputs[0] == "-1" {
		fmt.Println("Operation cancelled")
		return nil
	}

	fmt.Print("Criteria ID to unlink (number): ")
	fmt.Scanln(&inputs[1])
	if inputs[1] == "-1" {
		fmt.Println("Operation cancelled")
		return nil
	}

	projectID, _ := strconv.ParseUint(inputs[0], 10, 32)
	criteriaID, _ := strconv.ParseUint(inputs[1], 10, 32)

	assessment, err := h.assessmentCtrl.RetrieveAssessmentBySeniorProjectId(uint(projectID))
	if err != nil || assessment == nil {
		fmt.Println("The Senior Project's Assessment not found")
		return nil
	}

	err = h.linkCtrl.DeleteAssessmentCriteriaLink(assessment.ID, uint(criteriaID))
	if err != nil {
		return fmt.Errorf("error deleting link: %w", err)
	}

	fmt.Println("Successfully unlinked criteria from assessment")
	return nil
}
