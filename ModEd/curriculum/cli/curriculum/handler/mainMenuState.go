package handler

import (
	controller "ModEd/curriculum/controller"
)

type CurriculumCLIParams struct {
	CurriculumController controller.CurriculumControllerInterface
	CourseController     controller.CourseControllerInterface
	ClassController      controller.ClassControllerInterface
}

// NewMainMenuState creates the main menu for the curriculum module
func NewMainMenuState(params *CurriculumCLIParams) MenuState {
	mainState := NewBaseMenuState("Curriculum Module", nil)

	curriculumHandler := newCurriculumHandler(params.CurriculumController)
	courseHandler := newCourseHandler(params.CourseController)
	classHandler := newClassHandler(params.ClassController)

	curriculumState := NewCurriculumMenuState(curriculumHandler, mainState)
	courseState := NewCourseMenuState(courseHandler, mainState)
	classState := NewClassMenuState(classHandler, mainState)

	mainState.AddMenuItem("1", "Curriculum", func() (MenuState, error) {
		return curriculumState, nil
	})

	mainState.AddMenuItem("2", "Course", func() (MenuState, error) {
		return courseState, nil
	})

	mainState.AddMenuItem("3", "Class", func() (MenuState, error) {
		return classState, nil
	})

	mainState.AddExitItem()

	return mainState
}
