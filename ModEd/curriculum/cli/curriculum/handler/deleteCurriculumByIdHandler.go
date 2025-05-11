package handler

import (
	"ModEd/curriculum/controller"
	"ModEd/curriculum/utils"
	"fmt"
)

type deleteCurriculumByIdHandler struct {
	curriculumController controller.CurriculumControllerInterface
}

func NewDeleteCurriculumByIdHandler(curriculumController controller.CurriculumControllerInterface) *deleteCurriculumByIdHandler {
	return &deleteCurriculumByIdHandler{
		curriculumController: curriculumController,
	}
}

func (h *deleteCurriculumByIdHandler) Execute() error {
	curriculums, err := h.curriculumController.GetCurriculums()
	if err != nil {
		fmt.Println("Error getting curriculums:", err)
		return err
	}

	for _, curriculum := range curriculums {
		curriculum.Print()
	}

	curriculumId := utils.GetUserInputUint("Enter the curriculum Id to delete: ")

	confirm := utils.GetUserInput(fmt.Sprintf("Are you sure you want to delete curriculum with Id %d? (y/n): ", curriculumId))
	if confirm != "y" {
		fmt.Println("Deletion cancelled.")
		return nil
	}

	_, err = h.curriculumController.DeleteCurriculum(curriculumId)
	if err != nil {
		fmt.Println("Error deleting curriculum:", err)
		return err
	}

	fmt.Println("Curriculum deleted successfully!")
	return nil
}
