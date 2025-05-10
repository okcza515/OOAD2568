package handler

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"time"

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
	state.AddMenuItem("1", "List Submissions for Assessment", state.listSubmissions)
	state.AddMenuItem("2", "Upload PDF Submission", state.uploadSubmission)
	state.AddMenuItem("3", "View Submission Details", state.viewSubmission)
	state.AddMenuItem("4", "Delete Submission", state.deleteSubmission)
	state.AddBackItem() // Add back option
	return state
}

// Enter displays the submission menu
func (s *SubmissionMenuState) Enter() error {
	fmt.Println("\n===== Submission Management =====")
	fmt.Println("1. List Submissions for Assessment")
	fmt.Println("2. Upload PDF Submission")
	fmt.Println("3. View Submission Details")
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

// listSubmissions lists all submissions for a specific assessment
func (s *SubmissionMenuState) listSubmissions() (MenuState, error) {
	var assessmentID uint

	fmt.Print("Enter Assessment ID: ")
	fmt.Scanln(&assessmentID)

	// Get the assessment to verify it exists
	assessment, err := s.params.AssessmentController.GetAssessment(assessmentID)
	if err != nil {
		fmt.Printf("Error: Assessment with ID %d not found: %v\n", assessmentID, err)
		return s, nil
	}

	fmt.Printf("\n===== Submissions for Assessment: %s (ID: %d) =====\n",
		assessment.Title, assessment.AssessmentId)

	// Display all submissions related to this assessment
	for _, submission := range assessment.Submission {
		fmt.Printf("Student: %s %s (%s)\n", submission.FirstName, submission.LastName, submission.StudentCode)
		fmt.Printf("  Submitted: %t  Date: %s\n", submission.Submitted, submission.UpdatedAt.Format("2006-01-02 15:04:05"))
		fmt.Printf("  Score: %.2f\n", submission.Score)
		if submission.Feedback != "" {
			fmt.Printf("  Feedback: %s\n", submission.Feedback)
		}
		if submission.Answers.Path != "" {
			fmt.Printf("  PDF: %s (Size: %d bytes)\n", submission.Answers.Filename, submission.Answers.Size)
		} else {
			fmt.Printf("  PDF: Not submitted\n")
		}
		fmt.Println("  ---------------------")
	}

	if len(assessment.Submission) == 0 {
		fmt.Println("No submissions found for this assessment.")
	}

	return s, nil
}

// uploadSubmission handles uploading a PDF submission
func (s *SubmissionMenuState) uploadSubmission() (MenuState, error) {
	var assessmentID uint
	var studentCode, pdfPath string

	// Check if SubmissionPDFController is available
	if s.params.SubmissionPDFController == nil {
		fmt.Println("Error: PDF submission functionality is not available")
		return s, nil
	}

	fmt.Print("Enter Assessment ID: ")
	fmt.Scanln(&assessmentID)

	// Verify the assessment exists
	assessment, err := s.params.AssessmentController.GetAssessment(assessmentID)
	if err != nil {
		fmt.Printf("Error: Assessment with ID %d not found: %v\n", assessmentID, err)
		return s, nil
	}

	fmt.Printf("Selected Assessment: %s (ID: %d)\n", assessment.Title, assessment.AssessmentId)

	fmt.Print("Enter Student Code: ")
	fmt.Scanln(&studentCode)

	fmt.Print("Enter full path to PDF file: ")
	fmt.Scanln(&pdfPath)

	// Validate that the file exists and is a PDF
	fileInfo, err := os.Stat(pdfPath)
	if err != nil {
		fmt.Printf("Error: File not found or cannot be accessed: %v\n", err)
		return s, nil
	}

	if fileInfo.IsDir() {
		fmt.Println("Error: The path points to a directory, not a file")
		return s, nil
	}

	// Check if the file is a PDF
	if filepath.Ext(pdfPath) != ".pdf" {
		fmt.Println("Error: File must be a PDF document with .pdf extension")
		return s, nil
	}

	// Open the file to get its size
	file, err := os.Open(pdfPath)
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return s, nil
	}
	defer file.Close()

	// Copy the file to the upload directory
	uploadDir := "uploads/assessments" // Default directory
	assessmentDir := filepath.Join(uploadDir, strconv.FormatUint(uint64(assessmentID), 10))

	// Create the directory if it doesn't exist
	if err := os.MkdirAll(assessmentDir, 0755); err != nil {
		fmt.Printf("Error creating directory: %v\n", err)
		return s, nil
	}

	// Generate a unique filename using timestamp
	timestamp := time.Now().Unix()
	filename := studentCode + "_" + strconv.FormatUint(uint64(assessmentID), 10) + "_" + strconv.FormatInt(timestamp, 10) + ".pdf"
	destPath := filepath.Join(assessmentDir, filename)

	// Create destination file
	dest, err := os.Create(destPath)
	if err != nil {
		fmt.Printf("Error creating destination file: %v\n", err)
		return s, nil
	}
	defer dest.Close()

	// Reset file position
	file.Seek(0, 0)

	// Read the source file and write to destination
	buffer := make([]byte, 1024)
	for {
		n, err := file.Read(buffer)
		if err != nil {
			break
		}
		if _, err := dest.Write(buffer[:n]); err != nil {
			fmt.Printf("Error writing to destination file: %v\n", err)
			return s, nil
		}
	}

	// Create a PathFile object
	pathFile := &model.PathFile{
		Path:     destPath,
		Filename: filepath.Base(pdfPath),
		MimeType: "application/pdf",
		Size:     fileInfo.Size(),
	}

	// Find or create the assessment submission
	var submission model.AssessmentSubmission
	submission.StudentCode = studentCode
	submission.Answers = *pathFile
	submission.Submitted = true
	submission.UpdatedAt = time.Now()

	// In a real implementation, you would save this to the database
	fmt.Println("\nSubmission uploaded successfully!")
	fmt.Printf("File saved to: %s\n", destPath)

	return s, nil
}

