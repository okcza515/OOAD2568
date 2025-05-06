// MEP-1005

package core

import (
	"errors"

	"gorm.io/gorm"
)

type BaseController[T RecordInterface] struct {
	db *gorm.DB
}

func NewBaseController[T RecordInterface](db *gorm.DB) *BaseController[T] {
	return &BaseController[T]{db: db}
}

func (controller *BaseController[T]) Insert(data T) error {
	return controller.db.Create(&data).Error
}

func (controller *BaseController[T]) InsertMany(data []T) error {
	return controller.db.Create(data).Error
}

func (controller *BaseController[T]) UpdateByID(data T) error {
	result := controller.db.Model(data).Where("id = ?", data.GetID()).Updates(data)
	if result.RowsAffected == 0 {
		return errors.New("record not found")
	}

	return result.Error
}

func (controller *BaseController[T]) UpdateByCondition(condition map[string]interface{}, data T) error {
	result := controller.db.Model(data).Where(condition).Updates(data)
	if result.RowsAffected == 0 {
		return errors.New("record not found")
	}
	return result.Error
}

func (controller *BaseController[T]) RetrieveByID(id uint, preloads ...string) (T, error) {
	var record T
	query := controller.db

	for _, preload := range preloads {
		query = query.Preload(preload)
	}

	if err := query.Where("id = ?", id).First(&record).Error; err != nil {
		return record, err
	}

	return record, nil
}

func (controller *BaseController[T]) RetrieveByCondition(condition map[string]interface{}, preloads ...string) (T, error) {
	var record T

	query := controller.db
	for _, preload := range preloads {
		query = query.Preload(preload)
	}

	if err := query.Where(condition).First(&record).Error; err != nil {
		return record, err
	}

	return record, nil
}

func (controller *BaseController[T]) DeleteByID(id uint) error {
	var record T
	result := controller.db.Where("id = ?", id).Delete(&record)
	if result.RowsAffected == 0 {
		return errors.New("record not found")
	}

	return result.Error
}

func (controller *BaseController[T]) DeleteByCondition(condition map[string]interface{}) error {
	var record T
	result := controller.db.Where(condition).Delete(&record)
	if result.RowsAffected == 0 {
		return errors.New("record not found")
	}

	return result.Error
}

func (controller *BaseController[T]) List(condition map[string]interface{}, preloads ...string) ([]T, error) {
	var records []T
	query := controller.db

	if condition != nil {
		query = query.Where(condition)
	}

	for _, preload := range preloads {
		query = query.Preload(preload)
	}

	if err := query.Find(&records).Error; err != nil {
		return nil, err
	}
	return records, nil
}

func (controller *BaseController[T]) ListPagination(condition map[string]interface{}, page, pageSize int, preloads ...string) ([]T, error) {
	var records []T
	var totalCount int64
	query := controller.db

	if condition != nil {
		query = query.Where(condition)
	}

	for _, preload := range preloads {
		query = query.Preload(preload)
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
