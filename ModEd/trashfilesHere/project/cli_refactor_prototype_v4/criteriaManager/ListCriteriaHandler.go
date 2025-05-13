package criteriaManager

import (
	"fmt"
)

type ListCriteriaHandler struct {
	*BaseCriteriaStrategy
}

func (h *ListCriteriaHandler) Execute() error {
	criteriaList, err := h.controller.List(map[string]interface{}{})
	if err != nil {
		return fmt.Errorf("error listing criteria: %w", err)
	}

	if len(criteriaList) == 0 {
		fmt.Println("No criteria found")
		return nil
	}

	fmt.Println("\nAssessment Criteria:")
	for _, c := range criteriaList {
		fmt.Printf("ID: %d, Name: %s\n", c.ID, c.CriteriaName)
	}
	return nil
}
