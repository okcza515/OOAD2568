// MEP-1010 Work Integrated Learning (WIL)
package handler

import (
	"ModEd/asset/util"
	"ModEd/core/cli"
	"ModEd/curriculum/controller"
	"ModEd/curriculum/model"
	"ModEd/curriculum/utils"
	"errors"
	"fmt"
)

type WILProjectCurriculumMenuStateHandler struct {
	manager *cli.CLIMenuStateManager
	proxy   *controller.WILModuleProxy

	wilModuleMenuStateHandler *WILModuleMenuStateHandler
}

func NewWILProjectCurriculumMenuStateHandler(
	manager *cli.CLIMenuStateManager, proxy *controller.WILModuleProxy, wilModuleMenuStateHandler *WILModuleMenuStateHandler,
) *WILProjectCurriculumMenuStateHandler {
	return &WILProjectCurriculumMenuStateHandler{
		manager:                   manager,
		proxy:                     proxy,
		wilModuleMenuStateHandler: wilModuleMenuStateHandler,
	}
}

func (menu *WILProjectCurriculumMenuStateHandler) Render() {
	fmt.Println("\nWIL Project Curriculum Menu:")
	fmt.Println("1. Create WIL Course")
	fmt.Println("2. Create WIL Class")
	fmt.Println("3. List all of WIL Course")
	fmt.Println("4. List all of WIL Class")
	fmt.Println("0. Exit WIL Curriculum")
}

func (menu *WILProjectCurriculumMenuStateHandler) HandleUserInput(input string) error {

	switch input {
	case "1":
		err := menu.createWILCourse()
		if err != nil {
			fmt.Println("error! cannot use this function")
		}
	case "2":
		err := menu.createWILClass()
		if err != nil {
			fmt.Println("error! cannot use this function")
		}
	case "3":
		err := menu.listWILCourse()
		if err != nil {
			fmt.Println("error! cannot use this function")
		}
	case "4":
		err := menu.listWILClass()
		if err != nil {
			fmt.Println("error! cannot use this function")
		}
	case "0":
		menu.manager.SetState(menu.wilModuleMenuStateHandler)
		return nil
	default:
		fmt.Println("Invalid Command")
	}

	util.PressEnterToContinue()

	return nil
}

func (menu *WILProjectCurriculumMenuStateHandler) createWILCourse() error {
	courseName := utils.GetUserInput("Enter course name:")
	description := utils.GetUserInput("Enter course description:")
	semester := utils.GetUserInput("Enter semester:")

	course := &model.Course{
		Name:         courseName,
		Description:  description,
		CourseStatus: model.ACTIVE,
	}

	courseId, err := menu.proxy.WILProjectCurriculumController.CreateNewWILCourse(course, semester)
	if err != nil {
		return errors.New("error! cannot create WIL course: " + err.Error())
	} else {
		fmt.Printf("WIL Course created successfully with ID: %d\n", courseId)
	}
	return nil
}

func (menu *WILProjectCurriculumMenuStateHandler) createWILClass() error {
	courseId := utils.GetUserInputUint("Enter course Id:")
	section := utils.GetUserInputUint("Enter section:")

	class := &model.Class{
		CourseId: courseId,
		Section:  int(section),
	}

	classId, err := menu.proxy.WILProjectCurriculumController.CreateNewWILClass(class)
	if err != nil {
		return errors.New("error! cannot create WIL class: " + err.Error())
	} else {
		fmt.Printf("WIL Class created successfully with ID: %d\n", classId)
	}
	return nil
}

func (menu *WILProjectCurriculumMenuStateHandler) listWILCourse() error {
	courses, err := menu.proxy.WILProjectCurriculumController.RetrieveAllWILCourses()
	if err != nil {
		return errors.New("error! cannot retrieve WIL courses")
	}

	fmt.Println("------[WIL Course]------")
	for i, course := range courses {
		fmt.Printf("%d. %d %s %s %s %s\n", i+1, course.CourseId, course.Name, string(rune(course.CourseStatus)), course.CreatedAt.String(), course.UpdatedAt.String())
	}
	return nil
}

func (menu *WILProjectCurriculumMenuStateHandler) listWILClass() error {
	classes, err := menu.proxy.WILProjectCurriculumController.RetrieveAllWILClasses()
	if err != nil {
		return errors.New("error! cannot retrieve WIL classes")
	}

	fmt.Println("------[WIL Class]------")
	for i, class := range classes {
		fmt.Printf("%d. %d %d %s %s\n", i+1, class.ClassId, class.Section, class.CreatedAt.String(), class.UpdatedAt.String())
	}
	return nil
}
