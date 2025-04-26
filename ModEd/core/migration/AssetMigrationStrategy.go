package migration

import (
	"ModEd/asset/model"
)

type AssetMigrationStrategy struct {
}

func (s *AssetMigrationStrategy) GetModels() []interface{} {
	return []interface{}{
		&model.InstrumentLog{},
		&model.Instrument{},
		&model.BorrowInstrument{},
		&model.Category{},
		&model.Instrument{},
		&model.InstrumentLog{},
		&model.Supply{},
		&model.SupplyLog{},
	}
}
