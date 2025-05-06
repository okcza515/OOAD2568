// MEP-1010 Work Integrated Learning (WIL)
package handler

import (
	"ModEd/core"
	"ModEd/core/cli"
	"ModEd/core/validation"
	"ModEd/curriculum/controller"
	"ModEd/curriculum/model"
	"ModEd/curriculum/utils"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

type WILProjectMenuStateHandler struct {
	manager                   *cli.CLIMenuStateManager
	wrapper                   *controller.WILModuleWrapper
	wilModuleMenuStateHandler *WILModuleMenuStateHandler
	handler                   *utils.GeneralHandlerContext
}

func NewWILProjectMenuStateHandler(
	manager *cli.CLIMenuStateManager, wrapper *controller.WILModuleWrapper, wilModuleMenuStateHandler *WILModuleMenuStateHandler,
) *WILProjectMenuStateHandler {
	return &WILProjectMenuStateHandler{
		manager:                   manager,
		wrapper:                   wrapper,
		wilModuleMenuStateHandler: wilModuleMenuStateHandler,
		handler:                   utils.NewGeneralHandlerContext(),
	}
}

func (menu *WILProjectMenuStateHandler) Render() {
	menu.handler.SetMenuTitle("\nWIL Project Menu:")
	menu.handler.AddHandler("1", "Create WIL Project", utils.FuncStrategy{Action: menu.createCreateWILProject})
	menu.handler.AddHandler("2", "Edit WIL Project", utils.FuncStrategy{Action: menu.editWILProject})
	menu.handler.AddHandler("3", "Search WIL Project", nil)
	menu.handler.AddHandler("4", "List all WIL Project", utils.FuncStrategy{Action: menu.listAllWILProject})
	menu.handler.AddHandler("5", "Get WIL Project Detail By ID", utils.FuncStrategy{Action: menu.getWILProjectDetailByID})
	menu.handler.AddHandler("6", "Delete WIL Project By ID", utils.FuncStrategy{Action: menu.deleteWILProjectByID})
	menu.handler.AddBackHandler(utils.FuncStrategy{Action: func() error {
		menu.manager.SetState(menu.wilModuleMenuStateHandler)
		return nil
	}})

	menu.handler.ShowMenu()
}

func (menu *WILProjectMenuStateHandler) HandleUserInput(input string) error {
	err := menu.handler.HandleInput(input)
	if err != nil {
		return err
	}
	return nil
}

func (menu *WILProjectMenuStateHandler) createCreateWILProject() error {
	classId := core.ExecuteUserInputStep(core.UintInputStep{
		PromptText:    "Enter class Id:",
		FieldNameText: "ClassId",
	}).(uint)
	seniorProjectId := core.ExecuteUserInputStep(core.UintInputStep{
		PromptText:    "Enter Senior Project Id:",
		FieldNameText: "SeniorProjectId",
	}).(uint)
	companyId := core.ExecuteUserInputStep(core.UintInputStep{
		PromptText:    "Enter company Id:",
		FieldNameText: "CompanyId",
	}).(uint)
	mentor := core.ExecuteUserInputStep(core.StringInputStep{
		PromptText:    "Enter Mentor:",
		FieldNameText: "Mentor",
	}).(string)

	WILProject := model.WILProject{
		ClassId:         classId,
		SeniorProjectId: seniorProjectId,
		Company:         companyId,
		Mentor:          mentor,
	}

	validator := validation.NewModelValidator()
	err := validator.ModelValidate(WILProject)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	err = menu.wrapper.WILProjectController.Insert(WILProject)
	if err != nil {
		errMsg := errors.New("error! cannot create WIL Project: " + err.Error())
		return errMsg
	} else {
		fmt.Printf("WIL Project created successfully")
	}
	return nil
}

func (menu *WILProjectMenuStateHandler) editWILProject() error {
	WILProjectID := core.ExecuteUserInputStep(core.UintInputStep{
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
		temp := core.ExecuteUserInputStep(core.UintInputStep{
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
		temp := core.ExecuteUserInputStep(core.UintInputStep{
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
		temp := core.ExecuteUserInputStep(core.UintInputStep{
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
		temp := core.ExecuteUserInputStep(core.StringInputStep{
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
	WILProjectID := core.ExecuteUserInputStep(core.UintInputStep{
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
	WILProjectID := core.ExecuteUserInputStep(core.UintInputStep{
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
