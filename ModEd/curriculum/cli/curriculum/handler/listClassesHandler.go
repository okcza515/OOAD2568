package handler

import (
	"ModEd/curriculum/controller"
	"fmt"
)

type listClassesHandler struct {
	classController controller.ClassControllerInterface
}

func NewListClassesHandler(classController controller.ClassControllerInterface) *listClassesHandler {
	return &listClassesHandler{
		classController: classController,
	}
}

func (h *listClassesHandler) Execute() error {
	classes, err := h.classController.GetClasses()
	if err != nil {
		fmt.Println("Error getting classes:", err)
		return err
	}

	for _, class := range classes {
		class.Print()
	}
	return nil
}
