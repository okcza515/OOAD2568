// MEP-1003 Student Recruitment
package controller

import (
	"fmt"

	"gorm.io/gorm"
)

type SQLController[T any] interface {
	Create(model *T) error
	GetAll() ([]T, error)
	GetByID(id uint) (T, error)
	Update(model *T) error
	Delete(id uint) error
	ClearTable(tableName string) error
}

type GormSQLController[T any] struct {
	db *gorm.DB
}

func (c *GormSQLController[T]) GetDB() *gorm.DB {
	return c.db
}

func NewGormSQLController[T any](db *gorm.DB) *GormSQLController[T] {
	return &GormSQLController[T]{db}
}

func (c *GormSQLController[T]) Create(model *T) error {
	return c.db.Create(model).Error
}

func (c *GormSQLController[T]) GetAll() ([]T, error) {
	var models []T
	err := c.db.Find(&models).Error
	return models, err
}

func (c *GormSQLController[T]) GetByID(id uint) (T, error) {
	var model T
	err := c.db.First(&model, id).Error
	return model, err
}

func (c *GormSQLController[T]) Update(model *T) error {
	return c.db.Save(model).Error
}

func (c *GormSQLController[T]) Delete(id uint) error {
	var model T
	err := c.db.Delete(&model, id).Error
	return err
}

func (c *GormSQLController[T]) ClearTable(tableName string) error {
	query := fmt.Sprintf("DELETE FROM %s", tableName)
	return c.db.Exec(query).Error
}
