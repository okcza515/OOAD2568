package handler

import (
	"ModEd/curriculum/controller"
	"ModEd/curriculum/utils"
	"fmt"
)

type getCurriculumByIdHandler struct {
	curriculumController controller.CurriculumControllerInterface
}

func NewGetCurriculumByIdHandler(curriculumController controller.CurriculumControllerInterface) *getCurriculumByIdHandler {
	return &getCurriculumByIdHandler{
		curriculumController: curriculumController,
	}
}

func (h *getCurriculumByIdHandler) Execute() error {
	curriculumId := utils.GetUserInputUint("Enter the curriculum ID: ")
	curriculum, err := h.curriculumController.GetCurriculum(curriculumId)
	if err != nil {
		fmt.Println("Error getting curriculum:", err)
		return err
	}
	curriculum.Print()
	return nil
}
