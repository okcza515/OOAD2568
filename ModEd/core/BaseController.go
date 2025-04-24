package core

import (
	"reflect"

	"gorm.io/gorm"
)

type BaseController[T RecordInterface] struct {
	db *gorm.DB
}

func NewBaseController[T RecordInterface](db *gorm.DB) *BaseController[T] {
	return &BaseController[T]{db: db}
}

func (controller *BaseController[T]) Insert(data T) error {
	return controller.db.Create(data).Error
}

func (controller *BaseController[T]) InsertMany(data []T) error {
	return controller.db.Create(data).Error
}

func (controller *BaseController[T]) UpdateByID(data T) error {
	return controller.db.Model(data).Where("id = ?", data.GetID()).Updates(data).Error
}

func (controller *BaseController[T]) RetrieveByID(id uint, preloads ...string) (T, error) {
	var zero T

	ptr := reflect.New(reflect.TypeOf(zero).Elem()).Interface().(T)

	query := controller.db
	for _, preload := range preloads {
		query = query.Preload(preload)
	}

	if err := query.Where("id = ?", id).First(ptr).Error; err != nil {
		return zero, err
	}

	return ptr, nil
}

func (controller *BaseController[T]) DeleteByID(id uint) error {
	var record T
	return controller.db.Where("id = ?", id).Delete(&record).Error
}

func (controller *BaseController[T]) List(condition map[string]interface{}) ([]T, error) {
	var records []T
	query := controller.db

	if condition != nil {
		query = query.Where(condition)
	}

	if err := query.Find(&records).Error; err != nil {
		return nil, err
	}
	return records, nil
}

func (controller *BaseController[T]) ListPagination(condition map[string]interface{}, page, pageSize int) ([]T, error) {
	var records []T
	var totalCount int64
	query := controller.db

	if condition != nil {
		query = query.Where(condition)
	}

	if err := query.Model(new(T)).Count(&totalCount).Error; err != nil {
		return nil, err
	}

	offset := (page - 1) * pageSize

	if err := query.Offset(offset).Limit(pageSize).Find(&records).Error; err != nil {
		return nil, err
	}

	return records, nil
}
