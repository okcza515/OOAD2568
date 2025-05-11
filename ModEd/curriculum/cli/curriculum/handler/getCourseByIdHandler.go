package handler

import (
	"ModEd/curriculum/controller"
	"ModEd/curriculum/utils"
	"fmt"
)

type getCourseByIdHandler struct {
	courseController controller.CourseControllerInterface
}

func NewGetCourseByIdHandler(courseController controller.CourseControllerInterface) *getCourseByIdHandler {
	return &getCourseByIdHandler{
		courseController: courseController,
	}
}

func (h *getCourseByIdHandler) Execute() error {
	courseId := utils.GetUserInputUint("Enter the course ID: ")
	course, err := h.courseController.GetCourse(courseId)
	if err != nil {
		fmt.Println("Error getting course:", err)
		return err
	}
	course.Print()
	return nil
}
