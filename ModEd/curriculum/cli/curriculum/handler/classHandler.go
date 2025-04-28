// MEP-1002
package handler

import (
	controller "ModEd/curriculum/controller"
	"ModEd/curriculum/utils"
	"fmt"
)

const (
	defaultClassDataPath = "../../data/curriculum/class.json"
)

func RunClassCLIHandler(classController controller.ClassControllerInterface) {
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
			err := listClasses(classController)
			if err != nil {
				fmt.Println("Error listing classes:", err)
			}
		case "3":
			err := getClassById(classController)
			if err != nil {
				fmt.Println("Error getting class:", err)
			}
		case "4":
			err := updateClassById(classController)
			if err != nil {
				fmt.Println("Error updating class:", err)
			}
		case "5":
			err := deleteClassById(classController)
			if err != nil {
				fmt.Println("Error deleting class:", err)
			}
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
	fmt.Println("2. List all Classes")
	fmt.Println("3. Get Class by Id")
	fmt.Println("4. Update Class by Id")
	fmt.Println("5. Delete Class by Id")
	fmt.Println("0. Exit")
}

func listClasses(classController controller.ClassControllerInterface) (err error) {
	classes, err := classController.GetClasses()
	if err != nil {
		fmt.Println("Error getting classes:", err)
		return err
	}

	for _, class := range classes {
		class.Print()
	}
	return nil
}

func getClassById(classController controller.ClassControllerInterface) (err error) {
	classId := utils.GetUserInputUint("Enter the class ID: ")
	class, err := classController.GetClass(classId)
	if err != nil {
		fmt.Println("Error getting class:", err)
		return err
	}
	class.Print()
	return nil
}

func updateClassById(classController controller.ClassControllerInterface) (err error) {
	//TODO: Implement update class by ID
	return nil
}

func deleteClassById(classController controller.ClassControllerInterface) (err error) {
	classes, err := classController.GetClasses()
	if err != nil {
		fmt.Println("Error getting classes:", err)
		return err
	}

	for _, class := range classes {
		class.Print()
	}

	classId := utils.GetUserInputUint("Enter the class ID to delete: ")

	confirm := utils.GetUserInput(fmt.Sprintf("Are you sure you want to delete class with Id %d? (y/n): ", classId))
	if confirm != "y" {
		fmt.Println("Deletion cancelled.")
		return nil
	}

	_, err = classController.DeleteClass(classId)
	if err != nil {
		fmt.Println("Error deleting class:", err)
		return err
	}

	fmt.Println("Class deleted successfully!")
	return nil
}
