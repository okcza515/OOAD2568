// MEP-1003 Student Recruitment
package controller

import (
	"fmt"
)

type FormRound interface {
	GetForm() []string
	Validate(data map[string]string) (error)
}

var strategyRegistry = map[string]FormRound{
	"Portfolio":   &PortfolioStrategy{},
	"Quota":       &QuotaStrategy{},
	"Admission":   &AdmissionStrategy{},
	"Scholarship": &ScholarshipStrategy{},
}

func GetFormStrategy(roundName string) (FormRound, error) {
	strategy, exists := strategyRegistry[roundName]
	if !exists {
		return nil, fmt.Errorf("no strategy registered for now: %s", roundName)
	}
	return strategy, nil
}
