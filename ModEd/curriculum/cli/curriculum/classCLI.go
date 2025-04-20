package curriculum

import (
	controller "ModEd/curriculum/controller/curriculum"
	"ModEd/curriculum/utils"
	"fmt"
)

const (
	defaultClassDataPath = "../../data/curriculum/class.json"
)

func RunClassCLI(classController *controller.ClassController) {
	for {
		printClassMenu()
		choice := utils.GetUserChoice()

		switch choice {
		case "1":
			dataPath := utils.GetInputDataPath("class", defaultClassDataPath)
			_, err := classController.CreateSeedClass(dataPath)
			if err != nil {
				fmt.Println("Error creating seed class:", err)
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

func printClassMenu() {
	fmt.Println("\nClass Menu:")
	fmt.Println("1. Create Seed Class")
	fmt.Println("2. Not implemented yet...")
	fmt.Println("3. Not implemented yet...")
	fmt.Println("0. Exit")
}
