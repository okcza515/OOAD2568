// MEP-1010 Work Integrated Learning (WIL)
package handler

import (
	"ModEd/core"
	"ModEd/core/cli"
	"ModEd/core/handler"
	"ModEd/curriculum/controller"
	"ModEd/curriculum/utils"
	"errors"

	"ModEd/curriculum/model"
	"fmt"
)

type IndependentStudyMenuStateHandler struct {
	manager                   *cli.CLIMenuStateManager
	wrapper                   *controller.WILModuleWrapper
	wilModuleMenuStateHandler *WILModuleMenuStateHandler
	insertHandlerStrategy     *handler.InsertHandlerStrategy[model.IndependentStudy]
	handler                   *utils.GeneralHandlerContext
}

func NewIndependentStudyMenuStateHandler(
	manager *cli.CLIMenuStateManager, wrapper *controller.WILModuleWrapper, wilModuleMenuStateHandler *WILModuleMenuStateHandler,
) *IndependentStudyMenuStateHandler {
	return &IndependentStudyMenuStateHandler{
		manager:                   manager,
		wrapper:                   wrapper,
		wilModuleMenuStateHandler: wilModuleMenuStateHandler,
		insertHandlerStrategy:     handler.NewInsertHandlerStrategy[model.IndependentStudy](wrapper.IndependentStudyController),
		handler:                   utils.NewGeneralHandlerContext(),
	}
}

func (menu *IndependentStudyMenuStateHandler) Render() {
	menu.handler.SetMenuTitle("\nIndependent Study Menu")
	menu.handler.AddHandler("1", "Read independent study list from file", utils.FuncStrategy{
		Action: func() error {
			err := menu.insertHandlerStrategy.Execute()
			return err
		},
	})
	menu.handler.AddHandler("2", "Assign new independent study", utils.FuncStrategy{Action: menu.assignNewIndependentStudy})
	menu.handler.AddHandler("3", "Find IS by Id", utils.FuncStrategy{Action: menu.findISByID})
	menu.handler.AddHandler("4", "List all IS", utils.FuncStrategy{Action: menu.listAllIS})
	menu.handler.AddHandler("5", "Update IS by Id", utils.FuncStrategy{Action: menu.updateIS})
	menu.handler.AddHandler("6", "Delete Independent StudyD", utils.FuncStrategy{Action: menu.deleteIS})
	menu.handler.AddBackHandler(utils.FuncStrategy{
		Action: func() error {
			menu.manager.SetState(menu.wilModuleMenuStateHandler)
			return nil
		},
	})

	menu.handler.ShowMenu()
}

func (menu *IndependentStudyMenuStateHandler) HandleUserInput(input string) error {
	err := menu.handler.HandleInput(input)
	if err != nil {
		return err
	}
	return nil
}

func (menu *IndependentStudyMenuStateHandler) isInformationRenderer(is model.IndependentStudy) {
	fmt.Println("ID                :\t", is.ID)
	fmt.Println("Topic             :\t", is.IndependentStudyTopic)
	fmt.Println("Content           :\t", is.IndependentStudyContent)
	fmt.Println("Assign to group id:\t", is.WILProjectId)
	if time := is.TurnInDate; time == nil {
		fmt.Println("Turn-in date      :\t No turn-in date")
	} else {
		fmt.Println("Turn-in date      :\t", is.TurnInDate)
	}

}

func (menu *IndependentStudyMenuStateHandler) assignNewIndependentStudy() error {
	newIndependentStudy := model.IndependentStudy{}
	fmt.Println("\nAssign new indepndent study topic")
	var turnInDateTime string = ""
	newIndependentStudy.WILProject.ID = core.ExecuteUserInputStep(core.UintInputStep{
		PromptText:    "Enter WIL project group ID you want to assign : ",
		FieldNameText: "WIL project group ID",
	}).(uint)
	newIndependentStudy.IndependentStudyTopic = core.ExecuteUserInputStep(core.StringInputStep{
		PromptText:    "Enter Independent study topic: ",
		FieldNameText: "Independent study topic",
	}).(string)
	newIndependentStudy.IndependentStudyContent = core.ExecuteUserInputStep(core.StringInputStep{
		PromptText:    "Enter Independent study content: ",
		FieldNameText: "Independent study content",
	}).(string)

	fmt.Println("Do you want this IS to has turn-in date?")
	fmt.Println("1. Have turn-in date")
	fmt.Println("2. No turn-in date")
	flag := core.ExecuteUserInputStep(core.UintInputStep{
		PromptText:    "Enter your choice: ",
		FieldNameText: "your choice",
	}).(uint)

	if flag == 1 {
		turnInDate := core.ExecuteUserInputStep(core.StringInputStep{
			PromptText:    "Enter turn-in date [YYYY-mm-dd]: ",
			FieldNameText: "turn-in date",
		}).(string)
		turnTime := core.ExecuteUserInputStep(core.StringInputStep{
			PromptText:    "Enter turn-in time [hh:mm:ss]: ",
			FieldNameText: "turn-in time",
		}).(string)
		turnInDateTime = turnInDate + " " + turnTime
	}
	if err := menu.wrapper.IndependentStudyController.CreateIndependentStudy(&newIndependentStudy, turnInDateTime); err != nil {
		return err
	}
	fmt.Println("\nSuccessfully assign WIL!")
	fmt.Println("\n\t               Summary               ")
	fmt.Println("\t ************************************")
	menu.isInformationRenderer(newIndependentStudy)

	return nil
}

