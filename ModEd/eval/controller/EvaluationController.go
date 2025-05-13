//Chanawat Limpanatewin 65070503445
//MEP-1006

package controller

import (
	"ModEd/core"
	evalModel "ModEd/eval/model"
	"encoding/csv"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"gorm.io/gorm"
)

type EvaluationController struct {
	*core.BaseController[evalModel.Evaluation]
	db *gorm.DB
}

func NewEvaluationController(db *gorm.DB) *EvaluationController {
	return &EvaluationController{
		db:             db,
		BaseController: core.NewBaseController[evalModel.Evaluation](db),
	}
}

func (ec *EvaluationController) CreateEvaluation(studentCode, instructorCode string, assignmentId uint, score uint, comment string) error {
	newEvaluation := evalModel.Evaluation{
		StudentCode:    studentCode,
		InstructorCode: instructorCode,
		AssignmentId:   assignmentId,
		Score:          score,
		Comment:        comment,
		EvaluatedAt:    time.Now(),
	}

	if err := ec.db.Create(&newEvaluation).Error; err != nil {
		return err
	}

	return ec.saveToCSV(&newEvaluation)
}

func (ec *EvaluationController) ViewAllEvaluations() ([]evalModel.Evaluation, error) {
	return ec.List(nil, "Student", "Instructor", "Assignment")
}

func (ec *EvaluationController) ViewEvaluationByID(studentCode string) ([]evalModel.Evaluation, error) {
	condition := map[string]interface{}{
		"student_code": studentCode,
	}
	return ec.List(condition, "Student", "Instructor", "Assignment")
}

func (ec *EvaluationController) UpdateEvaluation(id uint, score uint, comment string) error {
	evaluation := &evalModel.Evaluation{}
	if err := ec.db.First(evaluation, id).Error; err != nil {
		return err
	}

	evaluation.Score = score
	evaluation.Comment = comment
	evaluation.EvaluatedAt = time.Now()

	if err := ec.db.Save(evaluation).Error; err != nil {
		return err
	}

	return ec.updateCSV(evaluation)
}

func (ec *EvaluationController) saveToCSV(evaluation *evalModel.Evaluation) error {
	csvPath := "../../data/quiz/Evaluation.csv"

	dir := filepath.Dir(csvPath)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("failed to create directory: %v", err)
	}

	fileExists := true
	if _, err := os.Stat(csvPath); os.IsNotExist(err) {
		fileExists = false
	}

	file, err := os.OpenFile(csvPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("failed to open CSV file: %v", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	if !fileExists {
		headers := []string{"student_code", "Instructor_code", "Assignment_id", "score", "comment", "evaluated_at\n"}
		if err := writer.Write(headers); err != nil {
			return fmt.Errorf("failed to write headers: %v", err)
		}
		writer.Flush()
		if err := writer.Error(); err != nil {
			return fmt.Errorf("failed to flush headers: %v", err)
		}
	}

	record := []string{
		evaluation.StudentCode,
		evaluation.InstructorCode,
		strconv.FormatUint(uint64(evaluation.AssignmentId), 10),
		strconv.FormatUint(uint64(evaluation.Score), 10),
		evaluation.Comment,
		evaluation.EvaluatedAt.Format(time.RFC3339),
	}

	if err := writer.Write(record); err != nil {
		return fmt.Errorf("failed to write record: %v", err)
	}
	writer.Flush()
	if err := writer.Error(); err != nil {
		return fmt.Errorf("failed to flush record: %v", err)
	}

	return nil
}

func (ec *EvaluationController) updateCSV(evaluation *evalModel.Evaluation) error {
	csvPath := "../../data/quiz/Evaluation.csv"

	dir := filepath.Dir(csvPath)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("failed to create directory: %v", err)
	}

	file, err := os.OpenFile(csvPath, os.O_RDWR, 0644)
	if err != nil {
		return fmt.Errorf("failed to open CSV file: %v", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return fmt.Errorf("failed to read CSV file: %v", err)
	}

	for i, record := range records {
		if i == 0 {
			continue
		}
		if record[0] == evaluation.StudentCode {
			records[i] = []string{
				evaluation.StudentCode,
				evaluation.InstructorCode,
				strconv.FormatUint(uint64(evaluation.AssignmentId), 10),
				strconv.FormatUint(uint64(evaluation.Score), 10),
				evaluation.Comment,
				evaluation.EvaluatedAt.Format(time.RFC3339),
			}
			break
		}
	}

	file.Seek(0, 0)
	file.Truncate(0)
	writer := csv.NewWriter(file)
	if err := writer.WriteAll(records); err != nil {
		return fmt.Errorf("failed to write records: %v", err)
	}
	writer.Flush()
	if err := writer.Error(); err != nil {
		return fmt.Errorf("failed to flush records: %v", err)
	}

	return nil
}
