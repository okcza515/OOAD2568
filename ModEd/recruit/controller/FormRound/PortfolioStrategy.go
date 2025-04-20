// MEP-1003 Student Recruitment
package form

import (
	"ModEd/recruit/model"
	"fmt"
)

type PortfolioStrategy struct{}

func (p *PortfolioStrategy) ApplyForm(applicant *model.Applicant) error {
	fmt.Print("Enter Portfolio URL: ")
	var portfolioURL string
	fmt.Scan(&portfolioURL)
	applicant.PortfolioURL = portfolioURL
	return nil
}
