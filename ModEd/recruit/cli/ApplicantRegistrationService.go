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
	"path/filepath"
	"strconv"
	"time"
)

type ApplicantRegistrationService interface {
	RegisterManually(scanner *bufio.Scanner)
	RegisterFromFile(scanner *bufio.Scanner)
	SelectApplicationRound() *model.ApplicationRound
	SelectFacultyAndDepartment() (*commonModel.Faculty, *commonModel.Department)
	SaveReportForApplicant(applicantID uint, roundID uint, faculty *commonModel.Faculty, department *commonModel.Department, program *commonModel.ProgramType, status string)
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

	fmt.Print("Enter Birth Date (YYYY-MM-DD): ")
	scanner.Scan()
	birthDateStr := scanner.Text()
	birthDate, err := time.Parse("2006-01-02", birthDateStr)
	if err != nil {
		fmt.Println("Invalid birth date format.")
		return
	}
	applicant.BirthDate = birthDate

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
	tgatFields := []*float32{&applicant.TGAT1, &applicant.TGAT2, &applicant.TGAT3}
	for i := range tgatFields {
		*tgatFields[i] = inputFloat(fmt.Sprintf("Enter TGAT%d Score: ", i+1))
	}

	tpatFields := []*float32{&applicant.TPAT1, &applicant.TPAT2, &applicant.TPAT3, &applicant.TPAT4, &applicant.TPAT5}
	for i := range tpatFields {
		*tpatFields[i] = inputFloat(fmt.Sprintf("Enter TPAT%d Score: ", i+1))
	}

	round := s.SelectApplicationRound()
	if round == nil {
		return
	}

	if !s.handleRoundFormData(round, &applicant) {
		return
	}

	if err := s.applicantCtrl.RegisterApplicant(&applicant); err != nil {
		fmt.Println("Registration failed:", err)
		return
	}

	faculty, department := s.SelectFacultyAndDepartment()
	if faculty == nil || department == nil {
		return
	}

	program := s.SelectProgram()

	compositeCriteria := s.criteriaCtrl.BuildCriteriaForApplicant(round.RoundName, faculty.Name, department.Name)
	status := model.Pending
	if compositeCriteria.IsSatisfiedBy(applicant) {
		status = model.Pending
	} else {
		status = model.Rejected
	}

	s.SaveReportForApplicant(applicant.ApplicantID, round.RoundID, faculty, department, &program, string(status))
	// fmt.Println("Registration successful! Your Applicantion Report ID is:", applicant.ApplicantID)
	util.WaitForEnter()
}

func (s *applicantRegistrationService) RegisterFromFile(scanner *bufio.Scanner) {

	curDir, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting current directory:", err)
		return
	}
	parentDir := filepath.Dir(curDir)

	defaultRegisDataPath := filepath.Join(parentDir, "recruit", "data", "RegisData.csv")
	filePath := defaultRegisDataPath

	applicants, err := s.applicantCtrl.ReadApplicantsFromFile(filePath)
	if err != nil {
		fmt.Println("Error reading applicants from file:", err)
		return
	}

	for _, a := range applicants {
		round := s.SelectApplicationRound()
		if round == nil {
			fmt.Println("No valid round selected. Skipping applicant.")
			continue
		}

		if !s.handleRoundFormData(round, &a) {
			fmt.Printf("Skipping %s %s due to invalid form data.\n", a.FirstName, a.LastName)
			continue
		}

		if err := s.applicantCtrl.RegisterApplicant(&a); err != nil {
			fmt.Printf("Failed to register %s %s: %v\n", a.FirstName, a.LastName, err)
			continue
		}

		faculty, department := s.SelectFacultyAndDepartment()
		if faculty == nil || department == nil {
			continue
		}

		program := s.SelectProgram()

		compositeCriteria := s.criteriaCtrl.BuildCriteriaForApplicant(round.RoundName, faculty.Name, department.Name)
		status := model.Pending
		if compositeCriteria.IsSatisfiedBy(a) {
			status = model.Pending
		} else {
			status = model.Rejected
		}

		s.SaveReportForApplicant(a.ApplicantID, round.RoundID, faculty, department, &program, string(status))
		fmt.Println("Registration successful! Your Applicant ID is:", a.ApplicantID)
	}
}

func (s *applicantRegistrationService) SelectApplicationRound() *model.ApplicationRound {
	rounds, err := s.applicationRoundCtrl.GetAllRounds()
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

func (s *applicantRegistrationService) SelectFacultyAndDepartment() (*commonModel.Faculty, *commonModel.Department) {
	faculties, err := s.facultyCtrl.GetAll()
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
	selectedFaculty := &faculties[facultyChoice-1]

	departments, err := s.departmentCtrl.GetByFaculty(selectedFaculty.Name)
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
	return selectedFaculty, &departments[deptChoice-1]
}

func (s *applicantRegistrationService) SaveReportForApplicant(applicantID uint, roundID uint, faculty *commonModel.Faculty, department *commonModel.Department, program *commonModel.ProgramType, status string) {
	report := model.ApplicationReport{
		ApplicantID:         applicantID,
		ApplicationRoundsID: roundID,
		Faculty:             faculty,
		Department:          department,
		Program:             program,
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

func (s *applicantRegistrationService) handleRoundFormData(round *model.ApplicationRound, a *model.Applicant) bool {
	strategy, err := controller.GetFormStrategy(round.RoundName)
	if err != nil || strategy == nil {
		fmt.Printf("Failed to get form strategy: %v\n", err)
		return false
	}

	bufio.NewReader(os.Stdin).ReadString('\n')
	roundData := make(map[string]string)
	scanner := bufio.NewScanner(os.Stdin)
	for _, roundField := range strategy.GetForm() {
		fmt.Printf("Enter %s: ", roundField)

		scanner.Scan()
		data := scanner.Text()
		roundData[roundField] = data
	}

	a.SetRoundInfo(roundData)
	return true
}

func (s *applicantRegistrationService) SelectProgram() commonModel.ProgramType {
	fmt.Println("\n==== Available Programs ====")
	for k, v := range commonModel.ProgramTypeLabel {
		fmt.Printf("%d. %s\n", k+1, v)
	}

	fmt.Print("Select a program: ")
	var choice int
	fmt.Scan(&choice)

	choice--
	if _, ok := commonModel.ProgramTypeLabel[commonModel.ProgramType(choice)]; !ok {
		fmt.Println("Invalid selection. Defaulting to Regular.")
		return commonModel.REGULAR
	}
	return commonModel.ProgramType(choice)
}
