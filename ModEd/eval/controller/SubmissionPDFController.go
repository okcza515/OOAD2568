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
	db        *gorm.DB
	uploadDir string
}

func NewSubmissionPDFController(db *gorm.DB, uploadDir string) *SubmissionPDFController {
	if uploadDir == "" {
		uploadDir = "uploads/assessments"
	}

	if err := os.MkdirAll(uploadDir, 0755); err != nil {
	}

	return &SubmissionPDFController{
		db:        db,
		uploadDir: uploadDir,
	}
}

func (c *SubmissionPDFController) GetUploadDir() string {
	return c.uploadDir
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

	assessmentDir := filepath.Join(c.uploadDir, strconv.FormatUint(uint64(assessmentID), 10))
	if err := os.MkdirAll(assessmentDir, 0755); err != nil {
		return nil, errors.New("failed to create upload directory")
	}

	filePath := filepath.Join(assessmentDir, filename)

	src, err := file.Open()
	if err != nil {
		return nil, errors.New("failed to open uploaded file")
	}
	defer src.Close()

	dst, err := os.Create(filePath)
	if err != nil {
		return nil, errors.New("failed to create destination file")
	}
	defer dst.Close()

	if _, err = io.Copy(dst, src); err != nil {
		return nil, errors.New("failed to save file")
	}

	pathFile := &model.PathFile{
		Path:     filePath,
		Filename: file.Filename,
		MimeType: "application/pdf",
		Size:     file.Size,
	}

	return pathFile, nil
}

func (c *SubmissionPDFController) ReadPDF(filePath string) (*os.File, error) {
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return nil, errors.New("PDF file does not exist")
	}

	file, err := os.Open(filePath)
	if err != nil {
		return nil, errors.New("failed to open PDF file")
	}

	return file, nil
}

func (c *SubmissionPDFController) UpdatePDF(file *multipart.FileHeader, oldFilePath string, assessmentID uint, studentCode string) (*model.PathFile, error) {
	if oldFilePath != "" {
		if _, err := os.Stat(oldFilePath); !os.IsNotExist(err) {
			if err := os.Remove(oldFilePath); err != nil {
				return nil, errors.New("failed to delete old PDF file")
			}
		}
	}

	return c.SavePDF(file, assessmentID, studentCode)
}

func (c *SubmissionPDFController) DeletePDF(filePath string) error {
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return errors.New("PDF file does not exist")
	}

	if err := os.Remove(filePath); err != nil {
		return errors.New("failed to delete PDF file")
	}

	return nil
}