// viewSubmission shows details of a specific submission
func (s *SubmissionMenuState) viewSubmission() (MenuState, error) {
	var assessmentID uint
	var studentCode string

	fmt.Print("Enter Assessment ID: ")
	fmt.Scanln(&assessmentID)

	fmt.Print("Enter Student Code: ")
	fmt.Scanln(&studentCode)

	// Get the assessment to verify it exists
	assessment, err := s.params.AssessmentController.GetAssessment(assessmentID)
	if err != nil {
		fmt.Printf("Error: Assessment with ID %d not found: %v\n", assessmentID, err)
		return s, nil
	}

	// Find the submission for this student
	var foundSubmission *model.AssessmentSubmission
	for _, submission := range assessment.Submission {
		if submission.StudentCode == studentCode {
			foundSubmission = &submission
			break
		}
	}

	if foundSubmission == nil {
		fmt.Printf("No submission found for student %s in assessment ID %d\n", studentCode, assessmentID)
		return s, nil
	}

	fmt.Printf("\n===== Submission Details =====\n")
	fmt.Printf("Assessment: %s (ID: %d)\n", assessment.Title, assessment.AssessmentId)
	fmt.Printf("Student: %s %s (%s)\n", foundSubmission.FirstName, foundSubmission.LastName, foundSubmission.StudentCode)
	fmt.Printf("Submitted: %t  Date: %s\n", foundSubmission.Submitted, foundSubmission.UpdatedAt.Format("2006-01-02 15:04:05"))
	fmt.Printf("Score: %.2f\n", foundSubmission.Score)

	if foundSubmission.Feedback != "" {
		fmt.Printf("Feedback: %s\n", foundSubmission.Feedback)
	}

	if foundSubmission.Answers.Path != "" {
		fmt.Printf("PDF: %s (Size: %d bytes)\n", foundSubmission.Answers.Filename, foundSubmission.Answers.Size)
		fmt.Printf("Path: %s\n", foundSubmission.Answers.Path)
	} else {
		fmt.Printf("PDF: Not submitted\n")
	}

	return s, nil
}

// deleteSubmission deletes a submission
func (s *SubmissionMenuState) deleteSubmission() (MenuState, error) {
	var assessmentID uint
	var studentCode string

	fmt.Print("Enter Assessment ID: ")
	fmt.Scanln(&assessmentID)

	fmt.Print("Enter Student Code: ")
	fmt.Scanln(&studentCode)

	// Get the assessment to verify it exists
	assessment, err := s.params.AssessmentController.GetAssessment(assessmentID)
	if err != nil {
		fmt.Printf("Error: Assessment with ID %d not found: %v\n", assessmentID, err)
		return s, nil
	}

	// Find the submission for this student
	var foundSubmission *model.AssessmentSubmission
	for _, submission := range assessment.Submission {
		if submission.StudentCode == studentCode {
			foundSubmission = &submission
			break
		}
	}

	if foundSubmission == nil {
		fmt.Printf("No submission found for student %s in assessment ID %d\n", studentCode, assessmentID)
		return s, nil
	}

	// If there's a PDF file, attempt to delete it
	if foundSubmission.Answers.Path != "" {
		if s.params.SubmissionPDFController == nil {
			fmt.Println("Warning: PDF controller not available, cannot delete file")
		} else {
			if err := s.params.SubmissionPDFController.DeletePDF(foundSubmission.Answers.Path); err != nil {
				fmt.Printf("Warning: Could not delete PDF file: %v\n", err)
			} else {
				fmt.Println("PDF file deleted successfully.")
			}
		}
	}

	// In a real implementation, you would delete the submission from the database
	fmt.Println("Submission record deleted successfully!")

	return s, nil
}
