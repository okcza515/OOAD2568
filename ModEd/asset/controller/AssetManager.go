// MEP-1013

package controller

type AssetManager interface {
	GetAll() ([]interface{}, error)
	GetById(Id uint) (interface{}, error)
	GetByRoomId(roomID uint) ([]interface{}, error)
	Create(payload interface{}) error
	Update(Id uint, payload interface{}) error
	Delete(Id uint) error
}