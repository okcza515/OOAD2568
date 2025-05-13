package handler

import (
	"ModEd/common/model"
	"ModEd/curriculum/controller"
	"ModEd/curriculum/utils"
	"fmt"
	"strconv"
)

type updateCurriculumByIdHandler struct {
	curriculumController controller.CurriculumControllerInterface
}

func NewUpdateCurriculumByIdHandler(curriculumController controller.CurriculumControllerInterface) *updateCurriculumByIdHandler {
	return &updateCurriculumByIdHandler{
		curriculumController: curriculumController,
	}
}

func (h *updateCurriculumByIdHandler) Execute() error {
	curriculumId := utils.GetUserInputUint("Enter the curriculum ID: ")
	curriculum, err := h.curriculumController.GetCurriculum(curriculumId)
	if err != nil {
		fmt.Println("Error getting curriculum:", err)
		return err
	}

	fmt.Println("\nCurrent curriculum information:")
	curriculum.Print()

	fmt.Println("\nEnter new values (leave blank to keep current value):")
	newName := utils.GetUserInput(fmt.Sprintf("Name [%s]: ", curriculum.Name))
	if newName != "" {
		curriculum.Name = newName
	}

	newStartYear := utils.GetUserInput(fmt.Sprintf("Start Year [%d] format (YYYY): ", curriculum.StartYear))
	if newStartYear != "" {
		startYear, err := strconv.Atoi(newStartYear)
		if err == nil {
			curriculum.StartYear = startYear
		} else {
			fmt.Println("Invalid start year format, keeping current value")
		}
	}

	newEndYear := utils.GetUserInput(fmt.Sprintf("End Year [%d] format (YYYY): ", curriculum.EndYear))
	if newEndYear != "" {
		endYear, err := strconv.Atoi(newEndYear)
		if err == nil {
			curriculum.EndYear = endYear
		} else {
			fmt.Println("Invalid end year format, keeping current value")
		}
	}

	fmt.Println("Program Type Choice:")
	for key, value := range model.ProgramTypeLabel {
		fmt.Printf("%d. %s\n", key, value)
	}
	newProgramType := utils.GetUserInput(fmt.Sprintf("Program Type [%s]: ", curriculum.ProgramType))
	if newProgramType != "" {
		programType, err := strconv.Atoi(newProgramType)
		if err == nil {
			curriculum.ProgramType = model.ProgramType(programType)
		} else {
			fmt.Println("Invalid program type format, keeping current value")
		}
	}

	if err := curriculum.Validate(); err != nil {
		fmt.Println("Validation error:", err)
		return err
	}

	confirm := utils.GetUserInput("Are you sure you want to update this curriculum? (y/n): ")
	if confirmed, exists := confirmOptions[confirm]; !exists || !confirmed {
		fmt.Println("Update cancelled.")
		return nil
	}

	updatedCurriculum, err := h.curriculumController.UpdateCurriculum(curriculum)
	if err != nil {
		fmt.Println("Error updating curriculum:", err)
		return err
	}

	fmt.Println("Curriculum updated successfully:")
	updatedCurriculum.Print()

	return nil
}
