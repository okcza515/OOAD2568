package commands

import (
    "flag"
    "fmt"

    "ModEd/hr/controller"
    hrUtil "ModEd/hr/util"

    "gorm.io/gorm"
)

// usage : go run hr/cli/HumanResourceCLI.go delete -field="value"
// required field : id !!
func (c *DeleteStudentCommand) Execute(args []string, tx *gorm.DB) error {
    fs := flag.NewFlagSet("delete", flag.ExitOnError)
    studentID := fs.String("id", "", "Student ID to delete")
    fs.Parse(args)

    if err := hrUtil.ValidateRequiredFlags(fs, []string{"id"}); err != nil {
        fs.Usage()
        return fmt.Errorf("validation error: %v", err)
    }

    db := hrUtil.OpenDatabase(*hrUtil.DatabasePath)

    // Create a TransactionManager instance.
    tm := &hrUtil.TransactionManager{DB: db}

    err := tm.Execute(func(tx *gorm.DB) error {
        // Use StudentHRController to handle the deletion business logic.
        return controller.DeleteStudent(tx, *studentID)
    })

    if err != nil {
        return fmt.Errorf("failed to delete student: %v", err)
    }

    fmt.Println("Student deleted successfully!")
    return nil
}