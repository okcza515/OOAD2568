package handler

import (
	"ModEd/curriculum/controller"
	"ModEd/curriculum/utils"
	"fmt"
)

type deleteCourseByIdHandler struct {
	courseController controller.CourseControllerInterface
}

func NewDeleteCourseByIdHandler(courseController controller.CourseControllerInterface) *deleteCourseByIdHandler {
	return &deleteCourseByIdHandler{
		courseController: courseController,
	}
}

func (h *deleteCourseByIdHandler) Execute() error {
	courses, err := h.courseController.GetCourses()
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

	_, err = h.courseController.DeleteCourse(courseId)
	if err != nil {
		fmt.Println("Error deleting course:", err)
		return err
	}

	fmt.Println("Course deleted successfully!")
	return nil
}
