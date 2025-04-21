// MEP-1003 Student Recruitment
package controller

import "ModEd/recruit/model"

type Criteria interface {
	IsSatisfiedBy(applicant model.Applicant) bool
}
