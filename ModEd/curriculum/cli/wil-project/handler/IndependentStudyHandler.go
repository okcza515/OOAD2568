// MEP-1010 Work Integrated Learning (WIL)
package handler

import (
	"ModEd/asset/util"
	"ModEd/core/cli"
	"ModEd/curriculum/controller"
	// "ModEd/curriculum/model"
	// "ModEd/utils/deserializer"
	"fmt"
)

type IndependentStudyMenuStateHandler struct {
	manager *cli.CLIMenuStateManager
	proxy   *controller.WILModuleProxy

	wilModuleMenuStateHandler *WILModuleMenuStateHandler
}

func NewIndependentStudyMenuStateHandler(
	manager *cli.CLIMenuStateManager, proxy *controller.WILModuleProxy, wilModuleMenuStateHandler *WILModuleMenuStateHandler,
) *IndependentStudyMenuStateHandler {
	return &IndependentStudyMenuStateHandler{
		manager:                   manager,
		proxy:                     proxy,
		wilModuleMenuStateHandler: wilModuleMenuStateHandler,
	}
}

func (menu *IndependentStudyMenuStateHandler) Render() {
	fmt.Println("\nIndependent Study Menu:")
	fmt.Println("1.Migrate independent study")
	fmt.Println("0.Exit Independent Study module")
}

func (menu *IndependentStudyMenuStateHandler) HandleUserInput(input string) error {
		switch input {
		case "1":
			fmt.Println("Migrate Independent Study")
			// path := ""
			// fmt.Println("Please enter the path of the independent study(s) file (csv or json): ")
			// _, _ = fmt.Scanln(&path)
			// fd, err := deserializer.NewFileDeserializer(path)
			// if err != nil {
			// 	fmt.Println(err)
			// 	return nil
			// }
			// isModel := new([]model.IndependentStudy)
			// if err := fd.Deserialize(isModel); err != nil {
			// 	fmt.Println(err)
			// 	return nil
			// }
			// for _, is := range *isModel {
			// 	controller.IndependentStudyController.Insert(&is)
			// }
			// fmt.Println("Migration Success!")
		case "0":
			menu.manager.SetState(menu.wilModuleMenuStateHandler)
			return nil
		default:
			fmt.Println("Invalid Command")
		}

		util.PressEnterToContinue()

		return nil

}