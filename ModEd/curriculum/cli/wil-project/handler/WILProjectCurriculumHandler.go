// MEP-1010 Work Integrated Learning (WIL)
package handler

import (
	"ModEd/core"
	"ModEd/core/cli"
	"ModEd/core/handler"
	"ModEd/curriculum/controller"
	"ModEd/curriculum/model"
	"errors"
	"fmt"
)

type WILProjectCurriculumMenuStateHandler struct {
	manager                   *cli.CLIMenuStateManager
	wrapper                   *controller.WILModuleWrapper
	wilModuleMenuStateHandler *WILModuleMenuStateHandler
	handler                   *handler.HandlerContext
	backHandler               *handler.ChangeMenuHandlerStrategy
}

func NewWILProjectCurriculumMenuStateHandler(
	manager *cli.CLIMenuStateManager, wrapper *controller.WILModuleWrapper, wilModuleMenuStateHandler *WILModuleMenuStateHandler,
) *WILProjectCurriculumMenuStateHandler {

	return &WILProjectCurriculumMenuStateHandler{
		manager:                   manager,
		wrapper:                   wrapper,
		wilModuleMenuStateHandler: wilModuleMenuStateHandler,
		handler:                   handler.NewHandlerContext(),
		backHandler:               handler.NewChangeMenuHandlerStrategy(manager, wilModuleMenuStateHandler),
	}
}

func (menu *WILProjectCurriculumMenuStateHandler) Render() {
	menu.handler.SetMenuTitle("\nWIL Project Curriculum Menu:")
	menu.handler.AddHandler("1", "Create WIL Course", handler.FuncStrategy{Action: menu.createWILCourse})
	menu.handler.AddHandler("2", "Create WIL Class", handler.FuncStrategy{Action: menu.createWILClass})
	menu.handler.AddHandler("3", "List all of WIL Course", handler.FuncStrategy{Action: menu.listWILCourse})
	menu.handler.AddHandler("4", "List all of WIL Class", handler.FuncStrategy{Action: menu.listWILClass})
	menu.handler.AddBackHandler(menu.backHandler)

	menu.handler.ShowMenu()
}

func (menu *WILProjectCurriculumMenuStateHandler) HandleUserInput(input string) error {
	err := menu.handler.HandleInput(input)
	if err != nil {
		return err
	}
	return nil
}

func (menu *WILProjectCurriculumMenuStateHandler) createWILCourse() error {
	courseName := core.ExecuteUserInputStep(core.StringInputStep{
		PromptText:    "Enter course name:",
		FieldNameText: "Course Name",
	}).(string)
	description := core.ExecuteUserInputStep(core.StringInputStep{
		PromptText:    "Enter course description:",
		FieldNameText: "Course Description",
	}).(string)
	semester := core.ExecuteUserInputStep(core.StringInputStep{
		PromptText:    "Enter semester:",
		FieldNameText: "Semester",
	}).(string)

	course := &model.Course{
		Name:         courseName,
		Description:  description,
		CourseStatus: model.ACTIVE,
	}

	courseId, err := menu.wrapper.WILProjectCurriculumController.CreateNewWILCourse(course, semester)
	if err != nil {
		return errors.New("error! cannot create WIL course: " + err.Error())
	} else {
		fmt.Printf("WIL Course created successfully with ID: %d\n", courseId)
	}
	return nil
}

func (menu *WILProjectCurriculumMenuStateHandler) createWILClass() error {
	courseId := core.ExecuteUserInputStep(core.UintInputStep{
		PromptText:    "Enter course Id:",
		FieldNameText: "Course Id",
	}).(uint)
	section := core.ExecuteUserInputStep(core.UintInputStep{
		PromptText:    "Enter section:",
		FieldNameText: "Section",
	}).(uint)

	class := &model.Class{
		CourseId: courseId,
		Section:  int(section),
	}

	classId, err := menu.wrapper.WILProjectCurriculumController.CreateNewWILClass(class)
	if err != nil {
		return errors.New("error! cannot create WIL class: " + err.Error())
	} else {
		fmt.Printf("WIL Class created successfully with ID: %d\n", classId)
	}
	return nil
}

func (menu *WILProjectCurriculumMenuStateHandler) listWILCourse() error {
	courses, err := menu.wrapper.WILProjectCurriculumController.RetrieveAllWILCourses()
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
	classes, err := menu.wrapper.WILProjectCurriculumController.RetrieveAllWILClasses()
	if err != nil {
		return errors.New("error! cannot retrieve WIL classes")
	}

	fmt.Println("------[WIL Class]------")
	for i, class := range classes {
		fmt.Printf("%d. %d %d %s %s\n", i+1, class.ClassId, class.Section, class.CreatedAt.String(), class.UpdatedAt.String())
	}
	return nil
}
