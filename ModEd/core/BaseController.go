package core

import (
	"gorm.io/gorm"
)

type BaseController struct {
	db        *gorm.DB
	modelName string
}

func NewBaseController(modelName string, db *gorm.DB) *BaseController {
	controller := &BaseController{
		modelName: modelName,
		db:        db,
	}
	return controller
}

func (controller *BaseController) Insert(data RecordInterface) error {
	return controller.db.Create(data).Error
}

func (controller *BaseController) InsertMany(data []RecordInterface) error {
	err := controller.db.Transaction(func(tx *gorm.DB) error {
		for _, record := range data {
			if err := tx.Create(record).Error; err != nil {
				return err
			}
		}
		return nil
	})

	return err
}

func (controller *BaseController) UpdateByID(data RecordInterface) error {
	return controller.db.Model(data).Where("id = ?", (data).GetID()).Updates(data).Error
}

func (controller *BaseController) RetrieveByID(id uint, preloads ...string) (*RecordInterface, error) {
	record := CreateRecord(controller.modelName)

	query := controller.db
	for _, preload := range preloads {
		query = query.Preload(preload)
	}

	if err := query.Where("id = ?", id).First(&record).Error; err != nil {
		return nil, err
	}

	return record, nil
}

func (controller *BaseController) DeleteByID(id uint) error {
	var record RecordInterface
	return controller.db.Where("id = ?", id).Delete(&record).Error
}

func (controller *BaseController) Delete(data *RecordInterface) error {
	return controller.db.Where("id = ?", (*data).GetID()).Delete(data).Error
}

func (controller *BaseController) List(condition map[string]interface{}) ([]RecordInterface, error) {
	var records []RecordInterface
	query := controller.db
	query.Model(CreateRecord(controller.modelName))

	if condition != nil {
		query = query.Where(condition)
	}

	if err := query.Find(&records).Error; err != nil {
		return nil, err
	}
	return records, nil
}

func (controller *BaseController) ListPagination(condition map[string]interface{}, page int, pageSize int) (*PaginationResult, error) {
	query := controller.db
	var result PaginationResult
	result.Page = page
	result.PageSize = pageSize
	query.Model(CreateRecord(controller.modelName))

	if condition != nil {
		query = query.Where(condition)
	}

	if err := query.Count(&result.TotalCount).Error; err != nil {
		return nil, err
	}

	offset := (page - 1) * pageSize

	if err := query.Offset(offset).Limit(pageSize).Find(&result.Records).Error; err != nil {
		return nil, err
	}

	return &result, nil
}
