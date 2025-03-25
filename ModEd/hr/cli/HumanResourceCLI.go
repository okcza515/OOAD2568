package main

// import (
// 	"flag"
// 	"fmt"
// 	"os"

// 	"ModEd/hr/controller"
// 	hrUtil "ModEd/hr/util"
// )

// var (
// 	databasePath = flag.String("database", "data/ModEd.bin", "Path of SQLite Database")
// )

// func main() {
// 	flag.Parse()
// 	args := flag.Args()

// 	if len(args) < 1 {
// 		fmt.Println("Usage: go run humanresourcecli.go [-database=<path>] {list|add|update|delete|updateStatus|import|sync} [options]")
// 		os.Exit(1)
// 	}

// 	command := args[0]
// 	commandArgs := args[1:]

// 	db := hrUtil.OpenDatabase(*databasePath)
// 	hrController := controller.NewHRController(db)

// 	switch command {
// 	case "list":
// 		listStudents(hrController, commandArgs)
// 	case "update":
// 		updateStudent(hrController, commandArgs)
// 	case "add":
// 		addStudent(hrController, commandArgs)
// 	case "delete":
// 		deleteStudent(hrController, commandArgs)
// 	case "updateStatus":
// 		updateStudentStatus(hrController, commandArgs)
// 	case "import":
// 		importStudents(hrController, commandArgs)
// 	case "sync":
// 		synchronizeStudents(hrController)
// 	default:
// 		fmt.Printf("Unknown command: %s\n", command)
// 		fmt.Println("Available commands: list, add, update, delete, updateStatus, import, sync")
// 		os.Exit(1)
// 	}
// }

// // Functions now only handle CLI concerns like argument parsing and output formatting

// func listStudents(hrController *controller.HRController, args []string) {
// 	fs := flag.NewFlagSet("list", flag.ExitOnError)
// 	fs.Parse(args)

// 	// Business logic moved to controller
// 	studentInfos, err := hrController.ListStudents()
// 	if err != nil {
// 		fmt.Printf("Error listing students: %v\n", err)
// 		os.Exit(1)
// 	}

// 	// CLI-specific output formatting
// 	fmt.Println("Human Resource Student Info:")
// 	for _, s := range studentInfos {
// 		fmt.Printf("SID: %s | Name: %s %s | Gender: %s | CitizenID: %s | Phone: %s | Status: %s | Email: %s\n",
// 			s.StudentCode, s.FirstName, s.LastName, s.Gender, s.CitizenID, s.PhoneNumber,
// 			hrUtil.StatusToString(*s.Status), s.Email)
// 	}
// }

// func updateStudent(hrController *controller.HRController, args []string) {
// 	fs := flag.NewFlagSet("update", flag.ExitOnError)
// 	studentID := fs.String("id", "", "Student ID to update")
// 	firstName := fs.String("fname", "", "New First Name value")
// 	lastName := fs.String("lname", "", "New Last Name value")
// 	gender := fs.String("gender", "", "New Gender value")
// 	citizenID := fs.String("citizenID", "", "New Citizen ID value")
// 	phoneNumber := fs.String("phone", "", "New Phone Number value")
// 	emailStudent := fs.String("email", "", "New Email value")
// 	fs.Parse(args)

// 	hrUtil.ValidateFlags(fs, []string{"id"})

// 	// Business logic moved to controller
// 	if err := hrController.UpdateStudent(*studentID, *firstName, *lastName, *gender,
// 		*citizenID, *phoneNumber, *emailStudent); err != nil {
// 		fmt.Printf("Failed to update student info: %v\n", err)
// 		os.Exit(1)
// 	}

// 	fmt.Println("Student updated successfully!")
// }

// func addStudent(hrController *controller.HRController, args []string) {
// 	fs := flag.NewFlagSet("add", flag.ExitOnError)
// 	studentID := fs.String("id", "", "Student ID")
// 	firstName := fs.String("fname", "", "First Name")
// 	lastName := fs.String("lname", "", "Last Name")
// 	gender := fs.String("gender", "", "Gender")
// 	citizenID := fs.String("citizenID", "", "Citizen ID")
// 	phoneNumber := fs.String("phone", "", "Phone Number")
// 	fs.Parse(args)

// 	hrUtil.ValidateFlags(fs, []string{"id", "fname", "lname"})

// 	// Business logic moved to controller
// 	if err := hrController.CreateStudent(*studentID, *firstName, *lastName, *gender, *citizenID, *phoneNumber); err != nil {
// 		fmt.Printf("Failed to add student info: %v\n", err)
// 		os.Exit(1)
// 	}

// 	fmt.Println("Student added successfully!")
// }

// func deleteStudent(hrController *controller.HRController, args []string) {
// 	fs := flag.NewFlagSet("delete", flag.ExitOnError)
// 	studentID := fs.String("id", "", "Student ID to delete")
// 	fs.Parse(args)

// 	hrUtil.ValidateFlags(fs, []string{"id"})

// 	// Business logic in controller
// 	if err := hrController.StudentController.Delete(*studentID); err != nil {
// 		fmt.Printf("Failed to delete student info: %v\n", err)
// 		os.Exit(1)
// 	}

// 	fmt.Println("Student deleted successfully!")
// }

// func updateStudentStatus(hrController *controller.HRController, args []string) {
// 	fs := flag.NewFlagSet("updateStatus", flag.ExitOnError)
// 	studentID := fs.String("id", "", "Student ID to update status")
// 	status := fs.String("status", "", "New Status (ACTIVE, GRADUATED, or DROP)")
// 	fs.Parse(args)

// 	hrUtil.ValidateFlags(fs, []string{"id", "status"})

// 	newStatus, err := hrUtil.StatusFromString(*status)
// 	if err != nil {
// 		fmt.Printf("Error: %v\n", err)
// 		os.Exit(1)
// 	}

// 	// Business logic in controller
// 	if err := hrController.StudentController.UpdateStatus(*studentID, newStatus); err != nil {
// 		fmt.Printf("Failed to update student status: %v\n", err)
// 		os.Exit(1)
// 	}

// 	fmt.Printf("Student %s status successfully updated to %s!\n", *studentID, *status)
// }

// func importStudents(hrController *controller.HRController, args []string) {
// 	fs := flag.NewFlagSet("import", flag.ExitOnError)
// 	filePath := fs.String("path", "", "Path to CSV or JSON for HR student info")
// 	fs.Parse(args)

// 	if *filePath == "" {
// 		fmt.Println("Error: File path for HR student data is required.")
// 		fmt.Println("Usage: go run humanresourcecli.go [-database=<path>] import -path=<path>")
// 		os.Exit(1)
// 	}

// 	// Business logic moved to controller
// 	if err := hrController.ImportStudents(*filePath); err != nil {
// 		fmt.Printf("Import failed: %v\n", err)
// 		os.Exit(1)
// 	}

// 	fmt.Println("Students imported successfully!")
// }

// func synchronizeStudents(hrController *controller.HRController) {
// 	// Business logic in controller
// 	if err := hrController.SynchronizeStudents(); err != nil {
// 		fmt.Printf("Failed to synchronize students: %v\n", err)
// 		os.Exit(1)
// 	}

// 	fmt.Println("Students synchronized successfully!")
// }
