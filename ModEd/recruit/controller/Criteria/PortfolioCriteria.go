// MEP-1003 Student Recruitment
package criteria

import "ModEd/recruit/model"

type PortfolioCriteria struct{}

func (c *PortfolioCriteria) IsSatisfiedBy(applicant model.Applicant) bool {
	return applicant.GPAX >= 3.0 && applicant.PortfolioURL != ""
}
