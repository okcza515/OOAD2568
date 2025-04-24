package core

import (
	"gorm.io/gorm"
)

type StagingBaseController struct {
	db        *gorm.DB
	modelName string
}

func NewStagingStagingBaseController(modelName string, db *gorm.DB) *StagingBaseController {
	controller := &StagingBaseController{
		modelName: modelName,
		db:        db,
	}
	return controller
}

func (controller *StagingBaseController) Insert(data RecordInterface) error {
	return controller.db.Create(data).Error
}

func (controller *StagingBaseController) InsertMany(data interface{}) error {
	return controller.db.Create(data).Error
}

func (controller *StagingBaseController) UpdateByID(data RecordInterface) error {
	return controller.db.Model(data).Where("id = ?", (data).GetID()).Updates(data).Error
}

func (controller *StagingBaseController) RetrieveByID(id uint, preloads ...string) (*RecordInterface, error) {
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

func (controller *StagingBaseController) DeleteByID(id uint) error {
	var record RecordInterface
	return controller.db.Where("id = ?", id).Delete(&record).Error
}

func (controller *StagingBaseController) Delete(data *RecordInterface) error {
	return controller.db.Where("id = ?", (*data).GetID()).Delete(data).Error
}

func (controller *StagingBaseController) List(condition map[string]interface{}) ([]RecordInterface, error) {
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

func (controller *StagingBaseController) ListPagination(condition map[string]interface{}, page int, pageSize int) (*PaginationResult, error) {
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
