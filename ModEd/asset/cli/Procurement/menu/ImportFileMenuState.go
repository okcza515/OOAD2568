// MEP-1014
package menu

import (
	"ModEd/asset/cli/Procurement/helper"
	"ModEd/asset/controller"
	"ModEd/asset/util"
	"ModEd/core/cli"
	"ModEd/core/handler"
	"fmt"
)

type DepartmentMenuState struct {
	manager        *cli.CLIMenuStateManager
	handlerContext *handler.HandlerContext
}

func NewImportFileMenuState(manager *cli.CLIMenuStateManager) *DepartmentMenuState {
	facade, err := controller.CreateProcurementControllerFacade()
	if err != nil {
		fmt.Println("Failed to create ProcurementControllerFacade:", err)
		return nil
	}

	handlerContext := handler.NewHandlerContext()
	menu := &DepartmentMenuState{
		manager:        manager,
		handlerContext: handlerContext,
	}

	handlerContext.AddHandler("1", "Import Departments from CSV", handler.FuncStrategy{
		Action: func() error {
			path := util.GetStringInput("Insert your Department.csv path: ")
			err := helper.SeedDepartmentsFromCSV(facade.GetDB(), path)
			if err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println("Import successful.")
			}
			util.PressEnterToContinue()
			return nil
		},
	})

	handlerContext.AddBackHandler(handler.NewChangeMenuHandlerStrategy(manager, manager.GetState(string(MENU_PROCUREMENT_MAIN))))

	return menu
}

func (menu *DepartmentMenuState) Render() {
	fmt.Println()
	fmt.Println(":/procurement/department")
	fmt.Println()
	fmt.Println("Department Management Menu:")
	menu.handlerContext.ShowMenu()
	fmt.Println()
}

func (menu *DepartmentMenuState) HandleUserInput(input string) error {
	return menu.handlerContext.HandleInput(input)
}
