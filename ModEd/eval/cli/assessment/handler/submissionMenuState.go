package handler

import (
	"fmt"

	"ModEd/eval/model"
)

// SubmissionMenuState represents the submission management menu
type SubmissionMenuState struct {
	*BaseMenuState
	params *AssessmentCLIParams
}

// NewSubmissionMenuState creates a new submission menu state
func NewSubmissionMenuState(params *AssessmentCLIParams) *SubmissionMenuState {
	state := &SubmissionMenuState{
		BaseMenuState: NewBaseMenuState("Submission Management", nil),
		params:        params,
	}

	// Define menu items using AddMenuItem
	state.AddMenuItem("1", "List Submissions", state.listSubmissions)
	state.AddMenuItem("2", "Record Submission", state.recordSubmission)
	state.AddMenuItem("3", "Update Submission", state.updateSubmission)
	state.AddMenuItem("4", "Delete Submission", state.deleteSubmission)
	state.AddBackItem() // Add back option
	return state
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
	return s.BaseMenuState.HandleInput(input)
}

func (s *SubmissionMenuState) listSubmissions() (MenuState, error) {
	var submissions []*model.AnswerSubmission
	var err error

	submissions, err = s.params.SubmissionController.GetAllSubmissions()

	if err != nil {
		fmt.Printf("Error fetching submissions: %v\n", err)
		return s, nil
	}

	fmt.Println("\n===== Submissions =====")
	for i, sub := range submissions {
		fmt.Printf("%d. Student ID: %d, Exam ID: %d\n", i+1, sub.StudentID, sub.ExamID)
		fmt.Printf("   Score: %.2f\n", sub.Score)
		fmt.Println("   ---------------------")
	}

	return s, nil
}

func (s *SubmissionMenuState) recordSubmission() (MenuState, error) {
	var studentID, examID uint
	var content string

	fmt.Print("Enter Student ID: ")
	fmt.Scanln(&studentID)

	fmt.Print("Enter Exam ID: ")
	fmt.Scanln(&examID)

	fmt.Print("Enter Content/URL: ")
	fmt.Scanln(&content)

	submission := &model.AnswerSubmission{
		StudentID: studentID,
		ExamID:    examID,
		Score:     0.0, // Default score
	}

	_, err := s.params.SubmissionController.CreateSubmission(submission)
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

	submission, err := s.params.SubmissionController.GetSubmission(id)
	if err != nil {
		fmt.Printf("Error fetching submission: %v\n", err)
		return s, nil
	}

	var score float64

	fmt.Printf("Enter new score (current: %.2f): ", submission.Score)
	fmt.Scanln(&score)
	if score != 0 {
		submission.Score = score
	}

	_, err = s.params.SubmissionController.UpdateSubmission(submission)
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

	_, err := s.params.SubmissionController.DeleteSubmission(id)
	if err != nil {
		fmt.Printf("Error deleting submission: %v\n", err)
	} else {
		fmt.Println("Submission deleted successfully!")
	}

	return s, nil
}
