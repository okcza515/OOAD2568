package handler

type ClassMenuState struct {
	*BaseMenuState
	handler *classHandler
}

func NewClassMenuState(handler *classHandler, parent MenuState) *ClassMenuState {
	state := &ClassMenuState{
		BaseMenuState: NewBaseMenuState("Class", parent),
		handler:       handler,
	}

	state.AddMenuItem("1", "Create New Class", func() (MenuState, error) {
		err := handler.createClass()
		return state, err
	})

	state.AddMenuItem("2", "Create Seed Class", func() (MenuState, error) {
		err := handler.createSeedClass()
		return state, err
	})

	state.AddMenuItem("3", "List all Classes", func() (MenuState, error) {
		err := handler.listClasses()
		return state, err
	})

	state.AddMenuItem("4", "Get Class by Id", func() (MenuState, error) {
		err := handler.getClassById()
		return state, err
	})

	state.AddMenuItem("5", "Update Class by Id", func() (MenuState, error) {
		err := handler.updateClassById()
		return state, err
	})

	state.AddMenuItem("6", "Delete Class by Id", func() (MenuState, error) {
		err := handler.deleteClassById()
		return state, err
	})

	if parent == nil {
		state.AddExitItem()
	} else {
		state.AddBackItem()
	}

	return state
}
