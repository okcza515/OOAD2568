package handler

import (
	"ModEd/curriculum/controller"
	"fmt"
)

type listCurriculumsHandler struct {
	curriculumController controller.CurriculumControllerInterface
}

func NewListCurriculumsHandler(curriculumController controller.CurriculumControllerInterface) *listCurriculumsHandler {
	return &listCurriculumsHandler{
		curriculumController: curriculumController,
	}
}

func (h *listCurriculumsHandler) Execute() error {
	curriculums, err := h.curriculumController.GetCurriculums()
	if err != nil {
		fmt.Println("Error getting curriculums:", err)
		return err
	}

	for _, curriculum := range curriculums {
		curriculum.Print()
	}
	return nil
}
