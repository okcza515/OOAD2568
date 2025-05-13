package assessmentManager

import (
	"fmt"
	"strconv"
)

type UpdateLinkHandler struct {
	*BaseAssessmentStrategy
}

func (h *UpdateLinkHandler) Execute() error {
	inputs := make([]string, 3)

	fmt.Print("Senior Project ID (number, -1 to cancel): ")
	fmt.Scanln(&inputs[0])
	if inputs[0] == "-1" {
		fmt.Println("Operation cancelled")
		return nil
	}

	fmt.Print("Current Criteria ID (number): ")
	fmt.Scanln(&inputs[1])
	if inputs[1] == "-1" {
		fmt.Println("Operation cancelled")
		return nil
	}

	fmt.Print("New Criteria ID (number): ")
	fmt.Scanln(&inputs[2])
	if inputs[2] == "-1" {
		fmt.Println("Operation cancelled")
		return nil
	}

	projectID, _ := strconv.ParseUint(inputs[0], 10, 32)
	oldID, _ := strconv.ParseUint(inputs[1], 10, 32)
	newID, _ := strconv.ParseUint(inputs[2], 10, 32)

	assessment, err := h.assessmentCtrl.RetrieveByID(uint(projectID))
	if err != nil || assessment == nil {
		fmt.Println("Assessment not found")
		return nil
	}

	existing, err := h.linkCtrl.RetrieveAssessmentCriteriaLink(assessment.ID, uint(oldID))
	if err != nil || existing == nil {
		fmt.Println("Existing link not found")
		return nil
	}

	existing.AssessmentCriteriaId = uint(newID)
	err = h.linkCtrl.UpdateByID(existing)
	if err != nil {
		return fmt.Errorf("error updating link: %w", err)
	}

	fmt.Println("Successfully updated criteria link")
	return nil
}
