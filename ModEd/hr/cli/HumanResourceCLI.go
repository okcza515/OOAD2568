package main

import (
	"errors"
	"flag"
	"fmt"
	"os"

	commonController "ModEd/common/controller"
	commonModel "ModEd/common/model"
	"ModEd/hr/controller"
	hrModel "ModEd/hr/model"
	"ModEd/hr/util"

	hrUtil "ModEd/hr/util"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	// Global flags
	databasePath = flag.String("database", "data/ModEd.bin", "Path of SQLite Database")
)

func main() {
	// Parse global flags first
	flag.Parse()
	args := flag.Args()

	if len(args) < 1 {
		fmt.Println("Usage: go run humanresourcecli.go [-database=<path>] {list|add|update|delete} [options]")
		os.Exit(1)
	}

	// First argument after global flags is the command
	command := args[0]
	commandArgs := args[1:]

	// Execute the appropriate command
	switch command {
	case "list":
		listStudents(commandArgs)
	case "update":
		updateStudent(commandArgs)
	case "add":
		addStudent(commandArgs)
	case "delete":
		deleteStudent(commandArgs)
	case "updateStatus":
		updateStudentStatus(commandArgs)
	default:
		fmt.Printf("Unknown command: %s\n", command)
		fmt.Println("Available commands: list, add, update, delete")
		os.Exit(1)
	}
}

func openDatabase(database string) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(database), &gorm.Config{})
	if err != nil {
		fmt.Printf("Failed to connect to database: %v\n", err)
		os.Exit(1)
	}

	// Migrate student data from the common Student table to the HR StudentInfo table.
	if err := controller.MigrateStudentsToHR(db); err != nil {
		fmt.Printf("Migration Failed: %v\n", err)
		os.Exit(1)
	}
	return db
}

