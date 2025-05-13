package handler

import (
	"ModEd/curriculum/controller"
	"ModEd/curriculum/utils"
	"fmt"
)

type deleteClassByIdHandler struct {
	classController controller.ClassControllerInterface
}

func NewDeleteClassByIdHandler(classController controller.ClassControllerInterface) *deleteClassByIdHandler {
	return &deleteClassByIdHandler{
		classController: classController,
	}
}

func (h *deleteClassByIdHandler) Execute() error {
	classes, err := h.classController.GetClasses()
	if err != nil {
		fmt.Println("Error getting classes:", err)
		return err
	}

	for _, class := range classes {
		class.Print()
	}

	classId := utils.GetUserInputUint("Enter the class ID to delete: ")

	confirm := utils.GetUserInput(fmt.Sprintf("Are you sure you want to delete class with Id %d? (y/n): ", classId))
	if confirmed, exists := confirmOptions[confirm]; !exists || !confirmed {
		fmt.Println("Deletion cancelled.")
		return nil
	}

	_, err = h.classController.DeleteClass(classId)
	if err != nil {
		fmt.Println("Error deleting class:", err)
		return err
	}

	fmt.Println("Class deleted successfully!")
	return nil
}
