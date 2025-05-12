package handler

import (
	"ModEd/curriculum/controller"
	"ModEd/curriculum/utils"
	"fmt"
	"strconv"
	"time"
)

type updateClassByIdHandler struct {
	classController  controller.ClassControllerInterface
	courseController controller.CourseControllerInterface
}

func NewUpdateClassByIdHandler(classController controller.ClassControllerInterface, courseController controller.CourseControllerInterface) *updateClassByIdHandler {
	return &updateClassByIdHandler{
		classController:  classController,
		courseController: courseController,
	}
}

var confirmOptions = map[string]bool{
	"y": true,
	"n": false,
}

func (h *updateClassByIdHandler) Execute() error {
	classId := utils.GetUserInputUint("Enter the class ID: ")
	class, err := h.classController.GetClass(classId)
	if err != nil {
		fmt.Println("Error getting class:", err)
		return err
	}

	fmt.Println("\nCurrent class information:")
	class.Print()

	fmt.Println("\nEnter new values (leave blank to keep current value):")

	newCourseId := utils.GetUserInput(fmt.Sprintf("Course ID [%d]: ", class.CourseId))
	if newCourseId != "" {
		courseId, err := strconv.Atoi(newCourseId)
		if err == nil {
			class.CourseId = uint(courseId)
		} else {
			fmt.Println("Invalid course ID format, keeping current value")
		}
	}

	newSection := utils.GetUserInput(fmt.Sprintf("Section [%d]: ", class.Section))
	if newSection != "" {
		section, err := strconv.Atoi(newSection)
		if err == nil {
			class.Section = section
		} else {
			fmt.Println("Invalid section format, keeping current value")
		}
	}

	if !class.Schedule.IsZero() {
		newSchedule := utils.GetUserInput(fmt.Sprintf("Schedule [%s] format (YYYY-MM-DD HH:MM:SS): ", class.Schedule.Format("2006-01-02 15:04:05")))
		if newSchedule != "" {
			schedule, err := time.Parse("2006-01-02 15:04:05", newSchedule)
			if err == nil {
				class.Schedule = schedule
			} else {
				fmt.Println("Invalid schedule format, keeping current value. Use format: YYYY-MM-DD HH:MM:SS")
			}
		}
	} else {
		newSchedule := utils.GetUserInput("Schedule [none] format (YYYY-MM-DD HH:MM:SS): ")
		if newSchedule != "" {
			schedule, err := time.Parse("2006-01-02 15:04:05", newSchedule)
			if err == nil {
				class.Schedule = schedule
			} else {
				fmt.Println("Invalid schedule format, keeping no schedule. Use format: YYYY-MM-DD HH:MM:SS")
			}
		}
	}

	// Validate the updated class
	if err := class.Validate(); err != nil {
		fmt.Println("Validation error:", err)
		return err
	}

	confirm := utils.GetUserInput("Are you sure you want to update this class? (y/n): ")
	if confirmed, exists := confirmOptions[confirm]; !exists || !confirmed {
		fmt.Println("Update cancelled.")
		return nil
	}

	updatedClass, err := h.classController.UpdateClass(class)
	if err != nil {
		fmt.Println("Error updating class:", err)
		return err
	}

	fmt.Println("Class updated successfully:")
	updatedClass.Print()

	return nil
}
