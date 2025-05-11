package handler

import (
	"ModEd/curriculum/controller"
	"fmt"
)

type listCoursesHandler struct {
	courseController controller.CourseControllerInterface
}

func NewListCoursesHandler(courseController controller.CourseControllerInterface) *listCoursesHandler {
	return &listCoursesHandler{
		courseController: courseController,
	}
}

func (h *listCoursesHandler) Execute() error {
	courses, err := h.courseController.GetCourses()
	if err != nil {
		fmt.Println("Error listing courses:", err)
		return err
	}

	if len(courses) == 0 {
		fmt.Println("No courses found.")
		return nil
	}

	for _, course := range courses {
		course.Print()
	}
	return nil
}
