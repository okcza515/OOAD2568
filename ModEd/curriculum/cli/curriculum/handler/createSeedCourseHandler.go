package handler

import (
	"ModEd/curriculum/controller"
	"ModEd/curriculum/utils"
	"fmt"
)

const (
	defaultCourseDataPath = "../data/curriculum/course.json"
)

type createSeedCourseHandler struct {
	courseController controller.CourseControllerInterface
}

func NewCreateSeedCourseHandler(courseController controller.CourseControllerInterface) *createSeedCourseHandler {
	return &createSeedCourseHandler{
		courseController: courseController,
	}
}

func (h *createSeedCourseHandler) Execute() error {
	dataPath := utils.GetInputDataPath("course", defaultCourseDataPath)
	_, err := h.courseController.CreateSeedCourse(dataPath)
	if err != nil {
		fmt.Println("Error creating seed course:", err)
		return err
	}
	fmt.Println("Seed course created successfully")
	return nil
}
