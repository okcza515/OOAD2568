package assessmentManager

import (
	"fmt"
	"strconv"
)

type LinkCriteriaHandler struct {
	*BaseAssessmentStrategy
}

func (h *LinkCriteriaHandler) Execute() error {
	inputs := make([]string, 2)

	fmt.Print("Senior Project ID (number, -1 to cancel): ")
	fmt.Scanln(&inputs[0])
	if inputs[0] == "-1" {
		fmt.Println("Operation cancelled")
		return nil
	}

	fmt.Print("Criteria ID to link (number, -1 to cancel): ")
	fmt.Scanln(&inputs[1])
	if inputs[1] == "-1" {
		fmt.Println("Operation cancelled")
		return nil
	}

	projectID, _ := strconv.ParseUint(inputs[0], 10, 32)
	criteriaID, _ := strconv.ParseUint(inputs[1], 10, 32)

	existing, err := h.linkCtrl.ListProjectAssessmentCriteriaLinks(uint(projectID))
	if err != nil {
		return fmt.Errorf("error checking existing links: %w", err)
	}

	for _, link := range existing {
		if link.AssessmentCriteriaId == uint(criteriaID) {
			fmt.Println("This criteria is already linked")
			return nil
		}
	}

	assessment, err := h.assessmentCtrl.RetrieveAssessmentBySeniorProjectId(uint(projectID))
	if err != nil {
		return fmt.Errorf("error retrieving assessment: %w", err)
	}

	_, err = h.linkCtrl.InsertAssessmentCriteriaLink(assessment.ID, uint(criteriaID))
	if err != nil {
		return fmt.Errorf("error creating link: %w", err)
	}

	fmt.Println("Successfully linked criteria to assessment")
	return nil
}
