package curriculum

import (
	controller "ModEd/curriculum/controller/curriculum"
	"ModEd/curriculum/utils"
	"fmt"

	"gorm.io/gorm"
)

const (
	defaultClassDataPath = "../../data/curriculum/class.json"
)

func RunClassCLI() {
	database := utils.GetInputDatabasePath("class", defaultDBPath)

	db, err := utils.NewGormSqlite(&utils.GormConfig{
		DBPath: database,
		Config: &gorm.Config{},
	})
	if err != nil {
		panic(err)
	}

	classController := controller.NewClassController(db)

	handleClassChoice(classController)
	return
}

func printClassMenu() {
	fmt.Println("\nClass Menu:")
	fmt.Println("1. Create Seed Class")
	fmt.Println("2. Not implemented yet...")
	fmt.Println("3. Not implemented yet...")
	fmt.Println("0. Exit")
}

func handleClassChoice(classController controller.IClassController) {
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
