package handler

import (
	"ModEd/curriculum/controller"
	"ModEd/curriculum/model"
	"ModEd/curriculum/utils"
	"fmt"
	"strconv"
)

type updateCourseByIdHandler struct {
	courseController     controller.CourseControllerInterface
	curriculumController controller.CurriculumControllerInterface
}

func NewUpdateCourseByIdHandler(courseController controller.CourseControllerInterface, curriculumController controller.CurriculumControllerInterface) *updateCourseByIdHandler {
	return &updateCourseByIdHandler{
		courseController:     courseController,
		curriculumController: curriculumController,
	}
}

var optionalOptions = map[string]bool{
	"1": true,  // Yes
	"2": false, // No
}

func (h *updateCourseByIdHandler) Execute() error {
	courseId := utils.GetUserInputUint("Enter the course ID: ")
	course, err := h.courseController.GetCourse(courseId)
	if err != nil {
		fmt.Println("Error getting course:", err)
		return err
	}

	fmt.Println("\nCurrent course information:")
	course.Print()

	fmt.Println("\nEnter new values (leave blank to keep current value):")

	newName := utils.GetUserInput(fmt.Sprintf("Name [%s]: ", course.Name))
	if newName != "" {
		course.Name = newName
	}

	newDescription := utils.GetUserInput(fmt.Sprintf("Description [%s]: ", course.Description))
	if newDescription != "" {
		course.Description = newDescription
	}

	newCurriculumId := utils.GetUserInput(fmt.Sprintf("Curriculum ID [%d]: ", course.CurriculumId))
	if newCurriculumId != "" {
		curriculumId, err := strconv.Atoi(newCurriculumId)
		if err == nil {
			course.CurriculumId = uint(curriculumId)
		} else {
			fmt.Println("Invalid curriculum ID format, keeping current value")
		}
	}

	fmt.Println("Is this course optional?")
	fmt.Println("1. Yes")
	fmt.Println("2. No")
	optionalChoice := utils.GetUserInput(fmt.Sprintf("Optional [%v] (1/2): ", course.Optional))
	if newValue, exists := optionalOptions[optionalChoice]; exists {
		course.Optional = newValue
	}

	fmt.Println("Course Status options:")
	for key, value := range model.CourseStatusLabel {
		fmt.Printf("%d. %s\n", key, value)
	}
	statusChoice := utils.GetUserInput(fmt.Sprintf("Course Status [%s]: ", course.CourseStatus))
	if statusChoice != "" {
		status, err := strconv.Atoi(statusChoice)
		if err == nil && status >= 1 && status <= 2 {
			course.CourseStatus = model.CourseStatus(status)
		} else {
			fmt.Println("Invalid course status, keeping current value")
		}
	}

	if err := course.Validate(); err != nil {
		fmt.Println("Validation error:", err)
		return err
	}

	confirm := utils.GetUserInput("Are you sure you want to update this course? (y/n): ")
	if confirmed, exists := confirmOptions[confirm]; !exists || !confirmed {
		fmt.Println("Update cancelled.")
		return nil
	}

	updatedCourse, err := h.courseController.UpdateCourse(course)
	if err != nil {
		fmt.Println("Error updating course:", err)
		return err
	}

	fmt.Println("Course updated successfully:")
	updatedCourse.Print()

	return nil
}
