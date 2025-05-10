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

type SubmissionPDFController struct {
	db                   *gorm.DB
	assessmentController *AssessmentController
	uploadDir            string
}

func NewSubmissionPDFController(db *gorm.DB, assessmentController *AssessmentController, uploadDir string) *SubmissionPDFController {
	if uploadDir == "" {
		uploadDir = "uploads/assessments"
	}

	if err := os.MkdirAll(uploadDir, 0755); err != nil {
	}

	return &SubmissionPDFController{
		db:                   db,
		assessmentController: assessmentController,
		uploadDir:            uploadDir,
	}
}

func (c *SubmissionPDFController) SavePDF(file *multipart.FileHeader, assessmentID uint, studentCode string) (*model.PathFile, error) {
	if file == nil {
		return nil, errors.New("no file provided")
	}

	if file.Size > 10*1024*1024 {
		return nil, errors.New("file size exceeds maximum limit of 10MB")
	}

	if !strings.HasSuffix(strings.ToLower(file.Filename), ".pdf") {
		return nil, errors.New("only PDF files are accepted")
	}

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

// ReadPDF retrieves a PDF file by its path
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

// UpdatePDF updates an existing PDF file
func (c *SubmissionPDFController) UpdatePDF(file *multipart.FileHeader, oldFilePath string) (*model.PathFile, error) {
	// Delete old file if it exists
	if oldFilePath != "" {
		if _, err := os.Stat(oldFilePath); !os.IsNotExist(err) {
			if err := os.Remove(oldFilePath); err != nil {
				return nil, errors.New("failed to delete old PDF file")
			}
		}
	}

	// Validate new file
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

	// Use the same directory as the old file if it exists
	var dirPath string
	if oldFilePath != "" {
		dirPath = filepath.Dir(oldFilePath)
	} else {
		// Create a new directory using timestamp
		timestamp := time.Now().Unix()
		dirPath = filepath.Join(c.uploadDir, strconv.FormatInt(timestamp, 10))
		if err := os.MkdirAll(dirPath, 0755); err != nil {
			return nil, errors.New("failed to create upload directory")
		}
	}

	// Generate new filename using timestamp
	timestamp := time.Now().Unix()
	filename := "updated_" + strconv.FormatInt(timestamp, 10) + ".pdf"
	filePath := filepath.Join(dirPath, filename)

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

// DeletePDF deletes a PDF file
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