func listStudents(args []string) {
	fs := flag.NewFlagSet("list", flag.ExitOnError)
	fs.Parse(args)

	db := openDatabase(*databasePath)
	studentController := controller.CreateStudentHRController(db)
	studentInfos, err := studentController.GetAll()
	if err != nil {
		fmt.Printf("Error listing students: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Human Resource Student Info:")
	for _, s := range studentInfos {
		fmt.Printf("SID: %s | Name: %s %s | Gender: %s | CitizenID: %s | Phone: %s | Status: %s | Email: %s\n",
			s.SID, s.FirstName, s.LastName, s.Gender, s.CitizenID, s.PhoneNumber, hrUtil.StatusToString(s.Status), s.Email)
	}
}

func updateStudent(args []string) {
	fs := flag.NewFlagSet("update", flag.ExitOnError)
	studentID := fs.String("id", "", "Student ID to update")
	firstName := fs.String("fname", "", "New First Name value")
	lastName := fs.String("lname", "", "New Last Name value")
	gender := fs.String("gender", "", "New Gender value")
	citizenID := fs.String("citizenID", "", "New Citizen ID value")
	phoneNumber := fs.String("phone", "", "New Phone Number value")
	emailStudent := fs.String("email", "", "New Email value")
	fs.Parse(args)

	if *studentID == "" {
		fmt.Println("Error: Student ID is required.")
		fmt.Println("Usage: go run humanresourcecli.go [-database=<path>] update -id=<studentID> [options]")
		os.Exit(1)
	}

	db := openDatabase(*databasePath)
	studentController := controller.CreateStudentHRController(db)
	studentInfo, err := studentController.GetById(*studentID)
	if err != nil {
		fmt.Printf("Error retrieving student with ID %s: %v\n", *studentID, err)
		os.Exit(1)
	}

	// Update basic student information if provided
	if *firstName != "" {
		studentInfo.FirstName = *firstName
	}
	if *lastName != "" {
		studentInfo.LastName = *lastName
	}
	if *gender != "" {
		studentInfo.Gender = *gender
	}
	if *citizenID != "" {
		studentInfo.CitizenID = *citizenID
	}
	if *phoneNumber != "" {
		studentInfo.PhoneNumber = *phoneNumber
	}
	if *emailStudent != "" {
		studentInfo.Email = *emailStudent
	}

	if err := studentController.Update(studentInfo); err != nil {
		fmt.Printf("Failed to update student info: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Student updated successfully!")
}

func addStudent(args []string) {
	fs := flag.NewFlagSet("add", flag.ExitOnError)
	studentID := fs.String("id", "", "Student ID")
	firstName := fs.String("fname", "", "First Name")
	lastName := fs.String("lname", "", "Last Name")
	gender := fs.String("gender", "", "Gender")
	citizenID := fs.String("citizenID", "", "Citizen ID")
	phoneNumber := fs.String("phone", "", "Phone Number")
	fs.Parse(args)

	if *studentID == "" || *firstName == "" || *lastName == "" {
		fmt.Println("Error: Student ID, First Name, and Last Name are required.")
		fmt.Println("Usage: go run humanresourcecli.go [-database=<path>] add -id=<studentID> -fname=<firstName> -lname=<lastName> [options]")
		os.Exit(1)
	}

	db := openDatabase(*databasePath)
	studentController := controller.CreateStudentHRController(db)
	newStudent := hrModel.StudentInfo{
		Student: commonModel.Student{
			SID:       *studentID,
			FirstName: *firstName,
			LastName:  *lastName,
		},
		Gender:      *gender,
		CitizenID:   *citizenID,
		PhoneNumber: *phoneNumber,
	}

	if err := studentController.Insert(&newStudent); err != nil {
		fmt.Printf("Failed to add student info: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Student added successfully!")
}

func deleteStudent(args []string) {
	fs := flag.NewFlagSet("delete", flag.ExitOnError)
	studentID := fs.String("id", "", "Student ID to delete")
	fs.Parse(args)

	if *studentID == "" {
		fmt.Println("Error: Student ID is required.")
		fmt.Println("Usage: go run humanresourcecli.go [-database=<path>] delete -id=<studentID>")
		os.Exit(1)
	}

	db := openDatabase(*databasePath)
	studentController := controller.CreateStudentHRController(db)
	if err := studentController.Delete(*studentID); err != nil {
		fmt.Printf("Failed to delete student info: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Student deleted successfully!")
}

func updateStudentStatus(args []string) {
	fs := flag.NewFlagSet("updateStatus", flag.ExitOnError)
	studentID := fs.String("id", "", "Student ID to update status")
	status := fs.String("status", "", "New Status (ACTIVE, GRADUATED, or DROP)")
	fs.Parse(args)

	if *studentID == "" || *status == "" {
		fmt.Println("Error: Student ID and Status are required.")
		fmt.Println("Usage: go run humanresourcecli.go [-database=<path>] updateStatus -id=<studentID> -status=<ACTIVE|GRADUATED|DROP>")
		os.Exit(1)
	}

	// Convert the status string to enum
	newStatus, err := hrUtil.StatusFromString(*status)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	db := openDatabase(*databasePath)
	studentController := controller.CreateStudentHRController(db)

	// Use the dedicated controller method for updating status
	if err := studentController.UpdateStatus(*studentID, newStatus); err != nil {
		fmt.Printf("Failed to update student status: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Student %s status successfully updated to %s!\n", *studentID, *status)
}
func importStudents(args []string) {
	fs := flag.NewFlagSet("import", flag.ExitOnError)
	filePath := fs.String("path", "", "Path to CSV or JSON for HR student info (only studentid and HR fields).")
	fs.Parse(args)

	if *filePath == "" {
		fmt.Println("Error: File path for HR student data is required.")
		fmt.Println("Usage: go run humanresourcecli.go [-database=<path>] import -path=<path>")
		os.Exit(1)
	}

	if _, err := os.Stat(*filePath); errors.Is(err, os.ErrNotExist) {
		fmt.Printf("*** Error: File %s does not exist.\n", *filePath)
		os.Exit(1)
	}

	// Use a mapper tailored for HR data. Assume you have an HR-specific mapper that reads into hrModel.HRInfo.
	hrMapper, err := util.CreateMapper[hrModel.StudentInfo](*filePath)
	if err != nil {
		fmt.Printf("Failed to create HR mapper: %v\n", err)
		os.Exit(1)
	}

	hrRecords := hrMapper.Map() // hrRecords is []*hrModel.HRInfo

	db := openDatabase(*databasePath)
	hrController := controller.CreateStudentHRController(db)

	for _, hrRec := range hrRecords {
		// Create the common student controller.
		commonStudentController := commonController.CreateStudentController(db)
		// Retrieve common student data using the student's ID.
		commonStudent, err := commonStudentController.GetByStudentId(hrRec.SID)
		if err != nil {
			fmt.Printf("Failed to retrieve student %s from common data: %v\n", hrRec.SID, err)
			continue
		}

		// Merge the common student info with the HR-specific fields
		newStudent := hrModel.StudentInfo{
			Student: *commonStudent, // from the common Model
			Gender:  hrRec.Gender,   // HR field from CSV
			// Set additional HR fields from hrRec as needed.
		}

		// Upsert (insert or update) so that duplicate records are not created.
		if err := hrController.Upsert(&newStudent); err != nil {
			fmt.Printf("Failed to upsert student %s: %v\n", newStudent.SID, err)
			continue
		}
	}
	fmt.Println("Students imported successfully!")
}

func synchronizeStudents(args []string) {
	db := openDatabase(*databasePath)

	if err := controller.SynchronizeStudents(db); err != nil {
		fmt.Printf("Failed to synchronize students: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Students synchronized successfully!")
}
