// MEP-1010 Work Integrated Learning (WIL)
package handler

import (
	"ModEd/asset/util"
	"ModEd/core/cli"
	"ModEd/core/handler"
	"ModEd/curriculum/controller"
	"ModEd/curriculum/model"
	"ModEd/curriculum/utils"
	"errors"
	"fmt"
	"strings"
	"time"
)

type WILProjectApplicationMenuStateHandler struct {
	manager *cli.CLIMenuStateManager
	wrapper *controller.WILModuleWrapper

	wilModuleMenuStateHandler *WILModuleMenuStateHandler
	insertHandlerStrategy     *handler.InsertHandlerStrategy[model.WILProjectApplication]
}

func NewWILProjectApplicationMenuStateHandler(
	manager *cli.CLIMenuStateManager, wrapper *controller.WILModuleWrapper, wilModuleMenuStateHandler *WILModuleMenuStateHandler,
) *WILProjectApplicationMenuStateHandler {
	return &WILProjectApplicationMenuStateHandler{
		manager:                   manager,
		wrapper:                   wrapper,
		wilModuleMenuStateHandler: wilModuleMenuStateHandler,
		insertHandlerStrategy:     handler.NewInsertHandlerStrategy[model.WILProjectApplication](wrapper.WILProjectApplicationController),
	}
}

func (menu *WILProjectApplicationMenuStateHandler) Render() {
	fmt.Println("\nWIL Project Application Menu:")
	fmt.Println("1. Create WIL Project Application")
	fmt.Println("2. Edit WIL Project Application")
	fmt.Println("3. Search WIL Project Application")
	fmt.Println("4. List all WIL Project Application")
	fmt.Println("5. Load WIL Project Application From file")
	fmt.Println("back: Exit the module")
}

