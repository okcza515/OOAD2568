package handler

import "ModEd/eval/controller"

type AssessmentCLIParams struct {
	AssessmentController controller.AssessmentController
	SubmissionController controller.SubmissionController
	ResultController     controller.ResultController
}

type MainMenuState struct {
	*BaseMenuState
	Params              *AssessmentCLIParams
	AssessmentMenuState MenuState
	SubmissionMenuState MenuState
	ResultMenuState     MenuState
}

func NewMainMenuState(params *AssessmentCLIParams) *MainMenuState {
	mainState := &MainMenuState{
		BaseMenuState: NewBaseMenuState("Assessment Module", nil),
		Params:        params,
	}

	// Create sub-states
	mainState.AssessmentMenuState = NewAssessmentMenuState(mainState, params)
	mainState.SubmissionMenuState = NewSubmissionMenuState(mainState, params)
	mainState.ResultMenuState = NewResultMenuState(mainState, params)

	// Configure menu items
	mainState.AddMenuItem("1", "Assessment Management", func() (MenuState, error) {
		return mainState.AssessmentMenuState, nil
	})

	mainState.AddMenuItem("2", "Submission Management", func() (MenuState, error) {
		return mainState.SubmissionMenuState, nil
	})

	mainState.AddMenuItem("3", "Result Management", func() (MenuState, error) {
		return mainState.ResultMenuState, nil
	})

	mainState.AddExitItem()

	return mainState
}
