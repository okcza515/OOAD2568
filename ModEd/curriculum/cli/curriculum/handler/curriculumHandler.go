// MEP-1002
package handler

import (
	"ModEd/common/model"
	controller "ModEd/curriculum/controller"
	"ModEd/curriculum/utils"
	"errors"
	"fmt"
	"strconv"
)

const (
	defaultCurriculumDataPath = "../../data/curriculum/curriculum.json"
)

func RunCurriculumCLIHandler(curriculumController controller.CurriculumControllerInterface) {
	handler := newCurriculumHandler(curriculumController)

	menuManager := NewMenuManager(map[string]func() error{
		"1": handler.createSeedCurriculum,
		"2": handler.listCurriculums,
		"3": handler.getCurriculumById,
		"4": handler.updateCurriculumById,
		"5": handler.deleteCurriculumById,
		"0": func() error {
			fmt.Println("Exiting...")
			return ExitCommand
		},
	})

	for {

		choice := menuManager.HandlerUserInput(printCurriculumMenu)
		_, ok := menuManager.Actions[choice]
		if !ok {
			fmt.Println("Invalid option")
			continue
		}

		err := menuManager.Execute(choice)
		if err != nil {
			if errors.Is(err, ExitCommand) {
				return
			}
			fmt.Println("Error executing choice:", err)
		}
	}
}

func printCurriculumMenu() {
	fmt.Println("\nCurriculum Menu:")
	fmt.Println("1. Create Seed Curriculum")
	fmt.Println("2. List all Curriculums")
	fmt.Println("3. Get Curriculum by Id")
	fmt.Println("4. Update Curriculum by Id")
	fmt.Println("5. Delete Curriculum by Id")
	fmt.Println("0. Exit")
}

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
