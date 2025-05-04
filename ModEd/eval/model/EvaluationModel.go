// 65070503445
// MEP-1006

package model

import (
	"encoding/csv"
	"os"
	"strconv"
	"time"
)

type Evaluation struct {
	StudentCode    string
	InstructorCode string
	AssignmentID   *uint
	QuizID         *uint
	Score          uint
	Comment        string
	EvaluatedAt    time.Time
}

// ฟังก์ชันช่วยโหลดข้อมูล Evaluation จาก CSV
func LoadEvaluationsFromCSV(filePath string) ([]*Evaluation, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	var evaluations []*Evaluation
	for i, record := range records {
		if i == 0 {
			continue // skip header
		}

		score, _ := strconv.Atoi(record[4])
		var assignmentID, quizID *uint
		if record[2] != "" {
			id, _ := strconv.ParseUint(record[2], 10, 32)
			tmp := uint(id)
			assignmentID = &tmp
		}
		if record[3] != "" {
			id, _ := strconv.ParseUint(record[3], 10, 32)
			tmp := uint(id)
			quizID = &tmp
		}

		evaluatedAt, _ := time.Parse(time.RFC3339, record[6])

		evaluations = append(evaluations, &Evaluation{
			StudentCode:    record[0],
			InstructorCode: record[1],
			AssignmentID:   assignmentID,
			QuizID:         quizID,
			Score:          uint(score),
			Comment:        record[5],
			EvaluatedAt:    evaluatedAt,
		})
	}

	return evaluations, nil
}

func SaveEvaluationsToCSV(filePath string, evaluations []*Evaluation) error {
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Write header
	writer.Write([]string{"student_code", "instructor_code", "assignment_id", "quiz_id", "score", "comment", "evaluated_at"})

	for _, e := range evaluations {
		assignmentID := ""
		quizID := ""
		if e.AssignmentID != nil {
			assignmentID = strconv.FormatUint(uint64(*e.AssignmentID), 10)
		}
		if e.QuizID != nil {
			quizID = strconv.FormatUint(uint64(*e.QuizID), 10)
		}
		writer.Write([]string{
			e.StudentCode,
			e.InstructorCode,
			assignmentID,
			quizID,
			strconv.FormatUint(uint64(e.Score), 10),
			e.Comment,
			e.EvaluatedAt.Format(time.RFC3339),
		})
	}

	return nil
}
