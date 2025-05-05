package controller

import "ModEd/recruit/model"

type FilterStrategy[T any] interface {
	Filter(data []T) ([]T, error)
}
type InterviewFilterByInstructorID struct {
	InstructorID uint
}

func (f *InterviewFilterByInstructorID) Filter(interviews []model.Interview) ([]model.Interview, error) {
	var filtered []model.Interview
	for _, interview := range interviews {
		if interview.InstructorID == f.InstructorID {
			filtered = append(filtered, interview)
		}
	}
	return filtered, nil
}

type InterviewFilterByStatus struct {
	Status string
}

func (f *InterviewFilterByStatus) Filter(interviews []model.Interview) ([]model.Interview, error) {
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

type ApplicationReportFilterByID struct {
	ApplicationReportID uint
}

func (f *ApplicationReportFilterByID) Filter(data []model.ApplicationReport) ([]model.ApplicationReport, error) {
	var filtered []model.ApplicationReport
	for _, report := range data {
		if report.ApplicationReportID == f.ApplicationReportID {
			filtered = append(filtered, report)
		}
	}
	return filtered, nil
}
