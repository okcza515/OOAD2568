package controller

// MEP-1012 Asset

import (
	"ModEd/asset/model"
	"ModEd/core"
	"ModEd/core/migration"
	"ModEd/utils/deserializer"
)

type AssetControllerFacade struct {
	BorrowInstrument BorrowInstrumentControllerInterface
	Category         CategoryControllerInterface
	Instrument       InstrumentControllerInterface
	InstrumentLog    InstrumentLogControllerInterface
	Supply           SupplyControllerInterface
	SupplyLog        SupplyLogControllerInterface
}

func NewAssetControllerFacade() (*AssetControllerFacade, error) {
	db, err := migration.GetInstance().MigrateModule(core.MODULE_ASSET).BuildDB()

	if err != nil {
		return nil, err
	}

	facade := AssetControllerFacade{}

	//facade.BorrowInstrument = &BorrowInstrumentController{db: db, BaseController: core.NewBaseController[model.BorrowInstrument]("BorrowInstrument", db)}
	//facade.Category = &CategoryController{db: db, BaseController: core.NewBaseController[model.Category]("Category", db)}
	facade.Instrument = &InstrumentController{db: db, BaseController: core.NewBaseController[model.Instrument](db)}
	//facade.InstrumentLog = &InstrumentLogController{db: db, BaseController: core.NewBaseController[model.InstrumentLog]("InstrumentLog", db)}
	facade.Supply = &SupplyController{db: db, BaseController: core.NewBaseController[model.Supply](db)}
	//facade.SupplyLog = &SupplyLogController{db: db, BaseController: core.NewBaseController("SupplyLog", db)}

	return &facade, nil
}

func (facade *AssetControllerFacade) LoadSeedData() error {
	seedData := map[string]interface{}{
		//"BorrowInstrumentList": &[]model.BorrowInstrument{},
		"Category":       &[]model.Category{},
		"InstrumentList": &[]model.Instrument{},
		//"InstrumentLog":  &[]model.InstrumentLog{},
		"SupplyList": &[]model.Supply{},
		"SupplyLog":  &[]model.SupplyLog{},
	}

	for filename, m := range seedData {
		fd, err := deserializer.NewFileDeserializer("data/asset/" + filename + ".csv")
		if err != nil {
			return err
		}

		err = fd.Deserialize(m)
		if err != nil {
			return err
		}

		result := migration.GetInstance().DB.Create(m)
		if result.Error != nil {
			return result.Error
		}
	}

	return nil
}

func (facade *AssetControllerFacade) ResetDB() error {
	err := migration.GetInstance().DropAllTables()
	if err != nil {
		return err
	}

	_, err = migration.GetInstance().MigrateModule(core.MODULE_ASSET).BuildDB()
	if err != nil {
		return err
	}

	return nil
}

func (facade *AssetControllerFacade) ResetAndLoadDB() error {
	err := facade.ResetDB()
	if err != nil {
		return err
	}

	err = facade.LoadSeedData()
	if err != nil {
		return err
	}

	return nil
}
