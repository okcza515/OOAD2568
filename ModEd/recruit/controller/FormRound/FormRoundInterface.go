package form

import (
	"ModEd/recruit/model"
	"fmt"
)

type FormRound interface {
	ApplyForm(applicant *model.Applicant) error
}

var strategyRegistry = map[string]FormRound{
	"Portfolio":   &PortfolioStrategy{},
	"Scholarship": &ScholarshipStrategy{},
}

func GetFormStrategy(roundName string) (FormRound, error) {
	strategy, exists := strategyRegistry[roundName]
	if !exists {
		return nil, fmt.Errorf("no strategy registered for now: %s", roundName)
	}
	return strategy, nil
}
