package criteriaManager

import "fmt"

type DefineCriteriaHandler struct {
	*BaseCriteriaStrategy
}

func (h *DefineCriteriaHandler) Execute() error {
	name, err := h.getNameInput("Criteria Name (-1 to cancel): ")
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	err = h.controller.InsertAssessmentCriteria(name)
	if err != nil {
		return fmt.Errorf("error adding criteria: %w", err)
	}

	fmt.Println("Criteria added successfully!")
	return nil
}
