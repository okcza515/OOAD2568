// MEP-1003 Student Recruitment
package cli

import (
	common "ModEd/common/controller"
	commonModel "ModEd/common/model"
	"ModEd/recruit/controller"
	"ModEd/recruit/model"
	"ModEd/recruit/util"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"
)

func ApplicantRegistrationCLI(
	applicantCtrl *controller.ApplicantController,
	applicationRoundCtrl *controller.ApplicationRoundController,
	applicationReportCtrl *controller.ApplicationReportController,
	facultyCtrl *common.FacultyController,
	departmentCtrl *common.DepartmentController,
) {
	scanner := bufio.NewScanner(os.Stdin)
	util.ClearScreen()

	fmt.Println("==== Applicant Registration ====")
	fmt.Println("1. Register manually")
	fmt.Println("2. Register using CSV/JSON file")
	fmt.Print("Select option: ")
	scanner.Scan()
	choice := scanner.Text()

	criteriaCtrl := controller.NewApplicationCriteriaController()

	switch choice {
	case "1":
		registerManually(scanner, applicantCtrl, applicationRoundCtrl, applicationReportCtrl, facultyCtrl, departmentCtrl, criteriaCtrl)
	case "2":
		registerFromFile(scanner, applicantCtrl, applicationRoundCtrl, applicationReportCtrl, facultyCtrl, departmentCtrl, criteriaCtrl)
	default:
		fmt.Println("Invalid option.")
	}

}

func registerFromFile(
	scanner *bufio.Scanner,
	applicantCtrl *controller.ApplicantController,
	applicationRoundCtrl *controller.ApplicationRoundController,
	applicationReportCtrl *controller.ApplicationReportController,
	facultyCtrl *common.FacultyController,
	departmentCtrl *common.DepartmentController,
	criteriaCtrl *controller.ApplicationCriteriaController,
) {
	fmt.Print("Enter CSV or JSON file path: ")
	scanner.Scan()
	filePath := scanner.Text()

	applicants, err := applicantCtrl.ReadApplicantsFromFile(filePath)
	if err != nil {
		fmt.Println("Error reading applicants from file:", err)
		return
	}

	for _, a := range applicants {
		round := selectApplicationRound(applicationRoundCtrl)
		if round == nil {
			fmt.Println("No valid round selected. Skipping applicant.")
			continue
		}

		var strategy controller.FormRound
		strategy, err = controller.GetFormStrategy(round.RoundName)
		if err == nil && strategy != nil {
			if applyErr := strategy.ApplyForm(&a); applyErr != nil {
				fmt.Printf("Error applying form for %s %s: %v\n", a.FirstName, a.LastName, applyErr)
				continue
			}
		}

		if regErr := applicantCtrl.RegisterApplicant(&a); regErr != nil {
			fmt.Printf("Failed to register %s %s: %v\n", a.FirstName, a.LastName, regErr)
			continue
		}

		faculty, department := selectFacultyAndDepartment(facultyCtrl, departmentCtrl)
		if faculty == nil || department == nil {
			fmt.Println("No valid faculty or department selected. Skipping applicant.")
			// continue
		}

		compositeCriteria := criteriaCtrl.BuildCriteriaForApplicant(round.RoundName, faculty.Name, department.Name)
		status := model.Pending
		if compositeCriteria.IsSatisfiedBy(a) {
			status = model.Pending
		} else {
			status = model.Rejected
		}

		saveReportForApplicant(applicationReportCtrl, a.ApplicantID, round.RoundID, faculty, department, string(status))
		fmt.Println("Registration successful! Your Applicant ID is:", a.ApplicantID)
	}
}

