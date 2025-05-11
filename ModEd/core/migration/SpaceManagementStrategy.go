// MEP-1013
package migration

import (
	"ModEd/asset/model"
)

type SpaceManagementMigrationStrategy struct {
}

func (s *SpaceManagementMigrationStrategy) GetModels() []interface{} {
	return []interface{}{
		&model.InstrumentManagement{},
		&model.SupplyManagement{},
		&model.Booking{},
		&model.TimeTable{},
		&model.PermanentSchedule{},
		&model.Room{},
	}
}
