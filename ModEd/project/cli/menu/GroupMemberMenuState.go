package menu

import (
	"ModEd/core"
	"ModEd/core/cli"
	"ModEd/core/handler"
	"ModEd/project/cli_refactor_prototype_v3/handlers"
	"ModEd/project/controller"
	"ModEd/project/utils"
	"fmt"
)

type GroupMemberMenuState struct {
	manager        *cli.CLIMenuStateManager
	handlerContext *handler.HandlerContext
}

func NewGroupMemberMenuState(manager *cli.CLIMenuStateManager, storer *controller.InstanceStorer) *GroupMemberMenuState {
	handlerContext := handler.NewHandlerContext()
	handlerContext.SetMenuTitle("Group Member Management")

	io := core.NewMenuIO()
	h := handlers.NewGroupMemberHandler(storer)

	handlerContext.AddHandler("1", "View All Group Members", handler.FuncStrategy{
		Action: func() error {
			h.ViewAll(io)
			return nil
		},
	})

	handlerContext.AddHandler("2", "Add Group Member", handler.FuncStrategy{
		Action: func() error {
			h.Add(io)
			return nil
		},
	})

	handlerContext.AddHandler("3", "Update Group Member", handler.FuncStrategy{
		Action: func() error {
			h.Update(io)
			return nil
		},
	})

	handlerContext.AddHandler("4", "Delete Group Member", handler.FuncStrategy{
		Action: func() error {
			h.Delete(io)
			return nil
		},
	})

	backHandler := handler.NewChangeMenuHandlerStrategy(manager, manager.GetState("MAIN"))
	handlerContext.AddBackHandler(backHandler)

	return &GroupMemberMenuState{
		manager:        manager,
		handlerContext: handlerContext,
	}
}

func (menu *GroupMemberMenuState) Render() {
	fmt.Println()
	fmt.Println("::/project/group-member")
	fmt.Println()
	fmt.Println("Group Member Management Menu")
	menu.handlerContext.ShowMenu()
	fmt.Println("  exit:\tExit the program (or Ctrl+C is fine ¯\\_(ツ)_/¯)")
	fmt.Println()
}

func (menu *GroupMemberMenuState) HandleUserInput(input string) error {
	err := menu.handlerContext.HandleInput(input)
	if err != nil {
		fmt.Println("Error:", err)
	}

	if input != "back" {
		utils.PressEnterToContinue()
	}

	return nil
}