func (menu *IndependentStudyMenuStateHandler) findISByID() error {
	id := core.ExecuteUserInputStep(core.UintInputStep{
		PromptText:    "Enter IS ID: ",
		FieldNameText: "IS ID",
	}).(uint)
	fmt.Println("")
	fmt.Println("\tRequested Independent Study Record")
	is, err := menu.wrapper.IndependentStudyController.BaseController.RetrieveByID(id)
	if err != nil {
		return err
	}
	menu.isInformationRenderer(is)
	return nil
}

func (menu *IndependentStudyMenuStateHandler) listAllIS() error {
	independentStudies, err := menu.wrapper.IndependentStudyController.ListAllIndependentStudy()
	if err != nil {
		return err
	}

	fmt.Println("\n\t List of all Independent Study record")
	fmt.Println("\t ************************************")
	for _, indeindependentStudy := range independentStudies {
		menu.isInformationRenderer(indeindependentStudy)
		fmt.Println("---------------------------------")
	}
	return nil
}

func (menu *IndependentStudyMenuStateHandler) updateIS() error {
	var turnInDateTime string = ""
	id := core.ExecuteUserInputStep(core.UintInputStep{
		PromptText:    "Enter ID of Independent Study you you want to delete: ",
		FieldNameText: "IS ID",
	}).(uint)
	is, err := menu.wrapper.IndependentStudyController.BaseController.RetrieveByID(id)
	if err != nil {
		return errors.New("Failed to retrieved IS")
	}
	fmt.Println("\n\t                      current information")
	fmt.Println("\t **************************************************************")
	menu.isInformationRenderer(is)
	fmt.Println("\n")
	var msg string
	msg = ""

	for msg != "yes" && msg != "no" {
		msg = utils.GetUserInput("Do you want to change Independent Study topic? [yes/no]:")
	}
	if msg == "yes" {
		is.IndependentStudyTopic = core.ExecuteUserInputStep(core.StringInputStep{
			PromptText:    "Enter new IS topic: ",
			FieldNameText: "IndependentStudyTopic",
		}).(string)
	}
	msg = ""
	for msg != "yes" && msg != "no" {
		msg = utils.GetUserInput("Do you want to change Independent Study content? [yes/no]:")
	}
	if msg == "yes" {
		is.IndependentStudyContent = core.ExecuteUserInputStep(core.StringInputStep{
			PromptText:    "Enter new IS content: ",
			FieldNameText: "IndependentStudyTopic",
		}).(string)
	}
	msg = ""
	for msg != "yes" && msg != "no" {
		msg = utils.GetUserInput("Do you want to modify Turn-in date? [yes/no]:")
	}
	if msg == "yes" {
		fmt.Println("Do you want this IS to has turn-in date?")
		fmt.Println("1. Have turn-in date")
		fmt.Println("2. No turn-in date")
		var flag uint = 0
		for flag != 1 && flag != 2 {
			flag = core.ExecuteUserInputStep(core.UintInputStep{
				PromptText:    "Enter your choice: ",
				FieldNameText: "your choice",
			}).(uint)
		}
		if flag == 1 {
			turnInDate := core.ExecuteUserInputStep(core.StringInputStep{
				PromptText:    "Enter turn-in date [YYYY-mm-dd]: ",
				FieldNameText: "turn-in date",
			}).(string)
			turnTime := core.ExecuteUserInputStep(core.StringInputStep{
				PromptText:    "Enter turn-in time [hh:mm:ss]: ",
				FieldNameText: "turn-in time",
			}).(string)
			turnInDateTime = turnInDate + " " + turnTime
		}
		if err := menu.wrapper.IndependentStudyController.UpdateIndependentStudy(is, turnInDateTime); err != nil {
			return err
		}
		fmt.Println("\nSuccessfully updated!")
		fmt.Println("\n\t                      new information")
		fmt.Println("\t **************************************************************")
		is, err := menu.wrapper.IndependentStudyController.BaseController.RetrieveByID(id)
		if err != nil {
			return errors.New("Failed to retrieved IS")
		}
		menu.isInformationRenderer(is)
		fmt.Println("\n")
	}

	return nil
}

func (menu *IndependentStudyMenuStateHandler) deleteIS() error {
	id := core.ExecuteUserInputStep(core.UintInputStep{
		PromptText:    "Enter ID of Independent Study you you want to delete: ",
		FieldNameText: "IS ID",
	}).(uint)
	is, err := menu.wrapper.IndependentStudyController.BaseController.RetrieveByID(id)
	if err != nil {
		return errors.New("Failed to retrieved IS")
	}
	fmt.Println("\n\t The independent with the following information will be deleted")
	fmt.Println("\t **************************************************************")
	menu.isInformationRenderer(is)
	var msg string
	for msg != "yes" && msg != "no" {
		msg = utils.GetUserInput("Delete this record? [yes/no]:")
	}
	if msg == "yes" {
		if err := menu.wrapper.IndependentStudyController.DeleteIndependentStudiesByID(id); err != nil {
			return err
		}
		return nil
	}

	fmt.Println("Exiting delete operation cancled...")
	return nil
}
