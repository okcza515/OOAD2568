package controller

import "ModEd/recruit/model"

type FilterStrategy interface {
	Filter(interviews []model.Interview) ([]model.Interview, error)
}

type FilterByInstructorID struct {
	InstructorID uint
}

func (f *FilterByInstructorID) Filter(interviews []model.Interview) ([]model.Interview, error) {
	var filtered []model.Interview
	for _, interview := range interviews {
		if interview.InstructorID == f.InstructorID {
			filtered = append(filtered, interview)
		}
	}
	return filtered, nil
}

type FilterByStatus struct {
	Status string
}

func (f *FilterByStatus) Filter(interviews []model.Interview) ([]model.Interview, error) {
	if f.Status == "all" {
		return interviews, nil
	}

	var filtered []model.Interview
	for _, interview := range interviews {
		if string(interview.InterviewStatus) == f.Status {
			filtered = append(filtered, interview)
		}
	}
	return filtered, nil
}
