// MEP-1013

package controller

type AssetManager interface {
	GetAllAsset() ([]interface{}, error)
	GetAssetById(Id uint) (interface{}, error)
	GetAssetByRoomId(roomID uint) ([]interface{}, error)
	CreateAsset(payload interface{}) error
	UpdateAsset(Id uint, payload interface{}) error
	DeleteAsset(Id uint) error
}
