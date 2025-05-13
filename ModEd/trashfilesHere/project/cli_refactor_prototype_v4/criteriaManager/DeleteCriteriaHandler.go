package criteriaManager

import "fmt"

type DeleteCriteriaHandler struct {
	*BaseCriteriaStrategy
}

func (h *DeleteCriteriaHandler) Execute() error {
	id, err := h.getIDInput("Criteria ID to delete (-1 to cancel): ")
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	err = h.controller.DeleteAssessmentCriteria(id)
	if err != nil {
		return fmt.Errorf("error deleting criteria: %w", err)
	}

	fmt.Println("Criteria deleted successfully!")
	return nil
}