func (menu *WILProjectApplicationMenuStateHandler) HandleUserInput(input string) error {

	switch input {
	case "1":
		err := menu.createWILProjectApplication()
		if err != nil {
			fmt.Println("error! cannot use this function")
		}
	case "2":
		err := menu.editWILProjectApplication()
		if err != nil {
			fmt.Println("error! cannot use this function")
		}
	case "3":
		err := menu.searchWILProjectApplication()
		if err != nil {
			fmt.Println("error! cannot use this function")
		}
	case "4":
		_, err := menu.listAllWILProjectApplication()
		if err != nil {
			fmt.Println("error! cannot use this function")
		}
	case "5":
		err := menu.insertHandlerStrategy.Execute()
		if err != nil {
			fmt.Println(err.Error())
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

func (menu *WILProjectApplicationMenuStateHandler) createWILProjectApplication() error {
	WILProjectApplicationModel := model.WILProjectApplication{}

	numStudents := int(utils.GetUserInputUint("\nHow many students are in the project? 2 or 3: "))
	for numStudents != 2 && numStudents != 3 {
		fmt.Println("Invalid input. Please enter 2 or 3.")
		numStudents = int(utils.GetUserInputUint("\nHow many students are in the project? 2 or 3: "))
	}

	var StudentsId []string
	studentIdSet := make(map[string]bool)
	for len(StudentsId) < numStudents {
		studentId := utils.GetUserInput("\nEnter Student ID: ")
		if studentIdSet[studentId] {
			fmt.Println("\nStudent ID already exists. Please enter a different ID.")
			continue
		}
		studentIdSet[studentId] = true

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
	timeNow := time.Now()
	WILProjectApplicationModel.TurninDate = &timeNow

	result := menu.wrapper.WILProjectApplicationController.RegisterWILProjectsApplication(WILProjectApplicationModel, StudentsId)
	if result != nil {
		fmt.Println("\nError for WIL Project Application:", result)
		return errors.New("error! cannot create a WIL Project application")
	}
	return nil
}

func (menu *WILProjectApplicationMenuStateHandler) listAllWILProjectApplication() ([]model.WILProjectApplication, error) {
	fmt.Println("WIL Project Application List")
	applications, err := menu.wrapper.WILProjectApplicationController.ListWILProjectApplication()
	if err != nil {
		return nil, errors.New("error! cannot retrieve WIL Project application data")
	}

	for _, application := range applications {
		fmt.Printf("%s\n", application.ToString())
		fmt.Printf("Advisor %s %s\n", application.Advisor.FirstName, application.Advisor.LastName)
		fmt.Println("Students")
		for _, student := range application.Students {
			fmt.Printf("%s %s %s\n", student.StudentId, student.Student.FirstName, student.Student.LastName)
		}
		fmt.Println("===========================================================")
	}
	return applications, nil
}

func (menu *WILProjectApplicationMenuStateHandler) editWILProjectApplication() error {
	// Step 1: Display all WIL Project Applications
	fmt.Println("WIL Project Application List")
	applications, err := menu.listAllWILProjectApplication()

	if err != nil {
		return err
	}

	// Step 2: Enter the Project ID
	projectID := uint(utils.GetUserInputUint("\nEnter the Project ID to edit: "))
	var selectedApplication *model.WILProjectApplication
	for _, application := range applications {
		if application.ID == projectID {
			selectedApplication = &application
			break
		}
	}

	if selectedApplication == nil {
		fmt.Println("Invalid Project ID.")
		return nil
	}

	// Step 4: Select a column to update
	for {
		showProjectDetail(*selectedApplication)
		column := utils.GetUserInputUint("\nEnter the number of the column to update (1-10): ")
		if column == 10 {
			break
		}

		switch column {
		case 1:
			selectedApplication.ProjectName = utils.GetUserInput("\nEnter new Project Name: ")
		case 2:
			selectedApplication.ProjectDetail = utils.GetUserInput("\nEnter new Project Detail: ")
		case 3:
			selectedApplication.Semester = utils.GetUserInput("\nEnter new Semester: ")
		case 4:
			selectedApplication.CompanyId = uint(utils.GetUserInputUint("\nEnter new Company ID: "))
		case 5:
			selectedApplication.Mentor = utils.GetUserInput("\nEnter new Mentor Name: ")
		case 6:
			selectedApplication.AdvisorId = utils.GetUserInputUint("\nEnter new Advisor ID: ")
		case 7:
			for {
				newStatus := utils.GetUserInput("\nEnter new Application Status (e.g., Pending, Approved, Rejected): ")
				if newStatus == string(model.WIL_APP_PENDING) || newStatus == string(model.WIL_APP_APPROVED) || newStatus == string(model.WIL_APP_REJECTED) {
					selectedApplication.ApplicationStatus = newStatus
					break
				}
				fmt.Println("Invalid Application Status. Please enter a valid status (Pending, Approved, Rejected).")
			}
		case 8:
			newDate := utils.GetUserInput("\nEnter new Turn-in Date (YYYY-MM-DD): ")
			parsedDate, err := time.Parse("2006-01-02", newDate)
			if err != nil {
				fmt.Println("Invalid date format. Please try again.")
				continue
			}
			selectedApplication.TurninDate = &parsedDate
		case 9:
			fmt.Println("Editing Students is not supported in this menu.")
		default:
			fmt.Println("Invalid selection. Please try again.")
		}
	}

	err = menu.wrapper.WILProjectApplicationController.UpdateByID(*selectedApplication)
	if err != nil {
		fmt.Println("Error updating the project:", err)
		return errors.New("error! cannot update the WIL Project application")
	}

	fmt.Println("Project updated successfully.")
	return nil
}

func (menu *WILProjectApplicationMenuStateHandler) searchWILProjectApplication() error {
	searchTerm := utils.GetUserInput("\nEnter search term (Project Name or ID): ")

	applications, err := menu.wrapper.WILProjectApplicationController.ListWILProjectApplication()
	if err != nil {
		return errors.New("error! cannot retrieve WIL Project application data")
	}

	var foundApplications []model.WILProjectApplication
	for _, application := range applications {
		if strings.Contains(strings.ToLower(application.ProjectName), strings.ToLower(searchTerm)) || fmt.Sprintf("%d", application.ID) == searchTerm {
			foundApplications = append(foundApplications, application)
		}
	}

	if len(foundApplications) == 0 {
		fmt.Println("No WIL Project Applications found matching the search term.")
		return nil
	}

	fmt.Println("\nSearch Results:")
	for _, application := range foundApplications {
		fmt.Printf("%s\n", application.ToString())
		fmt.Printf("Advisor %s %s\n", application.Advisor.FirstName, application.Advisor.LastName)
		fmt.Println("Students")
		for _, student := range application.Students {
			fmt.Printf("%s %s %s\n", student.StudentId, student.Student.FirstName, student.Student.LastName)
		}
		fmt.Println("===========================================================")
	}

	return nil
}

func showProjectDetail(selectedApplication model.WILProjectApplication) {
	fmt.Println("\nSelected Project Details:")
	fmt.Printf("1. Project Name: %s\n", selectedApplication.ProjectName)
	fmt.Printf("2. Project Detail: %s\n", selectedApplication.ProjectDetail)
	fmt.Printf("3. Semester: %s\n", selectedApplication.Semester)
	fmt.Printf("4. Company ID: %d\n", selectedApplication.CompanyId)
	fmt.Printf("5. Mentor: %s\n", selectedApplication.Mentor)
	fmt.Printf("6. Advisor ID: %d\n", selectedApplication.AdvisorId)
	fmt.Printf("7. Application Status: %s\n", selectedApplication.ApplicationStatus)
	fmt.Printf("8. Turn-in Date: %s\n", selectedApplication.TurninDate)
	fmt.Println("9. Students (IDs):")
	for i, student := range selectedApplication.Students {
		fmt.Printf("   %d. %s\n", i+1, student.StudentId)
	}
	fmt.Println("10. Exit Edit Menu")
}
