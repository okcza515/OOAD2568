// MEP-1010 Work Integrated Learning (WIL)
package handler

import (
	"ModEd/curriculum/controller"
	"ModEd/curriculum/model"
	"ModEd/curriculum/utils"
	"errors"
	"fmt"
	"time"
)

func RunWILProjectApplicationHandler(controller *controller.WILModuleFacade) {
	for {
		printWILProjectApplicationModuleMenu()
		choice := utils.GetUserChoice()

		switch choice {
		case "1":
			err := createWILProjectApplication(controller)
			if err != nil {
				fmt.Println(err)
				continue
			}
			fmt.Println("\nWIL Project Application created successfully")
		case "2":
			fmt.Println("2 Not implemented yet...")
		case "3":
			fmt.Println("3 Not implemented yet...")
		case "4":
			err := listAllWILProjectApplication(controller)
			if err != nil {
				fmt.Println(err)
				continue
			}
		case "5":
			fmt.Println("5 Not implemented yet...")
		case "6":
			fmt.Println("6 Not implemented yet...")
		case "0":
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid option")
		}
	}
}

func printWILProjectApplicationModuleMenu() {
	fmt.Println("\nWIL Project Application Menu:")
	fmt.Println("1. Create WIL Project Application")
	fmt.Println("2. Edit WIL Project Application")
	fmt.Println("3. Search WIL Project Application")
	fmt.Println("4. List all WIL Project Application")
	fmt.Println("5. Get WIL Project Application By ID")
	fmt.Println("6. Delete WIL Project Application")
	fmt.Println("0. Exit WIL Module")
}

func createWILProjectApplication(controller *controller.WILModuleFacade) error {
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

	result := controller.WILProjectApplicationController.RegisterWILProjectsApplication(WILProjectApplicationModel, StudentsId)
	if result != nil {
		fmt.Println("\nError for WIL Project Application:", result)
		return errors.New("Error! cannot create a WIL Project application")
	}
	return nil
}

func listAllWILProjectApplication(controller *controller.WILModuleFacade) error {
	fmt.Println("\nWIL Project Application List\n")
	applications, err := controller.WILProjectApplicationController.ListWILProjectApplication()
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
