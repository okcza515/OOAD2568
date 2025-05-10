package handler

import (
	"ModEd/curriculum/controller"
	"ModEd/curriculum/utils"
	"fmt"
)

const (
	defaultClassDataPath = "../data/curriculum/class.json"
)

type createSeedClassHandler struct {
	classController controller.ClassControllerInterface
}

func NewCreateSeedClassHandler(classController controller.ClassControllerInterface) *createSeedClassHandler {
	return &createSeedClassHandler{
		classController: classController,
	}
}

func (h *createSeedClassHandler) Execute() error {
	dataPath := utils.GetInputDataPath("class", defaultClassDataPath)
	_, err := h.classController.CreateSeedClass(dataPath)
	if err != nil {
		fmt.Println("Error creating seed class:", err)
		return err
	}
	fmt.Println("Seed class created successfully")
	return nil
}
