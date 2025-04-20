package commands

// import (
// 	"ModEd/hr/controller"
// 	hrModel "ModEd/hr/model"
// 	"ModEd/hr/util"
// 	"flag"
// 	"fmt"


// 	"gorm.io/gorm"
// )

// // usage: go run hr/cli/HumanResourceCLI.go requestResignation -id="66050001" -reason="ย้ายคณะ"
// func (c *RequestResignationStudentCommand) Execute(args []string, tx *gorm.DB) error {
// 	fs := flag.NewFlagSet("requestResignation", flag.ExitOnError)
// 	ID := fs.String("id", "", "ID")
// 	reason := fs.String("reason", "", "Reason for resignation")
// 	fs.Parse(args)

// 	if err := util.ValidateRequiredFlags(fs, []string{"id", "reason", "role"}); err != nil {
// 		fs.Usage()
// 		return fmt.Errorf("Validation error: %v\n", err)
// 	}

// 	db := util.OpenDatabase(*util.DatabasePath)
// 	hrFacade := controller.NewHRFacade(db)

// 	request := hrModel.NewRequestResignationStudentBuilder().
// 		WithStudentID(*ID).
// 		WithReason(*reason).
// 		Build()

// 	if err := hrFacade.SubmitResignationStudentRequest(request); err != nil {
// 		return fmt.Errorf("Failed to submit resignation request: %v\n", err)
// 	}

// 	fmt.Println("Resignation request submitted successfully.")
// 	return nil
// }
