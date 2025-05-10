package handler

import "ModEd/eval/controller"

type AssessmentCLIParams struct {
	AssessmentController    controller.AssessmentController
	SubmissionController    controller.SubmissionController
	ResultController        controller.ResultController
	SubmissionPDFController *controller.SubmissionPDFController
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
	mainState.SubmissionMenuState = NewSubmissionMenuState(params)
	mainState.ResultMenuState = NewResultMenuState(params)

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

func NewAssessmentMenuState(parent MenuState, params *AssessmentCLIParams) MenuState {
	// Assuming a similar structure to other menu states
	assessmentState := &BaseMenuState{
		Name:   "Assessment Management",
		Parent: parent,
	}
	// Add menu items specific to assessment management here
	return assessmentState
}

// Add Enter method to MainMenuState
func (s *MainMenuState) Enter() error {
	// Implement the logic for entering the main menu state
	return nil
}

// Add Enter method to BaseMenuState
func (s *BaseMenuState) Enter() error {
	// Implement the logic for entering the base menu state
	return nil
}

// Add Exit method to MainMenuState
func (s *MainMenuState) Exit() error {
	// Implement the logic for exiting the main menu state
	return nil
}

// Add Exit method to BaseMenuState
func (s *BaseMenuState) Exit() error {
	// Implement the logic for exiting the base menu state
	return nil
}
