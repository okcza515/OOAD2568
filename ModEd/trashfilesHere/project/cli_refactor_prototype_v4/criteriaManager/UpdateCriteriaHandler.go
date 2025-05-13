package criteriaManager

import (
	"fmt"
)

type UpdateCriteriaHandler struct {
	*BaseCriteriaStrategy
}

func (h *UpdateCriteriaHandler) Execute() error {
	id, err := h.getIDInput("Criteria ID to update (-1 to cancel): ")
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	name, err := h.getNameInput("New Criteria Name: ")
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	criteria, err := h.controller.RetrieveByID(id)
	if err != nil || criteria == nil {
		return fmt.Errorf("criteria not found")
	}

	criteria.CriteriaName = name
	err = h.controller.UpdateByID(criteria)
	if err != nil {
		return fmt.Errorf("error updating criteria: %w", err)
	}

	fmt.Println("Criteria updated successfully!")
	return nil
}
