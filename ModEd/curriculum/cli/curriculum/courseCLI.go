package curriculum

import (
	controller "ModEd/curriculum/controller/curriculum"
	"ModEd/curriculum/utils"
	"fmt"
)

const (
	defaultCourseDataPath = "../../data/curriculum/course.json"
)

func RunCourseCLI(courseController *controller.CourseController) {

	for {
		printCourseMenu()
		choice := utils.GetUserChoice()

		switch choice {
		case "1":
			dataPath := utils.GetInputDataPath("course", defaultCourseDataPath)
			_, err := courseController.CreateSeedCourse(dataPath)
			if err != nil {
				fmt.Println("Error creating seed course:", err)
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

func printCourseMenu() {
	fmt.Println("\nCourse Menu:")
	fmt.Println("1. Create Seed Course")
	fmt.Println("2. Not implemented yet...")
	fmt.Println("3. Not implemented yet...")
	fmt.Println("0. Exit")
}
