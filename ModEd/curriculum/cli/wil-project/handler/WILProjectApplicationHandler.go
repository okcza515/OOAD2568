// MEP-1010 Work Integrated Learning (WIL)
package handler

import (
	"ModEd/asset/util"
	"ModEd/core"
	"ModEd/core/cli"
	"ModEd/core/handler"
	"ModEd/curriculum/controller"
	"ModEd/curriculum/model"
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

	numStudents := core.ExecuteUserInputStep(core.UintInputStep{
		PromptText:    "\nHow many students are in the project? 2 or 3: ",
		FieldNameText: "Number of Students",
	}).(uint)
	for numStudents != 2 && numStudents != 3 {
		fmt.Println("Invalid input. Please enter 2 or 3.")
		// numStudents = int(core.GetUserInputUint("\nHow many students are in the project? 2 or 3: "))
		numStudents = core.ExecuteUserInputStep(core.UintInputStep{
			PromptText:    "\nHow many students are in the project? 2 or 3: ",
			FieldNameText: "Number of Students",
		}).(uint)
	}

	var StudentsId []string
	studentIdSet := make(map[string]bool)
	for uint(len(StudentsId)) < numStudents {
		studentId := core.ExecuteUserInputStep(core.StringInputStep{
			PromptText:    "\nEnter Student ID: ",
			FieldNameText: "Student ID",
		}).(string)
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

	WILProjectApplicationModel.ProjectName = core.ExecuteUserInputStep(core.StringInputStep{
		PromptText:    "\nEnter Project Name: ",
		FieldNameText: "Project Name",
	}).(string)
	WILProjectApplicationModel.ProjectDetail = core.ExecuteUserInputStep(core.StringInputStep{
		PromptText:    "\nEnter Project Detail: ",
		FieldNameText: "Project Detail",
	}).(string)
	WILProjectApplicationModel.Semester = core.ExecuteUserInputStep(core.StringInputStep{
		PromptText:    "\nEnter Semester: ",
		FieldNameText: "Semester",
	}).(string)
	WILProjectApplicationModel.CompanyId = core.ExecuteUserInputStep(core.UintInputStep{
		PromptText:    "\nEnter Company Id: ",
		FieldNameText: "Company Id",
	}).(uint)
	WILProjectApplicationModel.Mentor = core.ExecuteUserInputStep(core.StringInputStep{
		PromptText:    "\nEnter Mentor Name: ",
		FieldNameText: "Mentor Name",
	}).(string)
	WILProjectApplicationModel.AdvisorId = core.ExecuteUserInputStep(core.UintInputStep{
		PromptText:    "\nEnter Advisor Id: ",
		FieldNameText: "Advisor Id",
	}).(uint)

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
	projectID := core.ExecuteUserInputStep(core.UintInputStep{
		PromptText:    "\nEnter the Project ID to edit: ",
		FieldNameText: "Project ID",
	}).(uint)
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
		column := core.ExecuteUserInputStep(core.UintInputStep{
			PromptText:    "\nEnter the number of the column to update (1-10): ",
			FieldNameText: "Column Number",
		}).(uint)
		if column == 10 {
			break
		}

		switch column {
		case 1:
			selectedApplication.ProjectName = core.ExecuteUserInputStep(core.StringInputStep{
				PromptText:    "\nEnter new Project Name: ",
				FieldNameText: "Project Name",
			}).(string)
		case 2:
			selectedApplication.ProjectDetail = core.ExecuteUserInputStep(core.StringInputStep{
				PromptText:    "\nEnter new Project Detail: ",
				FieldNameText: "Project Detail",
			}).(string)
		case 3:
			selectedApplication.Semester = core.ExecuteUserInputStep(core.StringInputStep{
				PromptText:    "\nEnter new Semester: ",
				FieldNameText: "Semester",
			}).(string)
		case 4:
			selectedApplication.CompanyId = core.ExecuteUserInputStep(core.UintInputStep{
				PromptText:    "\nEnter new Company ID: ",
				FieldNameText: "Company ID",
			}).(uint)
		case 5:
			selectedApplication.Mentor = core.ExecuteUserInputStep(core.StringInputStep{
				PromptText:    "\nEnter new Mentor Name: ",
				FieldNameText: "Mentor Name",
			}).(string)
		case 6:
			selectedApplication.AdvisorId = core.ExecuteUserInputStep(core.UintInputStep{
				PromptText:    "\nEnter new Advisor ID: ",
				FieldNameText: "Advisor ID",
			}).(uint)
		case 7:
			for {
				newStatus := core.ExecuteUserInputStep(core.StringInputStep{
					PromptText:    "\nEnter new Application Status (e.g., Pending, Approved, Rejected): ",
					FieldNameText: "Application Status",
				}).(string)
				if newStatus == string(model.WIL_APP_PENDING) || newStatus == string(model.WIL_APP_APPROVED) || newStatus == string(model.WIL_APP_REJECTED) {
					selectedApplication.ApplicationStatus = newStatus
					break
				}
				fmt.Println("Invalid Application Status. Please enter a valid status (Pending, Approved, Rejected).")
			}
		case 8:
			newDate := core.ExecuteUserInputStep(core.StringInputStep{
				PromptText:    "\nEnter new Turn-in Date (YYYY-MM-DD): ",
				FieldNameText: "Turn-in Date",
			}).(string)
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
	searchTerm := core.ExecuteUserInputStep(core.StringInputStep{
		PromptText:    "\nEnter search term (Project Name or ID): ",
		FieldNameText: "Search Term",
	}).(string)

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
