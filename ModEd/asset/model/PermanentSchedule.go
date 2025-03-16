package model

import (
	"ModEd/common/model"
	"time"
)

type PermanentSchedule struct {
	StartDate   time.Time
	EndDate     time.Time
	IsAvailable bool
	Faculty     model.Faculty
	Department  model.Department
	ProgramType model.ProgramType
	Classroom   string
	Course      string
	Section     int
	Class       string
}

/*func (ps *PermanentSchedule) SetStarttime(starttime time.Time) {
	if ps.Starttime.IsZero() {
		ps.Starttime = starttime
	}
}

func (ps *PermanentSchedule) SetEndtime(endtime time.Time) {
	if ps.Endtime.IsZero() {
		ps.Endtime = endtime
	}
}*/
