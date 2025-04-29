package controller

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"time"

	assignmentModel "ModEd/eval/model"

	"gorm.io/gorm"
)

type AssignmentController struct {
	csvPath     string
	assignments []assignmentModel.Assignment
}

func NewAssignmentController(csvPath string) (*AssignmentController, error) {
	controller := &AssignmentController{
		csvPath: csvPath,
	}

	// Load existing assignments from CSV
	err := controller.loadFromCSV()
	if err != nil {
		return nil, err
	}

	return controller, nil
}

func (c *AssignmentController) loadFromCSV() error {
	file, err := os.Open(c.csvPath)
	if err != nil {
		return err
	}
	defer file.Close()

	reader := csv.NewReader(file)

	// Skip header
	_, err = reader.Read()
	if err != nil {
		return err
	}

	c.assignments = []assignmentModel.Assignment{}

	// Read all records
	for {
		record, err := reader.Read()
		if err != nil {
			break
		}

		id, _ := strconv.ParseUint(record[0], 10, 32)
		released, _ := strconv.ParseBool(record[3])
		startDate, _ := time.Parse("2006-01-02", record[4])
		dueDate, _ := time.Parse("2006-01-02", record[5])

		assignment := assignmentModel.Assignment{
			Model:       gorm.Model{ID: uint(id)},
			Title:       record[1],
			Description: record[2],
			Released:    released,
			StartDate:   startDate,
			DueDate:     dueDate,
			Status:      record[6],
		}

		c.assignments = append(c.assignments, assignment)
	}

	return nil
}

func (c *AssignmentController) saveToCSV() error {
	file, err := os.Create(c.csvPath)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Write header
	writer.Write([]string{"ID", "Title", "Description", "Released", "StartDate", "DueDate", "Status"})

	// Write all assignments
	for _, a := range c.assignments {
		writer.Write([]string{
			strconv.FormatUint(uint64(a.ID), 10),
			a.Title,
			a.Description,
			strconv.FormatBool(a.Released),
			a.StartDate.Format("2006-01-02"),
			a.DueDate.Format("2006-01-02"),
			a.Status,
		})
	}

	return nil
}

func (c *AssignmentController) AddAssignment(assignment assignmentModel.Assignment) error {
	// Generate new ID
	maxID := uint(0)
	for _, a := range c.assignments {
		if a.ID > maxID {
			maxID = a.ID
		}
	}
	assignment.ID = maxID + 1

	c.assignments = append(c.assignments, assignment)
	return c.saveToCSV()
}

func (c *AssignmentController) DeleteAssignment(id uint) error {
	for i, a := range c.assignments {
		if a.ID == id {
			c.assignments = append(c.assignments[:i], c.assignments[i+1:]...)
			return c.saveToCSV()
		}
	}
	return fmt.Errorf("assignment not found")
}

func (c *AssignmentController) UpdateAssignment(assignment assignmentModel.Assignment) error {
	for i, a := range c.assignments {
		if a.ID == assignment.ID {
			c.assignments[i] = assignment
			return c.saveToCSV()
		}
	}
	return fmt.Errorf("assignment not found")
}

func (c *AssignmentController) GetAllAssignments() []assignmentModel.Assignment {
	return c.assignments
}

func (c *AssignmentController) GetAssignmentByID(id uint) (*assignmentModel.Assignment, error) {
	for _, a := range c.assignments {
		if a.ID == id {
			return &a, nil
		}
	}
	return nil, fmt.Errorf("assignment not found")
}

func (c *AssignmentController) GetAssignmentsByStatus(status string) []assignmentModel.Assignment {
	var result []assignmentModel.Assignment
	for _, a := range c.assignments {
		if a.Status == status {
			result = append(result, a)
		}
	}
	return result
}

func (c *AssignmentController) GetAssignmentsByDateRange(startDate, endDate time.Time) []assignmentModel.Assignment {
	var result []assignmentModel.Assignment
	for _, a := range c.assignments {
		if (a.StartDate.Equal(startDate) || a.StartDate.After(startDate)) &&
			(a.DueDate.Equal(endDate) || a.DueDate.Before(endDate)) {
			result = append(result, a)
		}
	}
	return result
}
