package controller

// MEP-1012 Asset

import (
	"ModEd/asset/model"
	"ModEd/core"
	"ModEd/core/migration"
	"ModEd/utils/deserializer"
	"errors"
)

type AssetControllerManager struct {
	BorrowInstrument BorrowInstrumentControllerInterface
	Category         CategoryControllerInterface
	Instrument       InstrumentControllerInterface
	InstrumentLog    InstrumentLogControllerInterface
	Supply           SupplyControllerInterface
	SupplyLog        SupplyLogControllerInterface
}

var assetInstance *AssetControllerManager

func GetAssetInstance() *AssetControllerManager {
	if assetInstance == nil {
		instance, err := NewAssetControllerFacade()
		if err != nil {
			panic(err)
		}

		assetInstance = instance
	}

	return assetInstance
}

func NewAssetControllerFacade() (*AssetControllerManager, error) {
	manager := &AssetControllerManager{}

	db := migration.GetInstance().DB

	if db == nil {
		return nil, errors.New("err: db not initialized")
	}

	//manager.BorrowInstrument = &BorrowInstrumentController{db: db, BaseController: core.NewBaseController[model.BorrowInstrument]("BorrowInstrument", db)}
	//manager.Category = &CategoryController{db: db, BaseController: core.NewBaseController[model.Category]("Category", db)}
	manager.Instrument = NewInstrumentController()
	manager.InstrumentLog = NewInstrumentLogController()
	manager.Supply = NewSupplyController()
	//manager.SupplyLog = &SupplyLogController{db: db, BaseController: core.NewBaseController("SupplyLog", db)}

	manager.Instrument.addObserver(manager.InstrumentLog)

	return manager, nil
}

func (manager *AssetControllerManager) LoadSeedData() error {
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

func (manager *AssetControllerManager) ResetDB() error {
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

func (manager *AssetControllerManager) ResetAndLoadDB() error {
	err := manager.ResetDB()
	if err != nil {
		return err
	}

	err = manager.LoadSeedData()
	if err != nil {
		return err
	}

	return nil
}
