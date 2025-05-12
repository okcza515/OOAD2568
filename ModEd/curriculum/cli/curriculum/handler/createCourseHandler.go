package handler

import (
	"ModEd/curriculum/controller"
	"ModEd/curriculum/model"
	"ModEd/curriculum/utils"
	"fmt"
	"strconv"
)

type createCourseHandler struct {
	courseController     controller.CourseControllerInterface
	curriculumController controller.CurriculumControllerInterface
}

func NewCreateCourseHandler(courseController controller.CourseControllerInterface, curriculumController controller.CurriculumControllerInterface) *createCourseHandler {
	return &createCourseHandler{
		courseController:     courseController,
		curriculumController: curriculumController,
	}
}

func (h *createCourseHandler) Execute() error {
	fmt.Println("\nCreate New Course:")

	curriculums, err := h.curriculumController.GetCurriculums()
	if err != nil {
		fmt.Println("Error fetching curriculums:", err)
		return err
	}

	if len(curriculums) == 0 {
		fmt.Println("No curriculums found. Please create a curriculum first.")
		return nil
	}

	fmt.Println("\nAvailable Curriculums:")
	for _, curriculum := range curriculums {
		fmt.Printf("ID: %d, Name: %s\n", curriculum.CurriculumId, curriculum.Name)
	}

	name := utils.GetUserInput("\nEnter course name: ")
	description := utils.GetUserInput("Enter course description: ")

	curriculumIdStr := utils.GetUserInput("Enter curriculum ID: ")
	curriculumId, err := strconv.Atoi(curriculumIdStr)
	if err != nil {
		fmt.Println("Invalid curriculum ID format")
		return err
	}

	fmt.Println("Is this course optional?")
	fmt.Println("1. Yes")
	fmt.Println("2. No")
	optionalChoice := utils.GetUserInput("Select option (1/2): ")
	var optional bool

	if newValue, exists := optionalOptions[optionalChoice]; exists {
		optional = newValue
	}

	fmt.Println("Course Status options:")
	for key, value := range model.CourseStatusLabel {
		fmt.Printf("%d. %s\n", key, value)
	}

	statusStr := utils.GetUserInput("Select course status (enter number): ")
	status, err := strconv.Atoi(statusStr)
	if err != nil {
		fmt.Println("Invalid course status")
		return err
	}

	course := &model.Course{
		Name:         name,
		Description:  description,
		CurriculumId: uint(curriculumId),
		Optional:     optional,
		CourseStatus: model.CourseStatus(status),
	}

	if err := course.Validate(); err != nil {
		fmt.Println("Validation error:", err)
		return err
	}

	fmt.Println("\nCourse to be created:")
	fmt.Printf("Name: %s\n", course.Name)
	fmt.Printf("Description: %s\n", course.Description)
	fmt.Printf("Curriculum ID: %d\n", course.CurriculumId)
	fmt.Printf("Optional: %v\n", course.Optional)
	fmt.Printf("Status: %s\n", model.CourseStatusLabel[course.CourseStatus])

	confirm := utils.GetUserInput("\nConfirm creation? (y/n): ")
	if confirm != "y" {
		fmt.Println("Creation cancelled.")
		return nil
	}

	_, err = h.courseController.CreateCourse(course)
	if err != nil {
		fmt.Println("Error creating course:", err)
		return err
	}

	fmt.Println("\nCourse created successfully!")
	return nil
}
