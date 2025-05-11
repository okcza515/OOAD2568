package commands

import (
	"ModEd/hr/util"
	"flag"
	"fmt"
	"strings"

	"gorm.io/gorm"
)

type ControllerReviewFunc func(requestID string, action string, reason string) error

func handleReviewCommand(
	args []string,
	tx *gorm.DB,
	commandName string,
	idDescription string,
	controllerFunc ControllerReviewFunc,
	entityType string,
) error {
	fs := flag.NewFlagSet(commandName, flag.ExitOnError)
	requestID := fs.String("id", "", idDescription)
	action := fs.String("action", "", "Action to perform (approve or reject)")
	reason := fs.String("reason", "", "Reason if the request is rejected")
	fs.Parse(args)

	validator := util.NewValidationChain(fs)
	validator.Field("id").Required()
	validator.Field("action").Required().AllowedValues([]string{"approve", "reject"})
	if strings.ToLower(*action) == "reject" {
		validator.Field("reason").Required()
	}
	if err := validator.Validate(); err != nil {
		fs.Usage()
		return fmt.Errorf("validation error: %v", err)
	}

	if err := controllerFunc(*requestID, strings.ToLower(*action), *reason); err != nil {
		return fmt.Errorf("failed to review %s request: %v", entityType, err)
	}

	fmt.Printf("%s request '%s' %sed successfully!\n", entityType, *requestID, strings.ToLower(*action))
	return nil
}
