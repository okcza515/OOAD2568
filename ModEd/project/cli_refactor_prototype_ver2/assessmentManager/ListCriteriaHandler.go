package assessmentManager

import (
	"fmt"
	"strconv"
)

type ListCriteriaHandler struct {
	*BaseAssessmentStrategy
}

func (h *ListCriteriaHandler) Execute() error {
	var projectIDStr string
	fmt.Print("Senior Project ID (number, -1 to cancel): ")
	fmt.Scanln(&projectIDStr)
	if projectIDStr == "-1" {
		fmt.Println("Operation cancelled")
		return nil
	}

	projectID, _ := strconv.ParseUint(projectIDStr, 10, 32)
	links, err := h.linkCtrl.ListProjectAssessmentCriteriaLinks(uint(projectID))
	if err != nil {
		return fmt.Errorf("error retrieving links: %w", err)
	}

	if len(links) == 0 {
		fmt.Println("No criteria linked to this assessment")
		return nil
	}

	fmt.Println("\nLinked Criteria:")
	for _, link := range links {
		criteria, err := h.criteriaCtrl.RetrieveByID(link.AssessmentCriteriaId)
		if err != nil || criteria == nil {
			fmt.Printf("  - Criteria ID %d: Error retrieving details\n", link.AssessmentCriteriaId)
			continue
		}
		fmt.Printf("  - %s (ID: %d)\n", criteria.CriteriaName, criteria.ID)
	}
	return nil
}
