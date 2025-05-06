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
	AssessmentID   uint
	AssessmentType string // "assignment" or "quiz"
	Score          uint
	Comment        string
	EvaluatedAt    time.Time
}

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
			continue
		}

		score, _ := strconv.Atoi(record[4])
		assessmentID, _ := strconv.ParseUint(record[2], 10, 32)
		assessmentType := record[3]

		evaluatedAt, _ := time.Parse(time.RFC3339, record[6])

		evaluations = append(evaluations, &Evaluation{
			StudentCode:    record[0],
			InstructorCode: record[1],
			AssessmentID:   uint(assessmentID),
			AssessmentType: assessmentType,
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

	writer.Write([]string{"student_code", "instructor_code", "assessment_id", "assessment_type", "score", "comment", "evaluated_at"})

	for _, e := range evaluations {
		writer.Write([]string{
			e.StudentCode,
			e.InstructorCode,
			strconv.FormatUint(uint64(e.AssessmentID), 10),
			e.AssessmentType,
			strconv.FormatUint(uint64(e.Score), 10),
			e.Comment,
			e.EvaluatedAt.Format(time.RFC3339),
		})
	}

	return nil
}
