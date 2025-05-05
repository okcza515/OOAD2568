package cli

import (
	"ModEd/core/cli"
	"ModEd/recruit/util"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type AdminDeleteInterviewMenuState struct {
	manager          *cli.CLIMenuStateManager
	interviewService AdminInterviewService
	parent           cli.MenuState
}

func NewAdminDeleteInterviewMenuState(
	manager *cli.CLIMenuStateManager,
	interviewService AdminInterviewService,
	parent cli.MenuState,
) *AdminDeleteInterviewMenuState {
	return &AdminDeleteInterviewMenuState{
		manager:          manager,
		interviewService: interviewService,
		parent:           parent,
	}
}

func (menu *AdminDeleteInterviewMenuState) Render() {
	util.ClearScreen()
	fmt.Print("Enter Interview ID to delete: ")
}

func (menu *AdminDeleteInterviewMenuState) HandleUserInput(input string) error {
	input = strings.TrimSpace(input)
	interviewID, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println("Invalid Interview ID.")
	} else {
		err := menu.interviewService.DeleteInterview(uint(interviewID))
		if err != nil {
			fmt.Println("Failed to delete interview:", err)
		} else {
			fmt.Println("Interview deleted successfully.")
		}
	}

	fmt.Println("\nPress Enter to return...")
	bufio.NewReader(os.Stdin).ReadBytes('\n')
	menu.manager.SetState(menu.parent)
	return nil
}
