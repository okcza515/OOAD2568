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
		if err := menu.deleteWILProjectByID(); err != nil {
			fmt.Println("error! cannot use this function")
		}
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
	classId := utils.ExecuteUserInputStep(utils.UintInputStep{
		PromptText:    "Enter class Id:",
		FieldNameText: "ClassId",
	}).(uint)
	seniorProjectId := utils.ExecuteUserInputStep(utils.UintInputStep{
		PromptText:    "Enter Senior Project Id:",
		FieldNameText: "SeniorProjectId",
	}).(uint)
	companyId := utils.ExecuteUserInputStep(utils.UintInputStep{
		PromptText:    "Enter company Id:",
		FieldNameText: "CompanyId",
	}).(uint)
	mentor := utils.ExecuteUserInputStep(utils.StringInputStep{
		PromptText:    "Enter Mentor:",
		FieldNameText: "Mentor",
	}).(string)

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
	WILProjectID := utils.ExecuteUserInputStep(utils.UintInputStep{
		PromptText:    "Enter WIL Project Id:",
		FieldNameText: "ClassId",
	}).(uint)

	WILProject, err := menu.wrapper.WILProjectController.RetrieveByID(WILProjectID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			fmt.Println("WIL Project Not Found")
			return nil
		} else {
			return err
		}
	}

	NewWILProject := &model.WILProject{}
	NewWILProject.ID = WILProjectID

	var msg string
	msg = ""
	for msg != "yes" && msg != "y" && msg != "no" && msg != "n" {
		msg = utils.GetUserInput(fmt.Sprintf("\nClassId : %d | Want to change ClassId [yes/no]: ", WILProject.ClassId))

	}
	if msg == "yes" || msg == "y" {
		temp := utils.ExecuteUserInputStep(utils.UintInputStep{
			PromptText:    "New Class Id: ",
			FieldNameText: "ClassId",
		}).(uint)
		NewWILProject.ClassId = temp
	}

	msg = ""
	for msg != "yes" && msg != "y" && msg != "no" && msg != "n" {
		msg = utils.GetUserInput(fmt.Sprintf("\nSeniorProjectId : %d | Want to change SeniorProjectId [yes/no]: ", WILProject.SeniorProjectId))
	}
	if msg == "yes" || msg == "y" {
		temp := utils.ExecuteUserInputStep(utils.UintInputStep{
			PromptText:    "New Senior Project Id: ",
			FieldNameText: "SeniorProjectId",
		}).(uint)
		NewWILProject.SeniorProjectId = temp
	}

	msg = ""
	for msg != "yes" && msg != "y" && msg != "no" && msg != "n" {
		msg = utils.GetUserInput(fmt.Sprintf("\nCompany : %d | Want to change Company [yes/no]: ", WILProject.Company))
	}
	if msg == "yes" || msg == "y" {
		temp := utils.ExecuteUserInputStep(utils.UintInputStep{
			PromptText:    "New Company: ",
			FieldNameText: "Company",
		}).(uint)
		NewWILProject.Company = temp
	}

	msg = ""
	for msg != "yes" && msg != "y" && msg != "no" && msg != "n" {
		msg = utils.GetUserInput(fmt.Sprintf("\nMentor : %s | Want to change Mentor [yes/no]: ", WILProject.Mentor))
	}
	if msg == "yes" || msg == "y" {
		temp := utils.ExecuteUserInputStep(utils.StringInputStep{
			PromptText:    "New Mentor: ",
			FieldNameText: "Mentor",
		}).(string)
		NewWILProject.Mentor = temp
	}

	if err := menu.wrapper.WILProjectController.UpdateByID(*NewWILProject); err != nil {
		fmt.Println(err.Error())
		return err
	}
	fmt.Println("Update Success")

	return nil
}

func (menu *WILProjectMenuStateHandler) listAllWILProject() error {
	fmt.Printf("\nWIL Project List\n\n")
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
	var WILProject model.WILProject
	var err error
	WILProjectID := utils.ExecuteUserInputStep(utils.UintInputStep{
		PromptText:    "Enter WIL Project Id:",
		FieldNameText: "WILProjectId",
	}).(uint)

	if WILProject, err = menu.wrapper.WILProjectController.RetrieveByID(WILProjectID); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			fmt.Println("WIL Project Not Found")
			return nil
		} else {
			return err
		}
	}

	fmt.Printf("ID: %d, Class ID: %d, Senior Project ID: %d, Company ID: %d, Mentor: %s\n",
		WILProject.ID, WILProject.ClassId, WILProject.SeniorProjectId, WILProject.Company, WILProject.Mentor)
	return nil
}

func (menu *WILProjectMenuStateHandler) deleteWILProjectByID() error {
	WILProjectID := utils.ExecuteUserInputStep(utils.UintInputStep{
		PromptText:    "Enter WIL Project Id:",
		FieldNameText: "WILProjectId",
	}).(uint)

	if _, err := menu.wrapper.WILProjectController.RetrieveByID(WILProjectID); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			fmt.Println("WIL Project Not Found")
			return nil
		} else {
			return err
		}
	}

	if err := menu.wrapper.WILProjectController.DeleteByID(WILProjectID); err != nil {
		return err
	}
	fmt.Printf("Delete WIL Project ID: %d Success\n", WILProjectID)
	return nil
}
