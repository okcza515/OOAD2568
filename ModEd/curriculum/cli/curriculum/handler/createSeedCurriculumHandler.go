package handler

import (
	"ModEd/curriculum/controller"
	"ModEd/curriculum/utils"
	"fmt"
)

const (
	defaultCurriculumDataPath = "../data/curriculum/curriculum.json"
)

type createSeedCurriculumHandler struct {
	curriculumController controller.CurriculumControllerInterface
}

func NewCreateSeedCurriculumHandler(curriculumController controller.CurriculumControllerInterface) *createSeedCurriculumHandler {
	return &createSeedCurriculumHandler{
		curriculumController: curriculumController,
	}
}

func (h *createSeedCurriculumHandler) Execute() error {
	dataPath := utils.GetInputDataPath("curriculum", defaultCurriculumDataPath)
	_, err := h.curriculumController.CreateSeedCurriculum(dataPath)
	if err != nil {
		fmt.Println("Error creating seed curriculum:", err)
		return err
	}
	fmt.Println("Seed curriculum created successfully")
	return nil
}
