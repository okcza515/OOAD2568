//MEP-1013

package controller

import (
	model "ModEd/asset/model/spacemanagement"
	"errors"
)

type AssetType int

const (
	Instrument AssetType = iota
	Supply
)

type AssetManagementController struct {
	InstrumentManagementAdapter *InstrumentManagementAdapter
	SupplyManagementAdapter     *SupplyManagementAdapter
}

func NewAssetManagementController(instrumentAdapter *InstrumentManagementAdapter, supplyAdapter *SupplyManagementAdapter) *AssetManagementController {
	return &AssetManagementController{
		InstrumentManagementAdapter: instrumentAdapter,
		SupplyManagementAdapter:     supplyAdapter,
	}
}

func (c *AssetManagementController) GetAllAsset(assetType AssetType) (interface{}, error) {
	switch assetType {
	case Instrument:
		return c.InstrumentManagementAdapter.getAllInstrumentManagement()
	case Supply:
		return c.SupplyManagementAdapter.getAllSupplyManagement()
	default:
		return nil, errors.New("invalid asset type")
	}
}

func (c *AssetManagementController) GetAssetById(assetType AssetType, Id uint) (interface{}, error) {
	switch assetType {
	case Instrument:
		return c.InstrumentManagementAdapter.getInstrumentManagementById(Id)
	case Supply:
		return c.SupplyManagementAdapter.getSupplyManagementById(Id)
	default:
		return nil, errors.New("invalid asset type")
	}
}

func (c *AssetManagementController) GetAssetByRoomId(assetType AssetType, roomID uint) (interface{}, error) {
	switch assetType {
	case Instrument:
		return c.InstrumentManagementAdapter.GetInstrumentManagementByRoomId(roomID)
	case Supply:
		return c.SupplyManagementAdapter.GetAssetManagementByRoomId(roomID)
	default:
		return nil, errors.New("invalid asset type")
	}
}

func (c *AssetManagementController) CreateAsset(assetType AssetType, payload interface{}) error {
	switch assetType {
	case Instrument:
		return c.InstrumentManagementAdapter.CreateInstrumentManagement(payload.(*model.InstrumentManagement))
	case Supply:
		return c.SupplyManagementAdapter.CreateAssetManagement(payload.(*model.SupplyManagement))
	default:
		return errors.New("invalid asset type")
	}
}

func (c *AssetManagementController) UpdateAsset(assetType AssetType, Id uint, payload interface{}) error {
	switch assetType {
	case Instrument:
		return c.InstrumentManagementAdapter.UpdateInstrumentManagement(Id, payload.(*model.InstrumentManagement))
	case Supply:
		return c.SupplyManagementAdapter.UpdateSupplyManagement(Id, payload.(*model.SupplyManagement))
	default:
		return errors.New("invalid asset type")
	}
}

func (c *AssetManagementController) DeleteAsset(assetType AssetType, Id uint) error {
	switch assetType {
	case Instrument:
		return c.InstrumentManagementAdapter.DeleteInstrumentManagement(Id)
	case Supply:
		return c.SupplyManagementAdapter.DeleteSupplyManagement(Id)
	default:
		return errors.New("invalid asset type")
	}
}
