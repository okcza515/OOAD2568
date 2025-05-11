package handler

import (
	"ModEd/curriculum/controller"
	"ModEd/curriculum/model"
	"ModEd/curriculum/utils"
	"fmt"
	"strconv"
	"time"
)

type createClassHandler struct {
	classController  controller.ClassControllerInterface
	courseController controller.CourseControllerInterface
}

func NewCreateClassHandler(classController controller.ClassControllerInterface, courseController controller.CourseControllerInterface) *createClassHandler {
	return &createClassHandler{
		classController:  classController,
		courseController: courseController,
	}
}

func (h *createClassHandler) Execute() error {
	fmt.Println("\nCreate New Class:")

	// List available courses for reference
	courses, err := h.courseController.GetCourses()
	if err != nil {
		fmt.Println("Error fetching courses:", err)
		return err
	}

	if len(courses) == 0 {
		fmt.Println("No courses found. Please create a course first.")
		return nil
	}

	fmt.Println("\nAvailable Courses:")
	for _, course := range courses {
		fmt.Printf("ID: %d, Name: %s\n", course.CourseId, course.Name)
	}

	courseIdStr := utils.GetUserInput("\nEnter course ID: ")
	courseId, err := strconv.Atoi(courseIdStr)
	if err != nil {
		fmt.Println("Invalid course ID format")
		return err
	}

	sectionStr := utils.GetUserInput("Enter section number: ")
	section, err := strconv.Atoi(sectionStr)
	if err != nil {
		fmt.Println("Invalid section number")
		return err
	}

	scheduleStr := utils.GetUserInput("Enter schedule (YYYY-MM-DD HH:MM:SS) or leave blank for none: ")
	var schedule time.Time
	if scheduleStr != "" {
		schedule, err = time.Parse("2006-01-02 15:04:05", scheduleStr)
		if err != nil {
			fmt.Println("Invalid schedule format. Expected: YYYY-MM-DD HH:MM:SS")
			return err
		}
	}

	class := &model.Class{
		CourseId: uint(courseId),
		Section:  section,
		Schedule: schedule,
	}

	if err := class.Validate(); err != nil {
		fmt.Println("Validation error:", err)
		return err
	}

	fmt.Println("\nClass to be created:")
	fmt.Printf("Course ID: %d\n", class.CourseId)
	fmt.Printf("Section: %d\n", class.Section)

	if !class.Schedule.IsZero() {
		fmt.Printf("Schedule: %s\n", class.Schedule.Format("2006-01-02 15:04:05"))
	} else {
		fmt.Println("Schedule: None")
	}

	confirm := utils.GetUserInput("\nConfirm creation? (y/n): ")
	if confirm != "y" {
		fmt.Println("Creation cancelled.")
		return nil
	}

	_, err = h.classController.CreateClass(class)
	if err != nil {
		fmt.Println("Error creating class:", err)
		return err
	}

	fmt.Println("\nClass created successfully!")
	return nil
}
