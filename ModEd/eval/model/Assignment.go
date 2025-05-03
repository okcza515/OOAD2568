package model

import (
	"fmt"
	"time"

	commonModel "ModEd/common/model"
	curriculumModel "ModEd/curriculum/model"

	"gorm.io/gorm"
)

// Assignment represents an assignment assessment
type Assignment struct {
	BaseAssessment `gorm:"embedded"`
	MaxFileSize    int64
	AllowedTypes   []string
	GroupSize      int
	IsGroup        bool
	Attachments    []Attachment
}

// Attachment represents a file attachment for an assignment
type Attachment struct {
	gorm.Model
	AssignmentID uint
	FileName     string
	FileSize     int64
	FileType     string
	UploadedAt   time.Time
	UploadedBy   commonModel.Instructor
}

// NewAssignment creates a new assignment with default values
func NewAssignment() *Assignment {
	return &Assignment{
		BaseAssessment: BaseAssessment{
			Type:   AssessmentTypeAssignment,
			Status: AssessmentStatusDraft,
		},
		MaxFileSize: 10 * 1024 * 1024, // 10MB default
		GroupSize:   1,
		IsGroup:     false,
	}
}

// Validate implements additional assignment-specific validation
func (a *Assignment) Validate() error {
	if err := a.BaseAssessment.Validate(); err != nil {
		return err
	}
	if a.MaxFileSize < 0 {
		return fmt.Errorf("max file size cannot be negative")
	}
	if a.GroupSize < 1 {
		return fmt.Errorf("group size must be at least 1")
	}
	return nil
}

// AddAttachment adds a file attachment to the assignment
func (a *Assignment) AddAttachment(attachment Attachment) {
	a.Attachments = append(a.Attachments, attachment)
}

// RemoveAttachment removes a file attachment from the assignment
func (a *Assignment) RemoveAttachment(attachmentID uint) {
	for i, attachment := range a.Attachments {
		if attachment.ID == attachmentID {
			a.Attachments = append(a.Attachments[:i], a.Attachments[i+1:]...)
			break
		}
	}
}

// GetAttachments returns all attachments for the assignment
func (a *Assignment) GetAttachments() []Attachment {
	return a.Attachments
}

// GetMaxFileSize returns the maximum allowed file size
func (a *Assignment) GetMaxFileSize() int64 {
	return a.MaxFileSize
}

// GetAllowedTypes returns the allowed file types
func (a *Assignment) GetAllowedTypes() []string {
	return a.AllowedTypes
}

// GetGroupSize returns the group size
func (a *Assignment) GetGroupSize() int {
	return a.GroupSize
}

// IsGroupAssignment returns whether this is a group assignment
func (a *Assignment) IsGroupAssignment() bool {
	return a.IsGroup
}
