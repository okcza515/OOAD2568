package controller

// MEP-1012 Asset

import (
	"ModEd/asset/model"
	"ModEd/core"
	"ModEd/core/migration"
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
		instance, err := newAssetControllerManager()
		if err != nil {
			panic(err)
		}

		assetInstance = instance
	}

	return assetInstance
}

func newAssetControllerManager() (*AssetControllerManager, error) {
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
	manager.SupplyLog = NewSupplyLogController()

	manager.Instrument.addObserver(manager.InstrumentLog)
	manager.Supply.addObserver(manager.SupplyLog)

	return manager, nil
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

	err = migration.GetInstance().
		AddSeedData("data/asset/Category.csv", &[]model.Category{}).
		//AddSeedData("data/asset/BorrowInstrument.csv", &[]model.BorrowInstrument{}).
		AddSeedData("data/asset/InstrumentList.csv", &[]model.Instrument{}).
		//AddSeedData("data/asset/InstrumentLog.csv", &[]model.InstrumentLog{}).
		AddSeedData("data/asset/SupplyList.csv", &[]model.Supply{}).
		AddSeedData("data/asset/SupplyLog.csv", &[]model.SupplyLog{}).
		LoadSeedData()
	if err != nil {
		return err
	}

	return nil
}
