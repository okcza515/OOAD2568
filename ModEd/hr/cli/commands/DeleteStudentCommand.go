package commands

import (
    "flag"
    "fmt"
    "os"
    "ModEd/hr/controller"
    hrUtil "ModEd/hr/util"

)
func (c *DeleteStudentCommand) Run(args []string) {
    fs := flag.NewFlagSet("delete", flag.ExitOnError)
    studentID := fs.String("id", "", "Student ID to delete")
    fs.Parse(args)

    if *studentID == "" {
        fmt.Println("Error: Student ID is required.")
        fmt.Println("Usage: go run humanresourcecli.go [-database=<path>] delete -id=<studentID>")
        os.Exit(1)
    }

    db := hrUtil.OpenDatabase(*hrUtil.DatabasePath)
    studentController := controller.CreateStudentHRController(db)
    if err := studentController.Delete(*studentID); err != nil {
        fmt.Printf("Failed to delete student info: %v\n", err)
        os.Exit(1)
    }

    fmt.Println("Student deleted successfully!")
}