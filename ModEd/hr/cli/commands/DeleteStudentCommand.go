package commands

import (
    "flag"
    "fmt"
    "os"
    "ModEd/hr/controller"
    hrUtil "ModEd/hr/util"

)
// usage : go run hr/cli/HumanResourceCLI.go delete -field="value"
// required field : id !!"
func (c *DeleteStudentCommand) Run(args []string) {
    fs := flag.NewFlagSet("delete", flag.ExitOnError)
    studentID := fs.String("id", "", "Student ID to delete")
    fs.Parse(args)

    if err := hrUtil.ValidateRequiredFlags(fs, []string{"id"}); err != nil {
        fmt.Printf("Validation error: %v\n", err)
        fs.Usage()
        os.Exit(1)
    }

    db := hrUtil.OpenDatabase(*hrUtil.DatabasePath)

	hrFacade := controller.NewHRFacade(db)
	if err := hrFacade.DeleteStudent(*studentID); err != nil {
    	fmt.Printf("Failed to delete student info: %v\n", err)
    	os.Exit(1)
}

    fmt.Println("Student deleted successfully!")
}