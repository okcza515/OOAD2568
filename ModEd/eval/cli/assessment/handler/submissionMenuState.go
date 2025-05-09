package handler

import (
	"fmt"
	"time"

	"ModEd/eval/model"
)

// SubmissionMenuState represents the submission management menu
type SubmissionMenuState struct {
	params *AssessmentCLIParams
}

// NewSubmissionMenuState creates a new submission menu state
func NewSubmissionMenuState(params *AssessmentCLIParams) *SubmissionMenuState {
	return &SubmissionMenuState{
		params: params,
	}
}

// Enter displays the submission menu
func (s *SubmissionMenuState) Enter() error {
	fmt.Println("\n===== Submission Management =====")
	fmt.Println("1. List Submissions")
	fmt.Println("2. Record Submission")
	fmt.Println("3. Update Submission")
	fmt.Println("4. Delete Submission")
	fmt.Println("back - Return to Assessment menu")
	return nil
}

// Exit handles exit from the submission menu
func (s *SubmissionMenuState) Exit() error {
	return nil
}

// HandleInput processes user input in the submission menu
func (s *SubmissionMenuState) HandleInput(input string) (MenuState, error) {
	switch input {
	case "1":
		fmt.Println("\n===== Submissions =====")
		fmt.Println("Submission listing will be implemented here")
		return s, nil
	case "2":
		fmt.Println("\n===== Record Submission =====")
		fmt.Println("Submission recording will be implemented here")
		return s, nil
	case "3":
		fmt.Println("\n===== Update Submission =====")
		fmt.Println("Submission update will be implemented here")
		return s, nil
	case "4":
		fmt.Println("\n===== Delete Submission =====")
		fmt.Println("Submission deletion will be implemented here")
		return s, nil
	case "back":
		return NewMainMenuState(s.params), nil
	default:
		return s, fmt.Errorf("invalid choice: %s", input)
	}
}

func (s *SubmissionMenuState) listSubmissions() (MenuState, error) {
	var assessmentID uint
	fmt.Print("Enter Assessment ID (0 for all): ")
	fmt.Scanln(&assessmentID)

	var submissions []*model.Submission
	var err error

	if assessmentID > 0 {
		submissions, err = s.params.SubmissionController.GetByAssessmentID(assessmentID)
	} else {
		submissions, err = s.params.SubmissionController.GetAll()
	}

	if err != nil {
		fmt.Printf("Error fetching submissions: %v\n", err)
		return s, nil
	}

	fmt.Println("\n===== Submissions =====")
	for i, sub := range submissions {
		fmt.Printf("%d. Student: %s, Assessment ID: %d\n", i+1, sub.StudentCode, sub.AssessmentID)
		fmt.Printf("   Submitted: %s, Status: %s\n", sub.SubmittedAt.Format("2006-01-02 15:04"), sub.Status)
		fmt.Println("   ---------------------")
	}

	return s, nil
}

func (s *SubmissionMenuState) recordSubmission() (MenuState, error) {
	var assessmentID uint
	var studentCode, content string

	fmt.Print("Enter Assessment ID: ")
	fmt.Scanln(&assessmentID)

	fmt.Print("Enter Student Code: ")
	fmt.Scanln(&studentCode)

	fmt.Print("Enter Content/URL: ")
	fmt.Scanln(&content)

	submission := &model.Submission{
		AssessmentID: assessmentID,
		StudentCode:  studentCode,
		Content:      content,
		SubmittedAt:  time.Now(),
		Status:       "Submitted",
	}

	err := s.params.SubmissionController.Create(submission)
	if err != nil {
		fmt.Printf("Error recording submission: %v\n", err)
	} else {
		fmt.Println("Submission recorded successfully!")
	}

	return s, nil
}

func (s *SubmissionMenuState) updateSubmission() (MenuState, error) {
	var id uint
	fmt.Print("Enter Submission ID to update: ")
	fmt.Scanln(&id)

	submission, err := s.params.SubmissionController.GetByID(id)
	if err != nil {
		fmt.Printf("Error fetching submission: %v\n", err)
		return s, nil
	}

	var content, status string

	fmt.Printf("Enter new content (current: %s): ", submission.Content)
	fmt.Scanln(&content)
	if content != "" {
		submission.Content = content
	}

	fmt.Printf("Enter new status (current: %s): ", submission.Status)
	fmt.Scanln(&status)
	if status != "" {
		submission.Status = status
	}

	submission.SubmittedAt = time.Now()

	err = s.params.SubmissionController.Update(submission)
	if err != nil {
		fmt.Printf("Error updating submission: %v\n", err)
	} else {
		fmt.Println("Submission updated successfully!")
	}

	return s, nil
}

func (s *SubmissionMenuState) deleteSubmission() (MenuState, error) {
	var id uint
	fmt.Print("Enter Submission ID to delete: ")
	fmt.Scanln(&id)

	err := s.params.SubmissionController.Delete(id)
	if err != nil {
		fmt.Printf("Error deleting submission: %v\n", err)
	} else {
		fmt.Println("Submission deleted successfully!")
	}

	return s, nil
}
