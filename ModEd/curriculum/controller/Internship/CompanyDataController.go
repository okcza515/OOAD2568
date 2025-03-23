package controller

import (
    model "ModEd/curriculum/model/Internship"
    "ModEd/utils/deserializer"
    "fmt"
    "gorm.io/gorm"
)

type CompanyDataController struct {
    Db *gorm.DB
}

func NewCompanyDataController(db *gorm.DB) *CompanyDataController {
    return &CompanyDataController{Db: db}
}

func (c *CompanyDataController) ImportCompaniesFromCSV(filePath string) error {
    fileDeserializer, err := deserializer.NewFileDeserializer(filePath)
    if err != nil {
        return fmt.Errorf("failed to create file deserializer: %w", err)
    }

    var companies []model.Company
    if err := fileDeserializer.Deserialize(&companies); err != nil {
        return fmt.Errorf("failed to deserialize CSV file: %w", err)
    }

    for _, company := range companies {
        if err := c.Db.Create(&company).Error; err != nil {
            return fmt.Errorf("failed to insert company into database: %w", err)
        }
    }

    return nil
}