// MEP-1002
package handler

import (
	controller "ModEd/curriculum/controller"
	"ModEd/curriculum/model"
	"ModEd/curriculum/utils"
	"fmt"
	"strconv"
)

const (
	defaultCourseDataPath = "../data/curriculum/course.json"
)

type courseHandler struct {
	courseController     controller.CourseControllerInterface
	curriculumController controller.CurriculumControllerInterface
}

func newCourseHandler(courseController controller.CourseControllerInterface, curriculumController controller.CurriculumControllerInterface) *courseHandler {
	return &courseHandler{
		courseController:     courseController,
		curriculumController: curriculumController,
	}
}

func (h *courseHandler) createSeedCourse() (err error) {
	dataPath := utils.GetInputDataPath("course", defaultCourseDataPath)
	_, err = h.courseController.CreateSeedCourse(dataPath)
	if err != nil {
		fmt.Println("Error creating seed course:", err)
		return err
	}
	return nil
}

func (h *courseHandler) listCourses() (err error) {
	courses, err := h.courseController.GetCourses()
	if err != nil {
		fmt.Println("Error getting courses:", err)
		return err
	}

	for _, course := range courses {
		course.Print()
	}
	return nil
}

func (h *courseHandler) getCourseById() (err error) {
	courseId := utils.GetUserInputUint("Enter the course ID: ")
	course, err := h.courseController.GetCourse(courseId)
	if err != nil {
		fmt.Println("Error getting course:", err)
		return err
	}
	course.Print()
	return nil
}

func (h *courseHandler) updateCourseById() (err error) {
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

	// Update Optional flag
	fmt.Println("Is this course optional?")
	fmt.Println("1. Yes")
	fmt.Println("2. No")
	optionalChoice := utils.GetUserInput(fmt.Sprintf("Optional [%v] (1/2): ", course.Optional))
	if optionalChoice == "1" {
		course.Optional = true
	} else if optionalChoice == "2" {
		course.Optional = false
	}

	// Update CourseStatus
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

	// Validate updated course
	if err := course.Validate(); err != nil {
		fmt.Println("Validation error:", err)
		return err
	}

	// Confirm update
	confirm := utils.GetUserInput("Are you sure you want to update this course? (y/n): ")
	if confirm != "y" {
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

func (h *courseHandler) deleteCourseById() (err error) {
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

func (h *courseHandler) createCourse() (err error) {
	fmt.Println("\nCreate New Course:")

	// List available curriculums for reference
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
	if optionalChoice == "1" {
		optional = true
	} else if optionalChoice == "2" {
		optional = false
	}

	fmt.Println("Course Status options:")
	for key, value := range model.CourseStatusLabel {
		fmt.Printf("%d. %s\n", key+1, value)
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

	fmt.Println("\nCourse created successfully:")

	return nil
}
