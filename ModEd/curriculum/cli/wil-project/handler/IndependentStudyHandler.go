// MEP-1010 Work Integrated Learning (WIL)
package handler

import (
	"ModEd/asset/util"
	"ModEd/core/cli"
	"ModEd/core/handler"
	"ModEd/curriculum/controller"
	"ModEd/curriculum/utils"

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
	fmt.Println("back: Exit the module")
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
			fmt.Print("Assign failed exiting with error [")
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

func (menu *IndependentStudyMenuStateHandler) wilInformationRenderer(id uint) error {
	is, err := menu.wrapper.IndependentStudyController.BaseController.RetrieveByID(id)
	if err != nil {
		return err
	}
	fmt.Println("Topic             :\t", is.IndependentStudyTopic)
	fmt.Println("Content           :\t", is.IndependentStudyContent)
	fmt.Println("Assign to group id:\t", is.WILProject.ID)
	if time := is.TurnInDate; time == nil {
		fmt.Println("Turn-in date      :\t No turn-in date")
	} else {
		fmt.Println("Turn-in date      :\t", is.TurnInDate)
	}
	return nil
}

//func (menu *IndependentStudyMenuStateHandler) readIndependentStudyFromFile() error {
//	fmt.Println("")
//	fmt.Println("Read IS list from file")
//	path := ""
//	fmt.Println("Please enter the path of the independent study(s) file (csv or json): ")
//	_, _ = fmt.Scanln(&path)
//	fd, err := deserializer.NewFileDeserializer(path)
//	if err != nil {
//		return err
//	}
//	isModel := new([]model.IndependentStudy)
//	if err := fd.Deserialize(isModel); err != nil {
//		return err
//	}
//
//	if err := menu.wrapper.IndependentStudyController.BaseController.InsertMany(*isModel); err != nil {
//		return err
//	}
//	fmt.Println("\nRead file Success!")
//	return nil
//}

func (menu *IndependentStudyMenuStateHandler) assignNewIndependentStudy() error {
	newIndependentStudy := model.IndependentStudy{}
	fmt.Println("\nAssign new indepndent study topic")
	var turnInDateTime string = ""
	newIndependentStudy.WILProject.ID = utils.ExecuteUserInputStep(utils.UintInputStep{
		PromptText:    "Enter WIL project group ID you want to assign : ",
		FieldNameText: "WIL project group ID",
	}).(uint)
	newIndependentStudy.IndependentStudyTopic = utils.ExecuteUserInputStep(utils.StringInputStep{
		PromptText:    "Enter Independent study topic: ",
		FieldNameText: "Independent study topic",
	}).(string)
	newIndependentStudy.IndependentStudyContent = utils.ExecuteUserInputStep(utils.StringInputStep{
		PromptText:    "Enter Independent study content: ",
		FieldNameText: "Independent study content",
	}).(string)

	fmt.Println("Do you want this IS to has turn-in date?")
	fmt.Println("1. Have turn-in date")
	fmt.Println("2. No turn-in date")
	flag := utils.ExecuteUserInputStep(utils.UintInputStep{
		PromptText:    "Enter your choice: ",
		FieldNameText: "your choice",
	}).(uint)

	if flag == 1 {
		turnInDate := utils.ExecuteUserInputStep(utils.StringInputStep{
			PromptText:    "Enter turn-in date [YYYY-mm-dd]: ",
			FieldNameText: "turn-in date",
		}).(string)
		turnTime := utils.ExecuteUserInputStep(utils.StringInputStep{
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
	id := utils.ExecuteUserInputStep(utils.UintInputStep{
		PromptText:    "Enter IS ID: ",
		FieldNameText: "IS ID",
	}).(uint)
	fmt.Println("")
	if err := menu.wilInformationRenderer(id); err != nil {
		return err
	}
	return nil
}
