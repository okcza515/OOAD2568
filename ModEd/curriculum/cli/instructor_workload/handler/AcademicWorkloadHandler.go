package handler

import (
	"ModEd/curriculum/controller"
	"ModEd/curriculum/model"
	"ModEd/curriculum/utils"
	"fmt"
	"time"
)

func RunAcademicWorkloadHandler(
	coursePlanController controller.CoursePlanService,
	classWorkloadController controller.ClassWorkloadService,
) {
	for {
		DisplayAcademicWorkloadModuleMenu()
		choice := utils.GetUserChoice()
		fmt.Println("choice: ", choice)

		switch choice {
		case "1":
			mockLecture := model.ClassLecture{
				ClassId:      1,
				LectureName:  "Introduction to OOAD",
				InstructorId: 1,
				StartTime:    time.Now(),
				EndTime:      time.Now().Add(2 * time.Hour),
			}
			classWorkloadController.AddClassLecture(&mockLecture)
		case "2":
			mockEditedLecture := model.ClassLecture{
				ClassId:      1,
				LectureName:  "Introduction to OOAD - Hardcore",
				InstructorId: 1,
				StartTime:    time.Now(),
				EndTime:      time.Now().Add(2 * time.Hour),
			}
			classWorkloadController.UpdateClassLecture(&mockEditedLecture)
		case "3":
			classWorkloadController.DeleteClassLecture(1)
		case "4":
			classWorkloadController.GetClassLecturesByClassId(1)

		case "5":
			mockMaterial := model.ClassMaterial{
				ClassId:  1,
				FileName: "Lecture1.pdf",
				FilePath: "path/to/lecture1.pdf",
			}
			classWorkloadController.AddClassMaterial(&mockMaterial)
		case "6":
			editedMockMaterial := model.ClassMaterial{
				ClassId:  1,
				FileName: "Lecture1_edited.pdf",
				FilePath: "path/to/lecture1_edited.pdf",
			}
			classWorkloadController.UpdateClassMaterial(&editedMockMaterial)
		case "7":
			classWorkloadController.DeleteClassMaterial(1)
		case "8":
			classWorkloadController.GetClassMaterialsByClassId(1)
		case "exit":
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid option")
		}
	}
}

func DisplayAcademicWorkloadModuleMenu() {
	fmt.Println("\nAcademic Workload Menu:")
	fmt.Println("1. Add Class Lecture")
	fmt.Println("2. Edit Class Lecture")
	fmt.Println("3. Delete Class Lecture")
	fmt.Println("4. List all Class Lectures")

	fmt.Println("5. Add Class Material")
	fmt.Println("6. Edit Class Material")
	fmt.Println("7. Delete Class Material")
	fmt.Println("8. Get Class Material By ID")

	fmt.Println("Type 'exit' to quit")
}
