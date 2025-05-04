package handler

type CourseMenuState struct {
	*BaseMenuState
	handler *courseHandler
}

func NewCourseMenuState(handler *courseHandler, parent MenuState) *CourseMenuState {
	state := &CourseMenuState{
		BaseMenuState: NewBaseMenuState("Course", parent),
		handler:       handler,
	}

	state.AddMenuItem("1", "Create Seed Course", func() (MenuState, error) {
		err := handler.createSeedCourse()
		return state, err
	})

	state.AddMenuItem("2", "List all Courses", func() (MenuState, error) {
		err := handler.listCourses()
		return state, err
	})

	state.AddMenuItem("3", "Get Course by Id", func() (MenuState, error) {
		err := handler.getCourseById()
		return state, err
	})

	state.AddMenuItem("4", "Update Course by Id", func() (MenuState, error) {
		err := handler.updateCourseById()
		return state, err
	})

	state.AddMenuItem("5", "Delete Course by Id", func() (MenuState, error) {
		err := handler.deleteCourseById()
		return state, err
	})

	if parent == nil {
		state.AddExitItem()
	} else {
		state.AddBackItem()
	}

	return state
}
