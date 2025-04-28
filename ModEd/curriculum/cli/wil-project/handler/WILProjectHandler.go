// MEP-1010 Work Integrated Learning (WIL)
package handler

import (
	"ModEd/asset/util"
	"ModEd/core/cli"
	"ModEd/curriculum/controller"
	"ModEd/curriculum/model"
	"ModEd/curriculum/utils"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

type WILProjectMenuStateHandler struct {
	manager *cli.CLIMenuStateManager
	wrapper *controller.WILModuleWrapper

	wilModuleMenuStateHandler *WILModuleMenuStateHandler
}

func NewWILProjectMenuStateHandler(
	manager *cli.CLIMenuStateManager, wrapper *controller.WILModuleWrapper, wilModuleMenuStateHandler *WILModuleMenuStateHandler,
) *WILProjectMenuStateHandler {
	return &WILProjectMenuStateHandler{
		manager:                   manager,
		wrapper:                   wrapper,
		wilModuleMenuStateHandler: wilModuleMenuStateHandler,
	}
}

func (menu *WILProjectMenuStateHandler) Render() {
	fmt.Println("\nWIL Project Menu:")
	fmt.Println("1. Create WIL Project")
	fmt.Println("2. Edit WIL Project")
	fmt.Println("3. Search WIL Project")
	fmt.Println("4. List all WIL Project")
	fmt.Println("5. Get WIL Project Detail By ID")
	fmt.Println("6. Delete WIL Project By ID")
	fmt.Println("back: Exit the module")
}

func (menu *WILProjectMenuStateHandler) HandleUserInput(input string) error {
	switch input {
	case "1":
		if err := menu.createCreateWILProject(); err != nil {
			fmt.Println("error! cannot use this function")
		}
	case "2":
		if err := menu.editWILProject(); err != nil {
			fmt.Println("error! cannot use this function")
		}
	case "3":
		fmt.Println("3 Not implemented yet...")
	case "4":
		if err := menu.listAllWILProject(); err != nil {
			fmt.Println("error! cannot use this function")
		}
	case "5":
		if err := menu.getWILProjectDetailByID(); err != nil {
			fmt.Println("error! cannot use this function")
		}
	case "6":
		fmt.Println("6 Not implemented yet...")
	case "back":
		menu.manager.SetState(menu.wilModuleMenuStateHandler)
		return nil
	default:
		fmt.Println("Invalid Command")
	}

	util.PressEnterToContinue()

	return nil
}

func (menu *WILProjectMenuStateHandler) createCreateWILProject() error {
	classId := utils.GetUserInputUint("Enter class Id:")
	seniorProjectId := utils.GetUserInputUint("Enter Senior Project Id:")
	companyId := utils.GetUserInputUint("Enter company Id:")
	mentor := utils.GetUserInput("Enter Mentor:")

	WILProject := model.WILProject{
		ClassId:         classId,
		SeniorProjectId: seniorProjectId,
		Company:         companyId,
		Mentor:          mentor,
	}

	err := menu.wrapper.WILProjectController.Insert(WILProject)
	if err != nil {
		return errors.New("error! cannot create WIL Project: " + err.Error())
	} else {
		fmt.Printf("WIL Project created successfully")
	}
	return nil
}

func (menu *WILProjectMenuStateHandler) editWILProject() error {
	return nil
}

func (menu *WILProjectMenuStateHandler) listAllWILProject() error {
	fmt.Println("\nWIL Project List\n")
	WILProjects, err := menu.wrapper.WILProjectController.RetrieveAllWILProjects()
	if err != nil {
		return err
	}

	for _, project := range WILProjects {
		fmt.Printf("ID: %d, Class ID: %d, Senior Project ID: %d, Company ID: %d, Mentor: %s\n",
			project.ID, project.ClassId, project.SeniorProjectId, project.Company, project.Mentor)
	}

	return nil
}

func (menu *WILProjectMenuStateHandler) getWILProjectDetailByID() error {
	WILProjectID := utils.GetUserInputUint("Enter WIL Project Id:")

	WILProject, err := menu.wrapper.WILProjectController.RetrieveByID(WILProjectID)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		fmt.Println("WIL Project Not Found")
	} else if err != nil {
		return err
	} else {
		fmt.Printf("ID: %d, Class ID: %d, Senior Project ID: %d, Company ID: %d, Mentor: %s\n",
			WILProject.ID, WILProject.ClassId, WILProject.SeniorProjectId, WILProject.Company, WILProject.Mentor)
	}

	return nil
}
