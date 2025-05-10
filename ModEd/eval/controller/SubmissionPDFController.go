package controller

import (
	"ModEd/eval/model"
	"errors"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"gorm.io/gorm"
)

// SubmissionPDFController handles operations for assessment submission PDF files
type SubmissionPDFController struct {
	db        *gorm.DB
	uploadDir string
}

// NewSubmissionPDFController creates a new SubmissionPDFController
func NewSubmissionPDFController(db *gorm.DB, uploadDir string) *SubmissionPDFController {
	if uploadDir == "" {
		uploadDir = "uploads/assessments"
	}

	// Create directory if it doesn't exist
	if err := os.MkdirAll(uploadDir, 0755); err != nil {
		// Log error but continue
	}

	return &SubmissionPDFController{
		db:        db,
		uploadDir: uploadDir,
	}
}

// GetUploadDir returns the upload directory
func (c *SubmissionPDFController) GetUploadDir() string {
	return c.uploadDir
}

// SavePDF saves a PDF file for a submission
func (c *SubmissionPDFController) SavePDF(file *multipart.FileHeader, assessmentID uint, studentCode string) (*model.PathFile, error) {
	// Validate file
	if file == nil {
		return nil, errors.New("no file provided")
	}

	// Check file size (10MB max)
	if file.Size > 10*1024*1024 {
		return nil, errors.New("file size exceeds maximum limit of 10MB")
	}

	// Check file type
	if !strings.HasSuffix(strings.ToLower(file.Filename), ".pdf") {
		return nil, errors.New("only PDF files are accepted")
	}

	// Generate unique filename using timestamp
	timestamp := time.Now().Unix()
	filename := studentCode + "_" + strconv.FormatUint(uint64(assessmentID), 10) + "_" + strconv.FormatInt(timestamp, 10) + ".pdf"

	// Create directory for the assessment if it doesn't exist
	assessmentDir := filepath.Join(c.uploadDir, strconv.FormatUint(uint64(assessmentID), 10))
	if err := os.MkdirAll(assessmentDir, 0755); err != nil {
		return nil, errors.New("failed to create upload directory")
	}

	// Full path for the file
	filePath := filepath.Join(assessmentDir, filename)

	// Open source file
	src, err := file.Open()
	if err != nil {
		return nil, errors.New("failed to open uploaded file")
	}
	defer src.Close()

	// Create destination file
	dst, err := os.Create(filePath)
	if err != nil {
		return nil, errors.New("failed to create destination file")
	}
	defer dst.Close()

	// Copy file content
	if _, err = io.Copy(dst, src); err != nil {
		return nil, errors.New("failed to save file")
	}

	// Create PathFile object
	pathFile := &model.PathFile{
		Path:     filePath,
		Filename: file.Filename,
		MimeType: "application/pdf",
		Size:     file.Size,
	}

	return pathFile, nil
}

// ReadPDF retrieves a PDF file
func (c *SubmissionPDFController) ReadPDF(filePath string) (*os.File, error) {
	// Check if file exists
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return nil, errors.New("PDF file does not exist")
	}

	// Open the file
	file, err := os.Open(filePath)
	if err != nil {
		return nil, errors.New("failed to open PDF file")
	}

	return file, nil
}

// UpdatePDF replaces an existing PDF file
func (c *SubmissionPDFController) UpdatePDF(file *multipart.FileHeader, oldFilePath string, assessmentID uint, studentCode string) (*model.PathFile, error) {
	// Delete old file if it exists
	if oldFilePath != "" {
		if _, err := os.Stat(oldFilePath); !os.IsNotExist(err) {
			if err := os.Remove(oldFilePath); err != nil {
				return nil, errors.New("failed to delete old PDF file")
			}
		}
	}

	// Save the new file
	return c.SavePDF(file, assessmentID, studentCode)
}

// DeletePDF removes a PDF file
func (c *SubmissionPDFController) DeletePDF(filePath string) error {
	// Check if file exists
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return errors.New("PDF file does not exist")
	}

	// Delete the file
	if err := os.Remove(filePath); err != nil {
		return errors.New("failed to delete PDF file")
	}

	return nil
}
