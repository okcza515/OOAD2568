package model

import (
	master "ModEd/common/model"
	curriculum "ModEd/curriculum/model"
	"time"
)

type PermanentSchedule struct {
	StartDate   time.Time
	EndDate     time.Time
	IsAvailable bool
	Faculty     master.Faculty
	Department  master.Department
	ProgramType master.ProgramType
	Classroom   string
	Course      curriculum.Course
	Class       curriculum.Class
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
