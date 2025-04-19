package criteria

import "ModEd/recruit/model"

type Criteria interface {
	IsSatisfiedBy(applicant model.Applicant) bool
}
