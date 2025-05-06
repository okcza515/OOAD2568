package controller

import (
	"ModEd/eval/model"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type AssessmentController struct {
	db *gorm.DB
}

func NewAssessmentController(db *gorm.DB) *AssessmentController {
	return &AssessmentController{db: db}
}

// CreateAssessment creates a new assessment
func (c *AssessmentController) CreateAssessment(ctx *gin.Context) {
	var assessment model.Assessment
	if err := ctx.ShouldBindJSON(&assessment); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Use AssessmentBuilder
	builder := model.NewAssessmentBuilder(assessment.Type)
	newAssessment := builder.
		SetTitle(assessment.Title).
		SetDescription(assessment.Description).
		SetDates(assessment.StartDate, assessment.DueDate).
		SetCourse(assessment.CourseId).
		SetInstructor(assessment.InstructorCode).
		Build()

	if err := c.db.Create(newAssessment).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, newAssessment)
}

// GetAssessment retrieves an assessment by ID
func (c *AssessmentController) GetAssessment(ctx *gin.Context) {
	id := ctx.Param("id")
	var assessment model.Assessment

	if err := c.db.First(&assessment, id).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Assessment not found"})
		return
	}

	ctx.JSON(http.StatusOK, assessment)
}

// UpdateAssessment updates an existing assessment
func (c *AssessmentController) UpdateAssessment(ctx *gin.Context) {
	id := ctx.Param("id")
	var assessment model.Assessment

	if err := c.db.First(&assessment, id).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Assessment not found"})
		return
	}

	if err := ctx.ShouldBindJSON(&assessment); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.db.Save(&assessment).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, assessment)
}

// DeleteAssessment deletes an assessment
func (c *AssessmentController) DeleteAssessment(ctx *gin.Context) {
	id := ctx.Param("id")
	var assessment model.Assessment

	if err := c.db.First(&assessment, id).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Assessment not found"})
		return
	}

	if err := c.db.Delete(&assessment).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Assessment deleted successfully"})
}

// ListAssessments lists all assessments with optional filters
func (c *AssessmentController) ListAssessments(ctx *gin.Context) {
	var assessments []model.Assessment
	query := c.db.Model(&model.Assessment{})

	// Apply filters
	if courseId := ctx.Query("course_id"); courseId != "" {
		query = query.Where("course_id = ?", courseId)
	}
	if instructorId := ctx.Query("instructor_id"); instructorId != "" {
		query = query.Where("instructor_code = ?", instructorId)
	}
	if assessmentType := ctx.Query("type"); assessmentType != "" {
		query = query.Where("type = ?", assessmentType)
	}

	if err := query.Find(&assessments).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, assessments)
}

// UpdateAssessmentStatus updates the status of an assessment
func (c *AssessmentController) UpdateAssessmentStatus(ctx *gin.Context) {
	id := ctx.Param("id")
	var assessment model.Assessment

	if err := c.db.First(&assessment, id).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Assessment not found"})
		return
	}

	var statusUpdate struct {
		Status model.AssessmentStatus `json:"status"`
	}

	if err := ctx.ShouldBindJSON(&statusUpdate); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	assessment.SetStatus(statusUpdate.Status)

	if err := c.db.Save(&assessment).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, assessment)
}

// SubmitAssessment handles student submissions
func (c *AssessmentController) SubmitAssessment(ctx *gin.Context) {
	assessmentId := ctx.Param("id")
	var submission model.AssessmentSubmission

	if err := ctx.ShouldBindJSON(&submission); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get the assessment to determine the type
	var assessment model.Assessment
	if err := c.db.First(&assessment, assessmentId).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Assessment not found"})
		return
	}

	// Use appropriate strategy based on assessment type
	var strategy model.SubmissionStrategy
	if assessment.Type == model.QuizType {
		strategy = &model.QuizSubmissionStrategy{}
	} else {
		strategy = &model.AssignmentSubmissionStrategy{}
	}

	// Validate and process submission
	if err := strategy.ValidateSubmission(&submission); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	submission.Submitted = true
	submission.SubmittedAt = time.Now()

	if err := c.db.Create(&submission).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, submission)
}

