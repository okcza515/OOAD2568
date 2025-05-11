package controller

import (
	"errors"
)

type QuotaStrategy struct{}

func (q *QuotaStrategy) GetForm() []string {
	return []string{"MathGrade", "ScienceGrade", "EnglishGrade"}
}

func (q *QuotaStrategy) Validate(data map[string]string) error {
	requiredFields := []string{"MathGrade", "ScienceGrade", "EnglishGrade"}
	for _, field := range requiredFields {
		if data[field] == "" {
			return errors.New("missing field: " + field)
		}
	}
	return nil
}
