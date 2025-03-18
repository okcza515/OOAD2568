package core

import (
	"gorm.io/gorm"
)

type BaseController[T any] struct {
	db *gorm.DB
}

func NewBaseController[T any](db *gorm.DB) *BaseController[T] {
	return &BaseController[T]{db: db}
}

func (c *BaseController[T]) Insert(data T) error {
	return c.db.Create(data).Error
}

func (c *BaseController[T]) UpdateByID(id uint, data *T) error {
	return c.db.Model(data).Where("id = ?", id).Updates(data).Error
}

func (c *BaseController[T]) RetrieveByID(id uint) (*T, error) {
	var records T
	if err := c.db.Where("id = ?", id).First(&records).Error; err != nil {
		return nil, err
	}
	return &records, nil
}

func (c *BaseController[T]) DeleteByID(id uint) error {
	var record T
	return c.db.Where("id = ?", id).Delete(&record).Error
}

func (c *BaseController[T]) List(condition map[string]interface{}) ([]T, error) {
	var records []T
	query := c.db

	if condition != nil {
		query = query.Where(condition)
	}

	if err := query.Find(&records).Error; err != nil {
		return nil, err
	}
	return records, nil
}

func (c *BaseController[T]) ListPagination(condition map[string]interface{}, page, pageSize int) ([]T, int64, error) {
	var records []T
	var totalCount int64
	query := c.db

	if condition != nil {
		query = query.Where(condition)
	}

	if err := query.Model(new(T)).Count(&totalCount).Error; err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * pageSize

	if err := query.Offset(offset).Limit(pageSize).Find(&records).Error; err != nil {
		return nil, 0, err
	}

	return records, totalCount, nil
}
