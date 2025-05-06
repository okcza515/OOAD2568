// MEP-1002
package handler

import (
	"ModEd/common/model"
	controller "ModEd/curriculum/controller"
	curriculumModel "ModEd/curriculum/model"
	"ModEd/curriculum/utils"
	"fmt"
	"strconv"
)

const (
	defaultCurriculumDataPath = "../data/curriculum/curriculum.json"
)

type curriculumHandler struct {
	curriculumController controller.CurriculumControllerInterface
}

func newCurriculumHandler(curriculumController controller.CurriculumControllerInterface) *curriculumHandler {
	return &curriculumHandler{
		curriculumController: curriculumController,
	}
}

func (h *curriculumHandler) createSeedCurriculum() (err error) {
	dataPath := utils.GetInputDataPath("curriculum", defaultCurriculumDataPath)
	_, err = h.curriculumController.CreateSeedCurriculum(dataPath)
	if err != nil {
		fmt.Println("Error creating seed curriculum:", err)
		return err
	}
	return nil
}
func (h *curriculumHandler) listCurriculums() (err error) {
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

func (h *curriculumHandler) getCurriculumById() (err error) {
	curriculumId := utils.GetUserInputUint("Enter the curriculum ID: ")
	curriculum, err := h.curriculumController.GetCurriculum(curriculumId)
	if err != nil {
		fmt.Println("Error getting curriculum:", err)
		return err
	}
	curriculum.Print()
	return nil
}

func (h *curriculumHandler) updateCurriculumById() (err error) {
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

	newStartYear := utils.GetUserInput(fmt.Sprintf("Start Year [%d]: ", curriculum.StartYear))
	if newStartYear != "" {
		startYear, err := strconv.Atoi(newStartYear)
		if err == nil {
			curriculum.StartYear = startYear
		} else {
			fmt.Println("Invalid start year format, keeping current value")
		}
	}

	newEndYear := utils.GetUserInput(fmt.Sprintf("End Year [%d]: ", curriculum.EndYear))
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

	// Validate updated curriculum
	if err := curriculum.Validate(); err != nil {
		fmt.Println("Validation error:", err)
		return err
	}

	confirm := utils.GetUserInput("Are you sure you want to update this curriculum? (y/n): ")
	if confirm != "y" {
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

func (h *curriculumHandler) deleteCurriculumById() (err error) {
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

func (h *curriculumHandler) createCurriculum() (err error) {
	fmt.Println("\nCreate New Curriculum:")

	name := utils.GetUserInput("Enter name: ")

	startYearStr := utils.GetUserInput("Enter start year: ")
	startYear, err := strconv.Atoi(startYearStr)
	if err != nil {
		fmt.Println("Invalid start year format")
		return err
	}

	endYearStr := utils.GetUserInput("Enter end year: ")
	endYear, err := strconv.Atoi(endYearStr)
	if err != nil {
		fmt.Println("Invalid end year format")
		return err
	}

	fmt.Println("Program Type Options:")
	for key, value := range model.ProgramTypeLabel {
		fmt.Printf("%d. %s\n", key, value)
	}

	programTypeStr := utils.GetUserInput("Select program type (enter number): ")
	programType, err := strconv.Atoi(programTypeStr)
	if err != nil {
		fmt.Println("Invalid program type")
		return err
	}

	departmentIdStr := utils.GetUserInput("Select department (enter number): ")
	departmentId, err := strconv.Atoi(departmentIdStr)
	if err != nil {
		fmt.Println("Invalid department ID")
		return err
	}

	curriculum := &curriculumModel.Curriculum{
		Name:         name,
		StartYear:    startYear,
		EndYear:      endYear,
		ProgramType:  model.ProgramType(programType),
		DepartmentId: uint(departmentId),
	}

	// Validate the curriculum
	if err := curriculum.Validate(); err != nil {
		fmt.Println("Validation error:", err)
		return err
	}

	fmt.Println("\nCurriculum to be created:")
	fmt.Printf("Name: %s\n", curriculum.Name)
	fmt.Printf("Start Year: %d\n", curriculum.StartYear)
	fmt.Printf("End Year: %d\n", curriculum.EndYear)
	fmt.Printf("Program Type: %s\n", model.ProgramTypeLabel[curriculum.ProgramType])

	confirm := utils.GetUserInput("\nConfirm creation? (y/n): ")
	if confirm != "y" {
		fmt.Println("Creation cancelled.")
		return nil
	}

	_, err = h.curriculumController.CreateCurriculum(curriculum)
	if err != nil {
		fmt.Println("Error creating curriculum:", err)
		return err
	}

	fmt.Println("\nCurriculum created successfully:")

	return nil
}
