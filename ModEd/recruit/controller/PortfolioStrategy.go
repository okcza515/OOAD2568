// MEP-1003 Student Recruitment
package controller

import (
	"fmt"
)

type PortfolioStrategy struct{}

func (s PortfolioStrategy) GetForm() []string {
	return []string{
		"Portfolio URL",
	}
}

func (s PortfolioStrategy) Validate(data map[string]string) (error) {
	for _, forms := range s.GetForm() {
		roundData := data[forms]
		if roundData == "" {
			return fmt.Errorf("missing data for %s", forms)
		}
	}
	return nil
}