func registerManually(
	scanner *bufio.Scanner,
	applicantCtrl *controller.ApplicantController,
	applicationRoundCtrl *controller.ApplicationRoundController,
	applicationReportCtrl *controller.ApplicationReportController,
	facultyCtrl *common.FacultyController,
	departmentCtrl *common.DepartmentController,
	criteriaCtrl *controller.ApplicationCriteriaController,
) {
	var applicant model.Applicant

	fmt.Print("Enter First Name: ")
	scanner.Scan()
	applicant.FirstName = scanner.Text()

	fmt.Print("Enter Last Name: ")
	scanner.Scan()
	applicant.LastName = scanner.Text()

	fmt.Print("Enter Email: ")
	scanner.Scan()
	applicant.Email = scanner.Text()

	fmt.Print("Enter Address: ")
	scanner.Scan()
	applicant.Address = scanner.Text()

	fmt.Print("Enter Phone Number: ")
	scanner.Scan()
	applicant.Phonenumber = scanner.Text()

	fmt.Print("Enter HS Program: ")
	scanner.Scan()
	applicant.HighSchool_Program = scanner.Text()

	inputFloat := func(label string) float32 {
		fmt.Print(label)
		scanner.Scan()
		val, _ := strconv.ParseFloat(scanner.Text(), 32)
		return float32(val)
	}

	applicant.GPAX = inputFloat("Enter GPAX: ")
	applicant.TGAT1 = inputFloat("Enter TGAT1 Score: ")
	applicant.TGAT2 = inputFloat("Enter TGAT2 Score: ")
	applicant.TGAT3 = inputFloat("Enter TGAT3 Score: ")
	applicant.TPAT1 = inputFloat("Enter TPAT1 Score: ")
	applicant.TPAT2 = inputFloat("Enter TPAT2 Score: ")
	applicant.TPAT3 = inputFloat("Enter TPAT3 Score: ")
	applicant.TPAT4 = inputFloat("Enter TPAT4 Score: ")
	applicant.TPAT5 = inputFloat("Enter TPAT5 Score: ")

	round := selectApplicationRound(applicationRoundCtrl)
	if round == nil {
		return
	}

	var strategy controller.FormRound

	strategy, err := controller.GetFormStrategy(round.RoundName)
	if err == nil && strategy != nil {
		if err := strategy.ApplyForm(&applicant); err != nil {
			fmt.Println("Error applying additional form:", err)
			return
		}
	}

	if err := applicantCtrl.RegisterApplicant(&applicant); err != nil {
		fmt.Println("Registration failed:", err)
		return
	}

	faculty, department := selectFacultyAndDepartment(facultyCtrl, departmentCtrl)
	if faculty == nil || department == nil {
		return
	}

	compositeCriteria := criteriaCtrl.BuildCriteriaForApplicant(round.RoundName, faculty.Name, department.Name)

	status := model.Pending
	if compositeCriteria.IsSatisfiedBy(applicant) {
		status = model.Pending
	} else {
		status = model.Rejected
	}

	saveReportForApplicant(applicationReportCtrl, applicant.ApplicantID, round.RoundID, faculty, department, string(status))
	fmt.Println("Registration successful! Your Applicant ID is:", applicant.ApplicantID)
	time.Sleep(8 * time.Second)
}

func selectApplicationRound(appRoundCtrl *controller.ApplicationRoundController) *model.ApplicationRound {
	rounds, err := appRoundCtrl.GetAllRounds()
	if err != nil || len(rounds) == 0 {
		fmt.Println("Error retrieving application rounds.")
		return nil
	}

	fmt.Println("\n==== Available Application Rounds ====")
	for i, round := range rounds {
		fmt.Printf("%d. %s\n", i+1, round.RoundName)
	}

	fmt.Print("Select an application round: ")
	var choice int
	fmt.Scan(&choice)
	if choice < 1 || choice > len(rounds) {
		fmt.Println("Invalid choice.")
		return nil
	}
	return rounds[choice-1]
}

func selectFacultyAndDepartment(
	facultyCtrl *common.FacultyController,
	departmentCtrl *common.DepartmentController,
) (*commonModel.Faculty, *commonModel.Department) {
	faculties, err := facultyCtrl.GetAll()
	if err != nil || len(faculties) == 0 {
		fmt.Println("Error retrieving faculties.")
		return nil, nil
	}

	fmt.Println("\n==== Available Faculties ====")
	for i, faculty := range faculties {
		fmt.Printf("%d. %s\n", i+1, faculty.Name)
	}
	fmt.Print("Select a faculty: ")
	var facultyChoice int
	fmt.Scan(&facultyChoice)
	if facultyChoice < 1 || facultyChoice > len(faculties) {
		fmt.Println("Invalid faculty.")
		return nil, nil
	}
	selectedFaculty := faculties[facultyChoice-1]

	departments, err := departmentCtrl.GetByFaculty(selectedFaculty.Name)
	if err != nil || len(departments) == 0 {
		fmt.Println("Error retrieving departments.")
		return nil, nil
	}
	fmt.Println("\n==== Available Departments ====")
	for i, dept := range departments {
		fmt.Printf("%d. %s\n", i+1, dept.Name)
	}
	fmt.Print("Select a department: ")
	var deptChoice int
	fmt.Scan(&deptChoice)
	if deptChoice < 1 || deptChoice > len(departments) {
		fmt.Println("Invalid department.")
		return nil, nil
	}
	selectedDepartments := departments[deptChoice-1]

	return selectedFaculty, selectedDepartments
}

func saveReportForApplicant(
	appReportCtrl *controller.ApplicationReportController,
	applicantID uint,
	roundID uint,
	faculty *commonModel.Faculty,
	department *commonModel.Department,
	status string,
) {
	report := model.ApplicationReport{
		ApplicationReportID: 0,
		ApplicantID:         applicantID,
		ApplicationRoundsID: roundID,
		Faculty:             faculty,
		Department:          department,
		ApplicationStatuses: model.ApplicationStatus(status),
	}
	err := appReportCtrl.SaveApplicationReport(&report)
	if err != nil {
		fmt.Printf("Failed to save report for applicant ID %d: %v\n", applicantID, err)
	} else {
		fmt.Printf("Application report saved for applicant ID %d\n", applicantID)
	}
}
