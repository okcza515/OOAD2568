// MEP-1010 Work Integrated Learning (WIL)
package handler

import (
	"ModEd/asset/util"
	"ModEd/core/cli"
	"ModEd/curriculum/controller"
	"ModEd/curriculum/utils"

	"ModEd/curriculum/model"
	"ModEd/utils/deserializer"
	"fmt"
)

type IndependentStudyMenuStateHandler struct {
	manager *cli.CLIMenuStateManager
	wrapper *controller.WILModuleWrapper

	wilModuleMenuStateHandler *WILModuleMenuStateHandler
}

func NewIndependentStudyMenuStateHandler(
	manager *cli.CLIMenuStateManager, wrapper *controller.WILModuleWrapper, wilModuleMenuStateHandler *WILModuleMenuStateHandler,
) *IndependentStudyMenuStateHandler {
	return &IndependentStudyMenuStateHandler{
		manager:                   manager,
		wrapper:                   wrapper,
		wilModuleMenuStateHandler: wilModuleMenuStateHandler,
	}
}

func (menu *IndependentStudyMenuStateHandler) Render() {
	fmt.Println("\nIndependent Study Menu:")
	fmt.Println("1.Read independent study list from file")
	fmt.Println("2.Assign new independent study")
	fmt.Println("3.Find IS by Id")
	fmt.Println("0.Exit Independent Study module")
}

func (menu *IndependentStudyMenuStateHandler) HandleUserInput(input string) error {
	switch input {
	case "1":
		if err := menu.readIndependentStudyFromFile(); err != nil {
			fmt.Print("Read failed exiting with error [")
			fmt.Print(err)
			fmt.Println("]")
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
	case "0":
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

func (menu *IndependentStudyMenuStateHandler) readIndependentStudyFromFile() error {
	fmt.Println("")
	fmt.Println("Read IS list from file")
	path := ""
	fmt.Println("Please enter the path of the independent study(s) file (csv or json): ")
	_, _ = fmt.Scanln(&path)
	fd, err := deserializer.NewFileDeserializer(path)
	if err != nil {
		return err
	}
	isModel := new([]model.IndependentStudy)
	if err := fd.Deserialize(isModel); err != nil {
		return err
	}

	if err := menu.wrapper.IndependentStudyController.BaseController.InsertMany(*isModel); err != nil {
		return err
	}
	fmt.Println("\nRead file Success!")
	return nil
}

func (menu *IndependentStudyMenuStateHandler) assignNewIndependentStudy() error {
	newIndependentStudy := model.IndependentStudy{}
	fmt.Println("\nAssign new indepndent study topic")
	var turnInDateTime string = ""
	newIndependentStudy.WILProjectId = utils.GetUserInputUint("Enter WIL project group ID you want to assign : ")
	newIndependentStudy.IndependentStudyTopic = utils.GetUserInput("Enter Independent study topic: ")
	newIndependentStudy.IndependentStudyContent = utils.GetUserInput("Enter Independent study content: ")

	fmt.Println("Do you want this IS to has turn-in date?")
	fmt.Println("1. Have turn-in date")
	fmt.Println("2. No turn-in date")
	flag := utils.GetUserInputUint("your choice>>")

	if flag == 1 {
		turnInDate := utils.GetUserInput("Enter turn-in date [YYYY-mm-dd]: ")
		turnTime := utils.GetUserInput("Enter turn-in date [hh:mm:ss]: ")
		turnInDateTime = turnInDate + " " + turnTime
	}
	if err := menu.wrapper.IndependentStudyController.CreateIndependentStudy(&newIndependentStudy, turnInDateTime); err != nil {
		return err
	}
	fmt.Println("\nSuccessfully assign WIL!")
	return nil
}

func (menu *IndependentStudyMenuStateHandler) findISByID() error {
	id := utils.GetUserInputUint("Enter IS ID:")
	fmt.Println("")
	if err := menu.wilInformationRenderer(id); err != nil {
		return err
	}
	return nil
}
