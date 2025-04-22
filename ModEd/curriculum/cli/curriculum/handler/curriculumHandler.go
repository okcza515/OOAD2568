// MEP-1002
package handler

import (
	controller "ModEd/curriculum/controller"
	"ModEd/curriculum/utils"
	"fmt"
)

const (
	defaultCurriculumDataPath = "../../data/curriculum/curriculum.json"
)

func RunCurriculumCLIHandler(curriculumController controller.CurriculumControllerInterface) {
	for {
		printCurriculumMenu()
		choice := utils.GetUserChoice()

		switch choice {
		case "1":
			dataPath := utils.GetInputDataPath("curriculum", defaultCurriculumDataPath)
			_, err := curriculumController.CreateSeedCurriculum(dataPath)
			if err != nil {
				fmt.Println("Error creating seed curriculum:", err)
			}
			return
		case "2":
			err := listCurriculums(curriculumController)
			if err != nil {
				fmt.Println("Error listing curriculums:", err)
			}
		case "3":
			err := getCurriculumById(curriculumController)
			if err != nil {
				fmt.Println("Error getting curriculum:", err)
			}
		case "4":
			err := updateCurriculumById(curriculumController)
			if err != nil {
				fmt.Println("Error updating curriculum:", err)
			}
		case "5":
			err := deleteCurriculumById(curriculumController)
			if err != nil {
				fmt.Println("Error deleting curriculum:", err)
			}
		case "0":
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid option")
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

func listCurriculums(curriculumController controller.CurriculumControllerInterface) (err error) {
	curriculums, err := curriculumController.GetCurriculums()
	if err != nil {
		fmt.Println("Error getting curriculums:", err)
		return err
	}

	for _, curriculum := range curriculums {
		curriculum.Print()
	}
	return nil
}

func getCurriculumById(curriculumController controller.CurriculumControllerInterface) (err error) {
	curriculumId := utils.GetUserInputUint("Enter the curriculum ID: ")
	curriculum, err := curriculumController.GetCurriculum(curriculumId)
	if err != nil {
		fmt.Println("Error getting curriculum:", err)
		return err
	}
	curriculum.Print()
	return nil
}

func updateCurriculumById(curriculumController controller.CurriculumControllerInterface) (err error) {
	return nil
}
func deleteCurriculumById(curriculumController controller.CurriculumControllerInterface) (err error) {
	curriculums, err := curriculumController.GetCurriculums()
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

	_, err = curriculumController.DeleteCurriculum(curriculumId)
	if err != nil {
		fmt.Println("Error deleting curriculum:", err)
		return err
	}

	fmt.Println("Curriculum deleted successfully!")
	return nil
}
