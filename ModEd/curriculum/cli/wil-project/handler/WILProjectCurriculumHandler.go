// MEP-1010 Work Integrated Learning (WIL)
package handler

import (
	"ModEd/curriculum/controller"
	"ModEd/curriculum/model"
	"ModEd/curriculum/utils"
	"fmt"
)

func RunWILProjectCurriculumHandler(controller *controller.WILModuleFacade) {
	for {
		printWILProjectCurriculumMenu()
		choice := utils.GetUserChoice()

		switch choice {
		case "1":
			courseName := utils.GetUserInput("Enter course name:")
			description := utils.GetUserInput("Enter course description:")
			semester := utils.GetUserInput("Enter semester:")

			course := &model.Course{
				Name:         courseName,
				Description:  description,
				CourseStatus: model.ACTIVE,
			}

			courseId, err := controller.WILProjectCurriculumController.CreateNewWILCourse(course, semester)
			if err != nil {
				fmt.Println("Error! cannot create WIL course:", err)
			} else {
				fmt.Printf("WIL Course created successfully with ID: %d\n", courseId)
			}

		case "2":
			courseId := utils.GetUserInputUint("Enter course Id:")
			section := utils.GetUserInputUint("Enter section:")

			class := &model.Class{
				CourseId: courseId,
				Section:  int(section),
			}

			classId, err := controller.WILProjectCurriculumController.CreateNewWILClass(class)
			if err != nil {
				fmt.Println("Error! cannot create WIL class:", err)
			} else {
				fmt.Printf("WIL Class created successfully with ID: %d\n", classId)
			}
		case "3":
			courses, err := controller.WILProjectCurriculumController.RetrieveAllWILCourses()
			if err != nil {
				fmt.Println("Error! cannot retrieve WIL courses")
			}

			fmt.Println("------[WIL Course]------")
			for i, course := range courses {
				fmt.Printf("%d. %d %s %s %s %s\n", i+1, course.CourseId, course.Name, string(rune(course.CourseStatus)), course.CreatedAt.String(), course.UpdatedAt.String())
			}

		case "4":
			classes, err := controller.WILProjectCurriculumController.RetrieveAllWILClasses()
			if err != nil {
				fmt.Println("Error! cannot retrieve WIL classes")
			}

			fmt.Println("------[WIL Class]------")
			for i, class := range classes {
				fmt.Printf("%d. %d %d %s %s\n", i+1, class.ClassId, class.Section, class.CreatedAt.String(), class.UpdatedAt.String())
			}
		case "0":
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid option")
		}
	}
}

func printWILProjectCurriculumMenu() {
	fmt.Println("\nWIL Project Curriculum Menu:")
	fmt.Println("1. Create WIL Course")
	fmt.Println("2. Create WIL Class")
	fmt.Println("3. List all of WIL Course")
	fmt.Println("4. List all of WIL Class")
	fmt.Println("0. Exit WIL Curriculum")
}
