package cli

import (
	"ModEd/core/cli"
	"ModEd/recruit/util"
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type AdminScheduleInterviewMenuState struct {
	manager *cli.CLIMenuStateManager
	service AdminScheduleInterviewService
	parent  cli.MenuState
}

func NewAdminScheduleInterviewMenuState(
	manager *cli.CLIMenuStateManager,
	service AdminScheduleInterviewService,
	parent cli.MenuState,
) *AdminScheduleInterviewMenuState {
	return &AdminScheduleInterviewMenuState{
		manager: manager,
		service: service,
		parent:  parent,
	}
}

func (menu *AdminScheduleInterviewMenuState) Render() {
	util.ClearScreen()
	fmt.Println("=== Schedule Interview ===")
	fmt.Print("Enter Instructor ID: ")
}

func (menu *AdminScheduleInterviewMenuState) HandleUserInput(input string) error {
	scanner := bufio.NewScanner(os.Stdin)

	instructorID, err := strconv.ParseUint(input, 10, 32)
	if err != nil {
		fmt.Println("Invalid Instructor ID.")
		util.WaitForEnter()
		menu.manager.SetState(menu.parent)
		return nil
	}

	fmt.Print("Enter Application Report ID: ")
	scanner.Scan()
	reportInput := scanner.Text()
	reportID, err := strconv.ParseUint(reportInput, 10, 32)
	if err != nil {
		fmt.Println("Invalid Application Report ID.")
		util.WaitForEnter()
		menu.manager.SetState(menu.parent)
		return nil
	}

	fmt.Print("Enter Scheduled Appointment (YYYY-MM-DD HH:MM:SS): ")
	scanner.Scan()
	scheduledTime := scanner.Text()

	err = menu.service.ScheduleInterview(uint(instructorID), uint(reportID), scheduledTime)
	if err != nil {
		fmt.Println("Error scheduling interview:", err)
	} else {
		fmt.Println("Interview scheduled successfully!")
	}

	util.WaitForEnter()
	menu.manager.SetState(menu.parent)
	return nil
}
