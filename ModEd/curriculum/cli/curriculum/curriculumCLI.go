package curriculum

import (
	controller "ModEd/curriculum/controller/curriculum"
	"ModEd/curriculum/utils"
	"fmt"
)

const (
	defaultCurriculumDataPath = "../../data/curriculum/curriculum.json"
)

func RunCurriculumCLI(curriculumController *controller.CurriculumController) {
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
			fmt.Println("Not implemented yet...")
		case "3":
			fmt.Println("Not implemented yet...")
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
	fmt.Println("2. Not implemented yet...")
	fmt.Println("3. Not implemented yet...")
	fmt.Println("0. Exit")
}
