// MEP-1003 Student Recruitment
package cli

import (
	"ModEd/recruit/controller"
	"ModEd/recruit/model"
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/google/uuid"
)

func RegisterApplicantCLI(
	applicantCtrl *controller.ApplicantController,
	applicationRoundCtrl *controller.ApplicationRoundController,
	applicationReportCtrl *controller.ApplicationReportController,
	facultyCtrl *controller.FacultyController,
	departmentCtrl *controller.DepartmentController,
) {

	var firstName, lastName, email, address, phone, hsProgram string
	var gpax, tgat1, tgat2, tgat3, tpat1, tpat2, tpat3, tpat4, tpat5 float32

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Enter First Name: ")
	scanner.Scan()
	firstName = scanner.Text()

	fmt.Print("Enter Last Name: ")
	scanner.Scan()
	lastName = scanner.Text()

	fmt.Print("Enter Email: ")
	scanner.Scan()
	email = scanner.Text()

	fmt.Print("Enter Address: ")
	scanner.Scan()
	address = scanner.Text()

	fmt.Print("Enter Phone Number: ")
	scanner.Scan()
	phone = scanner.Text()

	fmt.Print("Enter HS Program: ")
	scanner.Scan()
	hsProgram = scanner.Text()

	fmt.Print("Enter GPAX: ")
	scanner.Scan()
	parsedGpax, _ := strconv.ParseFloat(scanner.Text(), 32)
	gpax = float32(parsedGpax)

	fmt.Print("Enter TGAT1 Score: ")
	scanner.Scan()
	parsedTgat1, _ := strconv.ParseFloat(scanner.Text(), 32)
	tgat1 = float32(parsedTgat1)

	fmt.Print("Enter TGAT2 Score: ")
	scanner.Scan()
	parsedTgat2, _ := strconv.ParseFloat(scanner.Text(), 32)
	tgat2 = float32(parsedTgat2)

	fmt.Print("Enter TGAT3 Score: ")
	scanner.Scan()
	parsedTgat3, _ := strconv.ParseFloat(scanner.Text(), 32)
	tgat3 = float32(parsedTgat3)

	fmt.Print("Enter TPAT1 Score: ")
	scanner.Scan()
	parsedTpat1, _ := strconv.ParseFloat(scanner.Text(), 32)
	tpat1 = float32(parsedTpat1)

	fmt.Print("Enter TPAT2 Score: ")
	scanner.Scan()
	parsedTpat2, _ := strconv.ParseFloat(scanner.Text(), 32)
	tpat2 = float32(parsedTpat2)

	fmt.Print("Enter TPAT3 Score: ")
	scanner.Scan()
	parsedTpat3, _ := strconv.ParseFloat(scanner.Text(), 32)
	tpat3 = float32(parsedTpat3)

	fmt.Print("Enter TPAT4 Score: ")
	scanner.Scan()
	parsedTpat4, _ := strconv.ParseFloat(scanner.Text(), 32)
	tpat4 = float32(parsedTpat4)

	fmt.Print("Enter TPAT5 Score: ")
	scanner.Scan()
	parsedTpat5, _ := strconv.ParseFloat(scanner.Text(), 32)
	tpat5 = float32(parsedTpat5)

	applicant := model.Applicant{
		ApplicantID: uuid.New(),
		FirstName:   firstName,
		LastName:    lastName,
		Email:       email,
		Address:     address,
		Phonenumber: phone,
		GPAX:        gpax,
		HS_Program:  hsProgram,
		TGAT1:       tgat1,
		TGAT2:       tgat2,
		TGAT3:       tgat3,
		TPAT1:       tpat1,
		TPAT2:       tpat2,
		TPAT3:       tpat3,
		TPAT4:       tpat4,
		TPAT5:       tpat5,
	}

	err := applicantCtrl.RegisterApplicant(&applicant)
	if err != nil {
		fmt.Println("Registration failed:", err)
		return
	}

	rounds, err := applicationRoundCtrl.GetAllRounds()
	if err != nil {
		fmt.Println("Error retrieving application rounds:", err)
		return
	}

	// แสดงรายการรอบการสมัครให้ผู้ใช้เลือก
	fmt.Println("\n==== Available Application Rounds ====")
	for i, round := range rounds {
		fmt.Printf("%d. %s\n", i+1, round.RoundName)
	}

	// ให้ผู้ใช้เลือก Application Round
	var roundChoice int
	fmt.Print("Select an application round: ")
	fmt.Scan(&roundChoice)

	// ตรวจสอบว่าผู้ใช้เลือกหมายเลขที่ถูกต้องหรือไม่
	if roundChoice < 1 || roundChoice > len(rounds) {
		fmt.Println("Invalid choice. Please select a valid round.")
		return
	}

	selectedRound := rounds[roundChoice-1]
	fmt.Printf("You have selected the %s round.\n", selectedRound.RoundName)

	// ดึงรายการคณะทั้งหมด
	faculties, err := facultyCtrl.GetAllFaculties()
	if err != nil {
		fmt.Println("Error retrieving faculties:", err)
		return
	}

	// แสดงรายการคณะให้ผู้ใช้เลือก
	fmt.Println("\n==== Available Faculties ====")
	for i, faculty := range faculties {
		fmt.Printf("%d. %s\n", i+1, faculty.Name)
	}

	var facultyChoice int
	fmt.Print("Select a faculty: ")
	fmt.Scan(&facultyChoice)

	if facultyChoice < 1 || facultyChoice > len(faculties) {
		fmt.Println("Invalid choice. Please select a valid faculty.")
		return
	}

	selectedFaculty := faculties[facultyChoice-1]
	fmt.Printf("You have selected the Faculty of %s.\n", selectedFaculty.Name)

	// ดึงรายการสาขาในคณะที่เลือก
	departments, err := departmentCtrl.GetDepartmentsByFacultyID(selectedFaculty.FacultyID)
	if err != nil {
		fmt.Println("Error retrieving departments:", err)
		return
	}

	// แสดงรายการสาขาให้ผู้ใช้เลือก
	fmt.Println("\n==== Available Departments ====")
	for i, department := range departments {
		fmt.Printf("%d. %s\n", i+1, department.Name)
	}

	var departmentChoice int
	fmt.Print("Select a department: ")
	fmt.Scan(&departmentChoice)

	if departmentChoice < 1 || departmentChoice > len(departments) {
		fmt.Println("Invalid choice. Please select a valid department.")
		return
	}

	selectedDepartment := departments[departmentChoice-1]
	fmt.Printf("You have selected the Department of %s.\n", selectedDepartment.Name)

	applicationReport := model.ApplicationReport{
		ApplicationReportID: uuid.New(),
		ApplicantID:         applicant.ApplicantID, // ใช้ ApplicantID
		ApplicationRoundsID: selectedRound.RoundID, // ใช้ RoundID
		FacultyID:           selectedFaculty.FacultyID,
		DepartmentID:        selectedDepartment.DepartmentID,
		ApplicationStatuses: model.Pending,
	}

	err = applicationReportCtrl.SaveApplicationReport(&applicationReport)
	if err != nil {
		fmt.Println("Failed to save application report:", err)
		return
	}

	fmt.Println("Registration successful! Your Applicant ID is:", applicant.ApplicantID)
}
