// MEP-1003 Student Recruitment
package controller

import (
	"fmt"
)

type ScholarshipStrategy struct{}

func (s ScholarshipStrategy) GetForm() []string {
	return []string{
		"Family Yearly Income",
	}
}

func (s ScholarshipStrategy) Validate(data map[string]string) (error) {
	for _, forms := range s.GetForm() {
		roundData := data[forms]
		if roundData == "" {
			return fmt.Errorf("missing data for %s", forms)
		}
	}
	return nil
}
