// MEP-1002
package handler

import (
	controller "ModEd/curriculum/controller"
	"ModEd/curriculum/model"
	"ModEd/curriculum/utils"
	"fmt"
	"strconv"
	"time"
)

const (
	defaultClassDataPath = "../data/curriculum/class.json"
)

type classHandler struct {
	classController  controller.ClassControllerInterface
	courseController controller.CourseControllerInterface
}

func newClassHandler(classController controller.ClassControllerInterface, courseController controller.CourseControllerInterface) *classHandler {
	return &classHandler{
		classController:  classController,
		courseController: courseController,
	}
}

func (h *classHandler) createSeedClass() (err error) {
	dataPath := utils.GetInputDataPath("class", defaultClassDataPath)
	_, err = h.classController.CreateSeedClass(dataPath)
	if err != nil {
		fmt.Println("Error creating seed class:", err)
		return err
	}
	return nil
}

func (h *classHandler) listClasses() (err error) {
	classes, err := h.classController.GetClasses()
	if err != nil {
		fmt.Println("Error getting classes:", err)
		return err
	}

	for _, class := range classes {
		class.Print()
	}
	return nil
}

func (h *classHandler) getClassById() (err error) {
	classId := utils.GetUserInputUint("Enter the class ID: ")
	class, err := h.classController.GetClass(classId)
	if err != nil {
		fmt.Println("Error getting class:", err)
		return err
	}
	class.Print()
	return nil
}

func (h *classHandler) updateClassById() (err error) {
	classId := utils.GetUserInputUint("Enter the class ID: ")
	class, err := h.classController.GetClass(classId)
	if err != nil {
		fmt.Println("Error getting class:", err)
		return err
	}

	fmt.Println("\nCurrent class information:")
	class.Print()

	fmt.Println("\nEnter new values (leave blank to keep current value):")

	// Update CourseId
	newCourseId := utils.GetUserInput(fmt.Sprintf("Course ID [%d]: ", class.CourseId))
	if newCourseId != "" {
		courseId, err := strconv.Atoi(newCourseId)
		if err == nil {
			class.CourseId = uint(courseId)
		} else {
			fmt.Println("Invalid course ID format, keeping current value")
		}
	}

	// Update Section
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
		newSchedule := utils.GetUserInput(fmt.Sprintf("Schedule [%s]: ", class.Schedule.Format("2006-01-02 15:04:05")))
		if newSchedule != "" {
			schedule, err := time.Parse("2006-01-02 15:04:05", newSchedule)
			if err == nil {
				class.Schedule = schedule
			} else {
				fmt.Println("Invalid schedule format, keeping current value. Use format: YYYY-MM-DD HH:MM:SS")
			}
		}
	} else {
		newSchedule := utils.GetUserInput("Schedule [none]: ")
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
	if confirm != "y" {
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

func (h *classHandler) deleteClassById() (err error) {
	classes, err := h.classController.GetClasses()
	if err != nil {
		fmt.Println("Error getting classes:", err)
		return err
	}

	for _, class := range classes {
		class.Print()
	}

	classId := utils.GetUserInputUint("Enter the class ID to delete: ")

	confirm := utils.GetUserInput(fmt.Sprintf("Are you sure you want to delete class with Id %d? (y/n): ", classId))
	if confirm != "y" {
		fmt.Println("Deletion cancelled.")
		return nil
	}

	_, err = h.classController.DeleteClass(classId)
	if err != nil {
		fmt.Println("Error deleting class:", err)
		return err
	}

	fmt.Println("Class deleted successfully!")
	return nil
}

func (h *classHandler) createClass() (err error) {
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

	fmt.Println("\nClass created successfully:")

	return nil
}
