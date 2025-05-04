// MEP-1002
package handler

import (
	controller "ModEd/curriculum/controller"
	"ModEd/curriculum/utils"
	"fmt"
	"strconv"
	"time"
)

const (
	defaultClassDataPath = "../../data/curriculum/class.json"
)

type classHandler struct {
	classController controller.ClassControllerInterface
}

func newClassHandler(classController controller.ClassControllerInterface) *classHandler {
	return &classHandler{
		classController: classController,
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
