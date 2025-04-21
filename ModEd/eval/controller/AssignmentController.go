package controller

import (
	"net/http"
	"time"

	"ModEd/assignment/model"
	"ModEd/common/model"     // สมมุติว่า commonModel อยู่ path นี้
	"ModEd/curriculum/model" // สมมุติว่า curriculumModel อยู่ path นี้

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type AssignmentInput struct {
	Title       string    `json:"title"`
	Description string    `json:"description"`
	StartDate   time.Time `json:"start_date"`
	DueDate     time.Time `json:"due_date"`
	Status      string    `json:"status"`
}

func GetAssignments(c *gin.Context) {
	var assignments []model.Assignment
	db := c.MustGet("db").(*gorm.DB)

	if err := db.Preload("Submission").Find(&assignments).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, assignments)
}

func GetAssignmentByID(c *gin.Context) {
	var assignment model.Assignment
	db := c.MustGet("db").(*gorm.DB)

	id := c.Param("id")
	if err := db.Preload("Submission").First(&assignment, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Assignment not found"})
		return
	}
	c.JSON(http.StatusOK, assignment)
}

func CreateAssignment(c *gin.Context) {
	var input AssignmentInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db := c.MustGet("db").(*gorm.DB)

	assignment := model.Assignment{
		Title:       input.Title,
		Description: input.Description,
		StartDate:   input.StartDate,
		DueDate:     input.DueDate,
		Status:      input.Status,
	}

	if err := db.Create(&assignment).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, assignment)
}

func UpdateAssignment(c *gin.Context) {
	var assignment model.Assignment
	db := c.MustGet("db").(*gorm.DB)

	id := c.Param("id")
	if err := db.First(&assignment, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Assignment not found"})
		return
	}

	var input AssignmentInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	assignment.Title = input.Title
	assignment.Description = input.Description
	assignment.StartDate = input.StartDate
	assignment.DueDate = input.DueDate
	assignment.Status = input.Status

	if err := db.Save(&assignment).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, assignment)
}

func DeleteAssignment(c *gin.Context) {
	var assignment model.Assignment
	db := c.MustGet("db").(*gorm.DB)

	id := c.Param("id")
	if err := db.First(&assignment, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Assignment not found"})
		return
	}

	if err := db.Delete(&assignment).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Assignment deleted"})
}
