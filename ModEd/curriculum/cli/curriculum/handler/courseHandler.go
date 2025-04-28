// MEP-1002
package handler

import (
	controller "ModEd/curriculum/controller"
	"ModEd/curriculum/utils"
	"fmt"
)

const (
	defaultCourseDataPath = "../../data/curriculum/course.json"
)

func RunCourseCLIHandler(courseController controller.CourseControllerInterface) {
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
			err := listCourses(courseController)
			if err != nil {
				fmt.Println("Error listing courses:", err)
			}
		case "3":
			err := getCourseById(courseController)
			if err != nil {
				fmt.Println("Error getting course:", err)
			}
		case "4":
			err := updateCourseById(courseController)
			if err != nil {
				fmt.Println("Error updating course:", err)
			}
		case "5":
			err := deleteCourseById(courseController)
			if err != nil {
				fmt.Println("Error deleting course:", err)
			}
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
	fmt.Println("2. List all Courses")
	fmt.Println("3. Get Course by Id")
	fmt.Println("4. Update Course by Id")
	fmt.Println("5. Delete Course by Id")
	fmt.Println("0. Exit")
}

func listCourses(courseController controller.CourseControllerInterface) (err error) {
	courses, err := courseController.GetCourses()
	if err != nil {
		fmt.Println("Error getting courses:", err)
		return err
	}

	for _, course := range courses {
		course.Print()
	}
	return nil
}

func getCourseById(courseController controller.CourseControllerInterface) (err error) {
	courseId := utils.GetUserInputUint("Enter the course ID: ")
	course, err := courseController.GetCourse(courseId)
	if err != nil {
		fmt.Println("Error getting course:", err)
		return err
	}
	course.Print()
	return nil
}

func updateCourseById(courseController controller.CourseControllerInterface) (err error) {

	return nil
}

func deleteCourseById(courseController controller.CourseControllerInterface) (err error) {
	courses, err := courseController.GetCourses()
	if err != nil {
		fmt.Println("Error getting courses:", err)
		return err
	}

	for _, course := range courses {
		course.Print()
	}

	courseId := utils.GetUserInputUint("Enter the course Id to delete: ")

	confirm := utils.GetUserInput(fmt.Sprintf("Are you sure you want to delete course with Id %d? (y/n): ", courseId))
	if confirm != "y" {
		fmt.Println("Deletion cancelled.")
		return nil
	}

	_, err = courseController.DeleteCourse(courseId)
	if err != nil {
		fmt.Println("Error deleting course:", err)
		return err
	}

	fmt.Println("Course deleted successfully!")
	return nil
}
