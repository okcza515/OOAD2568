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
	"time"
)

type WILProjectApplicationMenuStateHandler struct {
	manager *cli.CLIMenuStateManager
	proxy   *controller.WILModuleProxy

	wilModuleMenuStateHandler *WILModuleMenuStateHandler
}

func NewWILProjectApplicationMenuStateHandler(
	manager *cli.CLIMenuStateManager, proxy *controller.WILModuleProxy, wilModuleMenuStateHandler *WILModuleMenuStateHandler,
) *WILProjectApplicationMenuStateHandler {
	return &WILProjectApplicationMenuStateHandler{
		manager:                   manager,
		proxy:                     proxy,
		wilModuleMenuStateHandler: wilModuleMenuStateHandler,
	}
}

func (menu *WILProjectApplicationMenuStateHandler) Render() {
	fmt.Println("\nWIL Project Application Menu:")
	fmt.Println("1. Create WIL Project Application")
	fmt.Println("2. Edit WIL Project Application")
	fmt.Println("3. Search WIL Project Application")
	fmt.Println("4. List all WIL Project Application")
	fmt.Println("5. Get WIL Project Application By ID")
	fmt.Println("6. Delete WIL Project Application")
	fmt.Println("0. Exit WIL Module")
}

func (menu *WILProjectApplicationMenuStateHandler) HandleUserInput(input string) error {

	switch input {
	case "1":
		err := menu.createWILProjectApplication()
		if err != nil {
			fmt.Println("error! cannot use this function")
		}
	case "2":
		err := menu.listAllWILProjectApplication()
		if err != nil {
			fmt.Println("error! cannot use this function")
		}
	case "3":
		fmt.Println("Get detail of an Instrument")
	case "4":
		fmt.Println("Update an Instrument")
	case "5":
		fmt.Println("Delete an Instrument")
	case "0":
		menu.manager.SetState(menu.wilModuleMenuStateHandler)
		return nil
	default:
		fmt.Println("Invalid Command")
	}

	util.PressEnterToContinue()

	return nil
}

func (menu *WILProjectApplicationMenuStateHandler) createWILProjectApplication() error {
	WILProjectApplicationModel := model.WILProjectApplication{}

	fmt.Println("\nRegistering WILProjectApplication model")

	numStudents := int(utils.GetUserInputUint("\nHow many students are in the project? 2 or 3: "))
	var StudentsId []string
	for len(StudentsId) < numStudents {
		studentId := utils.GetUserInput("\nEnter Student ID: ")
		for _, id := range StudentsId {
			if id == studentId {
				fmt.Println("\nStudent ID already exists. Please enter a different ID.")
				continue
			}
		}

		// TODO: Check if the student ID is valid
		// If valid, append to the slice
		// Else continue
		StudentsId = append(StudentsId, studentId)
	}

	WILProjectApplicationModel.ProjectName = utils.GetUserInput("\nEnter Project Name: ")
	WILProjectApplicationModel.ProjectDetail = utils.GetUserInput("\nEnter Project Detail: ")
	WILProjectApplicationModel.Semester = utils.GetUserInput("\nEnter Semester: ")
	WILProjectApplicationModel.CompanyId = uint(utils.GetUserInputUint("\nEnter Company Id: "))
	WILProjectApplicationModel.Mentor = utils.GetUserInput("\nEnter Mentor Name: ")
	WILProjectApplicationModel.AdvisorId = utils.GetUserInputUint("\nEnter Advisor Id: ")

	WILProjectApplicationModel.ApplicationStatus = string(model.WIL_APP_PENDING)
	WILProjectApplicationModel.TurninDate = time.Now().Format("2006-01-02 15:04:05")

	result := menu.proxy.WILProjectApplicationController.RegisterWILProjectsApplication(WILProjectApplicationModel, StudentsId)
	if result != nil {
		fmt.Println("\nError for WIL Project Application:", result)
		return errors.New("Error! cannot create a WIL Project application")
	}
	return nil
}

func (menu *WILProjectApplicationMenuStateHandler) listAllWILProjectApplication() error {
	fmt.Println("\nWIL Project Application List\n")
	applications, err := menu.proxy.WILProjectApplicationController.ListWILProjectApplication()
	if err != nil {
		return errors.New("Error! cannot retrieve WIL Project application data")
	}

	for _, application := range applications {
		fmt.Printf("%s %s %s\n", application.ProjectName, application.Advisor.FirstName, application.Advisor.LastName)
		fmt.Println("Students")
		for _, student := range application.Students {
			fmt.Printf("%s %s %s\n", student.StudentId, student.Student.FirstName, student.Student.LastName)
		}
		fmt.Println("===========================================================")
	}
	return nil
}
