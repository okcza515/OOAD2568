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
	"strconv"
	"time"
)

type ApplicantRegistrationService interface {
	RegisterManually(scanner *bufio.Scanner)
	RegisterFromFile(scanner *bufio.Scanner)
	SelectApplicationRound() *model.ApplicationRound
	SelectFacultyAndDepartment() (*commonModel.Faculty, *commonModel.Department)
	SaveReportForApplicant(applicantID uint, roundID uint, faculty *commonModel.Faculty, department *commonModel.Department, status string)
}

type applicantRegistrationService struct {
	applicantCtrl         *controller.ApplicantController
	applicationRoundCtrl  *controller.ApplicationRoundController
	applicationReportCtrl *controller.ApplicationReportController
	facultyCtrl           *common.FacultyController
	departmentCtrl        *common.DepartmentController
	criteriaCtrl          *controller.ApplicationCriteriaController
}

func NewApplicantRegistrationService(
	applicantCtrl *controller.ApplicantController,
	applicationRoundCtrl *controller.ApplicationRoundController,
	applicationReportCtrl *controller.ApplicationReportController,
	facultyCtrl *common.FacultyController,
	departmentCtrl *common.DepartmentController,
) ApplicantRegistrationService {
	return &applicantRegistrationService{
		applicantCtrl:         applicantCtrl,
		applicationRoundCtrl:  applicationRoundCtrl,
		applicationReportCtrl: applicationReportCtrl,
		facultyCtrl:           facultyCtrl,
		departmentCtrl:        departmentCtrl,
		criteriaCtrl:          controller.NewApplicationCriteriaController(),
	}
}

func (s *applicantRegistrationService) RegisterManually(scanner *bufio.Scanner) {
	var applicant model.Applicant

	// Collecting applicant details
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

	round := s.SelectApplicationRound()
	if round == nil {
		return
	}

	// Apply additional form strategy if needed
	strategy, err := controller.GetFormStrategy(round.RoundName)
	if err == nil && strategy != nil {
		if err := strategy.ApplyForm(&applicant); err != nil {
			fmt.Println("Error applying additional form:", err)
			return
		}
	}

	// Register the applicant
	if err := s.applicantCtrl.RegisterApplicant(&applicant); err != nil {
		fmt.Println("Registration failed:", err)
		return
	}

	// Select faculty and department
	faculty, department := s.SelectFacultyAndDepartment()
	if faculty == nil || department == nil {
		return
	}

	// Build criteria for applicant and determine status
	compositeCriteria := s.criteriaCtrl.BuildCriteriaForApplicant(round.RoundName, faculty.Name, department.Name)
	status := model.Pending
	if compositeCriteria.IsSatisfiedBy(applicant) {
		status = model.Pending
	} else {
		status = model.Rejected
	}

	// Save application report
	s.SaveReportForApplicant(applicant.ApplicantID, round.RoundID, faculty, department, string(status))
	// fmt.Println("Registration successful! Your Applicantion Report ID is:", applicant.ApplicantID)
	util.WaitForEnter()
}

func (s *applicantRegistrationService) RegisterFromFile(scanner *bufio.Scanner) {
	fmt.Print("Enter CSV or JSON file path: ")
	scanner.Scan()
	filePath := scanner.Text()

	// Read applicants from the file
	applicants, err := s.applicantCtrl.ReadApplicantsFromFile(filePath)
	if err != nil {
		fmt.Println("Error reading applicants from file:", err)
		return
	}

	// Register applicants from the file
	for _, a := range applicants {
		round := s.SelectApplicationRound()
		if round == nil {
			fmt.Println("No valid round selected. Skipping applicant.")
			continue
		}

		// Apply additional form strategy if needed
		strategy, err := controller.GetFormStrategy(round.RoundName)
		if err == nil && strategy != nil {
			if err := strategy.ApplyForm(&a); err != nil {
				fmt.Printf("Error applying form for %s %s: %v\n", a.FirstName, a.LastName, err)
				continue
			}
		}

		// Register the applicant
		if err := s.applicantCtrl.RegisterApplicant(&a); err != nil {
			fmt.Printf("Failed to register %s %s: %v\n", a.FirstName, a.LastName, err)
			continue
		}

		// Select faculty and department
		faculty, department := s.SelectFacultyAndDepartment()
		if faculty == nil || department == nil {
			continue
		}

		// Build criteria for applicant and determine status
		compositeCriteria := s.criteriaCtrl.BuildCriteriaForApplicant(round.RoundName, faculty.Name, department.Name)
		status := model.Pending
		if compositeCriteria.IsSatisfiedBy(a) {
			status = model.Pending
		} else {
			status = model.Rejected
		}

		// Save application report
		s.SaveReportForApplicant(a.ApplicantID, round.RoundID, faculty, department, string(status))
		fmt.Println("Registration successful! Your Applicant ID is:", a.ApplicantID)
	}
}

func (s *applicantRegistrationService) SelectApplicationRound() *model.ApplicationRound {
	rounds, err := s.applicationRoundCtrl.GetAllRounds()
	if err != nil || len(rounds) == 0 {
		fmt.Println("Error retrieving application rounds.")
		return nil
	}

	// Display available application rounds
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

func (s *applicantRegistrationService) SelectFacultyAndDepartment() (*commonModel.Faculty, *commonModel.Department) {
	// Retrieve faculties
	faculties, err := s.facultyCtrl.GetAll()
	if err != nil || len(faculties) == 0 {
		fmt.Println("Error retrieving faculties.")
		return nil, nil
	}

	// Display available faculties
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

	// Retrieve departments for selected faculty
	departments, err := s.departmentCtrl.GetByFaculty(selectedFaculty.Name)
	if err != nil || len(departments) == 0 {
		fmt.Println("Error retrieving departments.")
		return nil, nil
	}

	// Display available departments
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
	return selectedFaculty, departments[deptChoice-1]
}

func (s *applicantRegistrationService) SaveReportForApplicant(applicantID uint, roundID uint, faculty *commonModel.Faculty, department *commonModel.Department, status string) {
	report := model.ApplicationReport{
		ApplicantID:         applicantID,
		ApplicationRoundsID: roundID,
		Faculty:             faculty,
		Department:          department,
		ApplicationStatuses: model.ApplicationStatus(status),
	}

	err := s.applicationReportCtrl.SaveApplicationReport(&report)
	if err != nil {
		fmt.Printf("\n\033[1;31m[ERROR]\033[0m Failed to save report for applicant ID %d: %v\n", applicantID, err)
	} else {
		fmt.Println("\n\033[1;32m[SUCCESS]\033[0m Registration completed.")
		fmt.Printf("\033[1;36mApplication Report ID:\033[0m \033[1;34m%d\033[0m\n", report.ApplicationReportID)
		fmt.Println("Please remember this ID. You will need it for scheduling an interview or checking your status later.")
	}

	time.Sleep(2 * time.Second)
}
