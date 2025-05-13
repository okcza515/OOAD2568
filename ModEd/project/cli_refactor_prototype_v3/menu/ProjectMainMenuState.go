package menu

import (
	"ModEd/core/cli"
	"ModEd/core/handler"
	"ModEd/project/controller"
	"ModEd/project/utils"
)

type ProjectMainMenuState struct {
	manager        *cli.CLIMenuStateManager
	handlerContext *handler.HandlerContext
}

func NewProjectMainMenuState(manager *cli.CLIMenuStateManager, storer *controller.InstanceStorer) *ProjectMainMenuState {
	handlerContext := handler.NewHandlerContext()
	handlerContext.SetMenuTitle("Project Main Menu")

	menustate := &ProjectMainMenuState{manager: manager, handlerContext: handlerContext}
	manager.AddMenu("MAIN", menustate)

	handlerContext.AddHandler("1", "Senior Project Management", handler.NewChangeMenuHandlerStrategy(manager, NewSeniorProjectMenuState(manager, storer)))
	handlerContext.AddHandler("2", "Advisor Management", handler.NewChangeMenuHandlerStrategy(manager, NewAdvisorMenuState(manager, storer)))
	handlerContext.AddHandler("3", "Committee Management", handler.NewChangeMenuHandlerStrategy(manager, NewCommitteeMenuState(manager, storer)))
	handlerContext.AddHandler("4", "Group Member Management", handler.NewChangeMenuHandlerStrategy(manager, NewGroupMemberMenuState(manager, storer)))
	handlerContext.AddHandler("5", "Assignment Management", handler.NewChangeMenuHandlerStrategy(manager, NewAssignmentMenuState(manager, storer)))
	handlerContext.AddHandler("6", "Presentation Management", handler.NewChangeMenuHandlerStrategy(manager, NewPresentationMenuState(manager, storer)))
	handlerContext.AddHandler("7", "Report Management", handler.NewChangeMenuHandlerStrategy(manager, NewReportMenuState(manager, storer)))
	handlerContext.AddHandler("8", "Progress Tracking", handler.NewChangeMenuHandlerStrategy(manager, NewProgressMenuState(manager, storer)))
	handlerContext.AddHandler("9", "Assessment Scores", handler.NewChangeMenuHandlerStrategy(manager, NewAssessmentCriteriaLinkMenuState(manager, storer)))
	handlerContext.AddHandler("10", "Assignment Scores", handler.NewChangeMenuHandlerStrategy(manager, NewAssignmentScoreMenuState(manager, storer)))
	handlerContext.AddHandler("11", "Presentation Scores", handler.NewChangeMenuHandlerStrategy(manager, NewPresentationScoreMenuState(manager, storer)))
	handlerContext.AddHandler("12", "Report Scores", handler.NewChangeMenuHandlerStrategy(manager, NewReportScoreMenuState(manager, storer)))

	return menustate
}

func (menu *ProjectMainMenuState) Render() {
	utils.MenuTitle("Senior Project CLI")
	menu.handlerContext.ShowMenu()
}

func (menu *ProjectMainMenuState) HandleUserInput(input string) error {
	return menu.handlerContext.HandleInput(input)
}
