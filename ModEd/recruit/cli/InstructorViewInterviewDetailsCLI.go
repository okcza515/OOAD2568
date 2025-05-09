// MEP-1003 Student Recruitment
package cli

import (
	"ModEd/core/cli"
	"ModEd/recruit/controller"
	"ModEd/recruit/util"
	"fmt"
)

type InstructorViewInterviewDetailsMenuState struct {
	manager      *cli.CLIMenuStateManager
	instructorID uint
	filter       string
	service      InstructorViewInterviewDetailsService
	previous     cli.MenuState
	controller   *controller.InterviewController
}

func NewInstructorViewInterviewDetailsMenuState(
	manager *cli.CLIMenuStateManager,
	instructorID uint,
	filter string,
	service InstructorViewInterviewDetailsService,
	previous cli.MenuState,
	controller *controller.InterviewController,
) *InstructorViewInterviewDetailsMenuState {
	return &InstructorViewInterviewDetailsMenuState{
		manager:      manager,
		instructorID: instructorID,
		filter:       filter,
		service:      service,
		previous:     previous,
		controller:   controller,
	}
}

func (m *InstructorViewInterviewDetailsMenuState) Render() {
	util.ClearScreen()

	interviews, err := m.service.ViewInterviewDetails(m.instructorID, m.filter)
	if err != nil {
		fmt.Println("Error fetching interview details:", err)
		util.WaitForEnter()
		m.manager.SetState(m.previous)
		return
	}

	m.service.DisplayReport(interviews)

	fmt.Println("\nPress Enter to go back...")
}

func (m *InstructorViewInterviewDetailsMenuState) HandleUserInput(_ string) error {
	m.manager.SetState(m.previous)
	return nil
}
