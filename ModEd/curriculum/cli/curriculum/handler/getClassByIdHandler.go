package handler

import (
	"ModEd/curriculum/controller"
	"ModEd/curriculum/utils"
	"fmt"
)

type getClassByIdHandler struct {
	classController controller.ClassControllerInterface
}

func NewGetClassByIdHandler(classController controller.ClassControllerInterface) *getClassByIdHandler {
	return &getClassByIdHandler{
		classController: classController,
	}
}

func (h *getClassByIdHandler) Execute() error {
	classId := utils.GetUserInputUint("Enter the class ID: ")
	class, err := h.classController.GetClass(classId)
	if err != nil {
		fmt.Println("Error getting class:", err)
		return err
	}
	class.Print()
	return nil
}
