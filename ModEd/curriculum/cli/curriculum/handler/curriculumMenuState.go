package handler

type CurriculumMenuState struct {
	*BaseMenuState
	handler *curriculumHandler
}

func NewCurriculumMenuState(handler *curriculumHandler, parent MenuState) *CurriculumMenuState {
	state := &CurriculumMenuState{
		BaseMenuState: NewBaseMenuState("Curriculum", parent),
		handler:       handler,
	}

	state.AddMenuItem("1", "Create New Curriculum", func() (MenuState, error) {
		err := handler.createCurriculum()
		return state, err
	})

	state.AddMenuItem("2", "Create Seed Curriculum", func() (MenuState, error) {
		err := handler.createSeedCurriculum()
		return state, err
	})

	state.AddMenuItem("3", "List all Curriculums", func() (MenuState, error) {
		err := handler.listCurriculums()
		return state, err
	})

	state.AddMenuItem("4", "Get Curriculum by Id", func() (MenuState, error) {
		err := handler.getCurriculumById()
		return state, err
	})

	state.AddMenuItem("5", "Update Curriculum by Id", func() (MenuState, error) {
		err := handler.updateCurriculumById()
		return state, err
	})

	state.AddMenuItem("6", "Delete Curriculum by Id", func() (MenuState, error) {
		err := handler.deleteCurriculumById()
		return state, err
	})

	if parent == nil {
		state.AddExitItem()
	} else {
		state.AddBackItem()
	}

	return state
}
