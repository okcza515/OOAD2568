// MEP-1010 Work Integrated Learning (WIL)
package handler

import (
	"ModEd/asset/util"
	"ModEd/core"
	"ModEd/core/cli"
	"ModEd/core/handler"
	"ModEd/curriculum/controller"

	"ModEd/curriculum/model"
	"fmt"
)

type IndependentStudyMenuStateHandler struct {
	manager *cli.CLIMenuStateManager
	wrapper *controller.WILModuleWrapper

	wilModuleMenuStateHandler *WILModuleMenuStateHandler
	insertHandlerStrategy     *handler.InsertHandlerStrategy[model.IndependentStudy]
}

func NewIndependentStudyMenuStateHandler(
	manager *cli.CLIMenuStateManager, wrapper *controller.WILModuleWrapper, wilModuleMenuStateHandler *WILModuleMenuStateHandler,
) *IndependentStudyMenuStateHandler {
	return &IndependentStudyMenuStateHandler{
		manager:                   manager,
		wrapper:                   wrapper,
		wilModuleMenuStateHandler: wilModuleMenuStateHandler,
		insertHandlerStrategy:     handler.NewInsertHandlerStrategy[model.IndependentStudy](wrapper.IndependentStudyController),
	}
}

func (menu *IndependentStudyMenuStateHandler) Render() {
	fmt.Println("\nIndependent Study Menu:")
	fmt.Println("1.Read independent study list from file")
	fmt.Println("2.Assign new independent study")
	fmt.Println("3.Find IS by Id")
	fmt.Println("4.Update IS by Id")
	fmt.Println("5.Delete Independent Study")
	fmt.Println("back: Exit Independent Study module")
}

func (menu *IndependentStudyMenuStateHandler) HandleUserInput(input string) error {
	switch input {
	case "1":
		err := menu.insertHandlerStrategy.Execute()
		if err != nil {
			fmt.Println("error! cannot use this function")
		}
	case "2":
		if err := menu.assignNewIndependentStudy(); err != nil {
			fmt.Print("Assign failed exiting with error [")
			fmt.Print(err)
			fmt.Println("]")
		}
	case "3":
		if err := menu.findISByID(); err != nil {
			fmt.Print("Retrive failed exiting with error [")
			fmt.Print(err)
			fmt.Println("]")
		}
	case "4":
		fmt.Println("Not ready")

	case "5":
		if err := menu.deleteIS(); err != nil {
			fmt.Print("Delete failed exiting with error [")
			fmt.Print(err)
			fmt.Println("]")
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

func (menu *IndependentStudyMenuStateHandler) isInformationRenderer(id uint) error {
	is, err := menu.wrapper.IndependentStudyController.BaseController.RetrieveByID(id)
	if err != nil {
		return err
	}
	fmt.Println("Topic             :\t", is.IndependentStudyTopic)
	fmt.Println("Content           :\t", is.IndependentStudyContent)
	fmt.Println("Assign to group id:\t", is.WILProjectId)
	if time := is.TurnInDate; time == nil {
		fmt.Println("Turn-in date      :\t No turn-in date")
	} else {
		fmt.Println("Turn-in date      :\t", is.TurnInDate)
	}
	return nil
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
	return nil
}

func (menu *IndependentStudyMenuStateHandler) findISByID() error {
	id := core.ExecuteUserInputStep(core.UintInputStep{
		PromptText:    "Enter IS ID: ",
		FieldNameText: "IS ID",
	}).(uint)
	fmt.Println("")
	fmt.Println("\tRequested Independent Study Record")
	if err := menu.isInformationRenderer(id); err != nil {
		return err
	}
	return nil
}

func (menu *IndependentStudyMenuStateHandler) deleteIS() error {
	id := core.ExecuteUserInputStep(core.UintInputStep{
		PromptText:    "Enter ID of Independent Study you you want to delete: ",
		FieldNameText: "IS ID",
	}).(uint)
	if err := menu.wrapper.IndependentStudyController.BaseController.DeleteByID(id); err != nil {
		return err
	}
	return nil
}
