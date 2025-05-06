package assessmentManager

import (
	"ModEd/project/controller"
	"ModEd/project/utils"
	"fmt"
	"strconv"
)

type ListCriteriaHandler struct {
	criteriaCtrl *controller.AssessmentCriteriaController
	linkCtrl     *controller.AssessmentCriteriaLinkController
}

func (h *ListCriteriaHandler) GetInputSequence() []utils.InputPrompt {
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
	}
}

func (h *ListCriteriaHandler) ExecuteMenuComponent(inputs []string) {
	projectID, _ := strconv.ParseUint(inputs[0], 10, 32)

	links, err := h.linkCtrl.ListProjectAssessmentCriteriaLinks(uint(projectID))
	if err != nil {
		fmt.Printf("Error retrieving links: %v\n", err)
		return
	}

	if len(links) == 0 {
		fmt.Println("No criteria linked to this assessment")
		return
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
}
