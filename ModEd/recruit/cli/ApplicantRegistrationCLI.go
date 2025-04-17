package cli

import (
	"ModEd/recruit/controller"
	"ModEd/recruit/model"
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func RegisterApplicantCLI(
	applicantCtrl *controller.ApplicantController,
	applicationRoundCtrl *controller.ApplicationRoundController,
	applicationReportCtrl *controller.ApplicationReportController,
	facultyCtrl *controller.FacultyController,
	departmentCtrl *controller.DepartmentController,
) {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("==== Applicant Registration ====")
	fmt.Println("1. Register manually")
	fmt.Println("2. Register using CSV/JSON file")
	fmt.Print("Select option: ")
	scanner.Scan()
	choice := scanner.Text()

	switch choice {
	case "1":
		registerManually(scanner, applicantCtrl, applicationRoundCtrl, applicationReportCtrl, facultyCtrl, departmentCtrl)
	case "2":
		registerFromFile(scanner, applicantCtrl, applicationRoundCtrl, applicationReportCtrl, facultyCtrl, departmentCtrl)
	default:
		fmt.Println("Invalid option.")
	}
}

func registerFromFile(
	scanner *bufio.Scanner,
	applicantCtrl *controller.ApplicantController,
	applicationRoundCtrl *controller.ApplicationRoundController,
	applicationReportCtrl *controller.ApplicationReportController,
	facultyCtrl *controller.FacultyController,
	departmentCtrl *controller.DepartmentController,
) {
	fmt.Print("Enter CSV or JSON file path: ")
	scanner.Scan()
	filePath := scanner.Text()

	applicants, err := applicantCtrl.ReadApplicantsFromFile(filePath)
	if err != nil {
		fmt.Println("Error reading applicants from file:", err)
		return
	}

	round := selectApplicationRound(applicationRoundCtrl)
	if round == nil {
		return
	}

	faculty, department := selectFacultyAndDepartment(facultyCtrl, departmentCtrl)
	if faculty == nil || department == nil {
		return
	}

	for _, a := range applicants {
		err := applicantCtrl.RegisterApplicant(&a)
		if err != nil {
			fmt.Printf("Failed to register %s %s: %v\n", a.FirstName, a.LastName, err)
			continue
		}
		saveReportForApplicant(applicationReportCtrl, a.ApplicantID, round.RoundID, faculty.FacultyID, department.DepartmentID)
	}
}

func registerManually(
	scanner *bufio.Scanner,
	applicantCtrl *controller.ApplicantController,
	applicationRoundCtrl *controller.ApplicationRoundController,
	applicationReportCtrl *controller.ApplicationReportController,
	facultyCtrl *controller.FacultyController,
	departmentCtrl *controller.DepartmentController,
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

	if err := applicantCtrl.RegisterApplicant(&applicant); err != nil {
		fmt.Println("Registration failed:", err)
		return
	}

	round := selectApplicationRound(applicationRoundCtrl)
	if round == nil {
		return
	}

	faculty, department := selectFacultyAndDepartment(facultyCtrl, departmentCtrl)
	if faculty == nil || department == nil {
		return
	}

	saveReportForApplicant(applicationReportCtrl, applicant.ApplicantID, round.RoundID, faculty.FacultyID, department.DepartmentID)
	fmt.Println("Registration successful! Your Applicant ID is:", applicant.ApplicantID)
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
	facultyCtrl *controller.FacultyController,
	departmentCtrl *controller.DepartmentController,
) (*model.Faculty, *model.Department) {
	faculties, err := facultyCtrl.GetAllFaculties()
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

	departments, err := departmentCtrl.GetDepartmentsByFacultyID(selectedFaculty.FacultyID)
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
	return &selectedFaculty, &departments[deptChoice-1]
}

func saveReportForApplicant(
	appReportCtrl *controller.ApplicationReportController,
	applicantID uint,
	roundID uint,
	facultyID uint,
	departmentID uint,
) {
	report := model.ApplicationReport{
		ApplicationReportID: 0,
		ApplicantID:         applicantID,
		ApplicationRoundsID: roundID,
		FacultyID:           facultyID,
		DepartmentID:        departmentID,
		ApplicationStatuses: model.Pending,
	}
	err := appReportCtrl.SaveApplicationReport(&report)
	if err != nil {
		fmt.Printf("Failed to save report for applicant ID %d: %v\n", applicantID, err)
	} else {
		fmt.Printf("Application report saved for applicant ID %d\n", applicantID)
	}
}
