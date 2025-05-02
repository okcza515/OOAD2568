package handler

import (
	"ModEd/core/cli"
	"ModEd/curriculum/controller"
	"ModEd/curriculum/utils"
	"fmt"
)

type InternshipEvaluationHandler struct {
	manager *cli.CLIMenuStateManager
	wrapper *controller.InternshipModuleWrapper
}

func NewIInternshipEvaluationHandler(manager *cli.CLIMenuStateManager, wrapper *controller.InternshipModuleWrapper) *InternshipEvaluationHandler {
	return &InternshipEvaluationHandler{
		manager: manager,
		wrapper: wrapper,
	}
}

func (handler *InternshipEvaluationHandler) Render() {
	fmt.Println("\n==== Internship Evaluation System ====")
	fmt.Println("1. Review Evaluation")
	fmt.Println("2. Report Evaluation")
	fmt.Println("3. Search Review and Report Scores")
	fmt.Println("Type 'exit' to quit")
	fmt.Print("Enter your choice: ")
}

func (handler *InternshipEvaluationHandler) HandleInput(input string) error {
	switch input {
	case "1":
		if err := handler.Review_Evaluation(); err != nil {
			return fmt.Errorf("failed to review evaluation: %w", err)
		}
	case "2":
		if err := handler.Report_Evaluation(); err != nil {
			return fmt.Errorf("failed to report evaluation: %w", err)
		}
	case "3":
		if err := handler.Search_Review_Report_Scores(); err != nil {
			return fmt.Errorf("failed to search review and report scores: %w", err)
		}
	case "exit":
		fmt.Println("Exiting Internship Evaluation System.")
		return nil
	default:
		fmt.Println("Invalid choice. Please try again.")
	}
	return nil
}

func (handler *InternshipEvaluationHandler) Review_Evaluation() error {

	id := utils.GetUserInput("Enter the ID of the internship application to delete: ")

	application, err := handler.wrapper.InternshipApplication.GetInternshipApplicationByID(id)
	if err != nil {
		return fmt.Errorf("failed to retrieve internship application: %w", err)
	}

	InstructorScore := int(utils.GetUserInputUint("Enter the review score (0-100): "))
	application.SupervisorReview.InstructorScore = InstructorScore
	MentorScore := int(utils.GetUserInputUint("Enter the review score (0-100): "))
	application.SupervisorReview.MentorScore = MentorScore
	err = handler.wrapper.InternshipApplication.UpdateInternshipApplication(application)
	if err != nil {
		return fmt.Errorf("failed to update internship application: %w", err)
	}
	fmt.Println("Internship application updated successfully.")

	return nil
}

func (handler *InternshipEvaluationHandler) Report_Evaluation() error {
	id := utils.GetUserInput("Enter the ID of the internship application to delete: ")

	application, err := handler.wrapper.InternshipApplication.GetInternshipApplicationByID(id)
	if err != nil {
		return fmt.Errorf("failed to retrieve internship application: %w", err)
	}

	ReportScore := int(utils.GetUserInputUint("Enter the report score (0-100): "))
	application.InternshipReport.ReportScore = ReportScore
	err = handler.wrapper.InternshipApplication.UpdateInternshipApplication(application)
	if err != nil {
		return fmt.Errorf("failed to update internship application: %w", err)
	}
	fmt.Println("Internship application updated successfully.")

	return nil
}

func (handler *InternshipEvaluationHandler) Search_Review_Report_Scores() error {
	id := utils.GetUserInput("Enter the ID of the internship application to delete: ")

	application, err := handler.wrapper.InternshipApplication.GetInternshipApplicationByID(id)
	if err != nil {
		return fmt.Errorf("failed to retrieve internship application: %w", err)
	}

	fmt.Printf("Internship Application ID: %s\n", application.StudentCode)
	fmt.Printf("Company Name: %s\n", application.Company.CompanyName)
	fmt.Printf("Review Scores:\n")
	fmt.Printf("Instructor Score: %d\n", application.SupervisorReview.InstructorScore)
	fmt.Printf("Mentor Score: %d\n", application.SupervisorReview.MentorScore)
	fmt.Printf("Report Scores:\n")
	fmt.Printf("Report Score: %d\n", application.InternshipReport.ReportScore)

	return nil
}
