// MEP-1014
package controller

import (
	model "ModEd/asset/model"

	"gorm.io/gorm"
)

type AcceptanceTestController struct {
	db *gorm.DB
}

func (c *AcceptanceTestController) CreateAcceptanceTest(body *model.AcceptanceTest) error {
	return c.db.Create(body).Error
}

func (c *AcceptanceTestController) ListAllAcceptanceTest() (*[]model.AcceptanceTest, error) {
	var acceptancetest []model.AcceptanceTest
	err := c.db.Find(&acceptancetest).Error
	return &acceptancetest, err
}

func (c *AcceptanceTestController) GetAcceptanceTestByID(id uint) (*model.AcceptanceTest, error) {
	var acceptancetest model.AcceptanceTest
	err := c.db.First(&acceptancetest, id).Error
	return &acceptancetest, err
}

func (c *AcceptanceTestController) GetQuotationDetailsByTOR(torID uint) ([]model.QuotationDetail, error) {
	var quotations []model.Quotation

	err := c.db.Preload("Details").
		Where("tor_id = ?", torID).
		Find(&quotations).Error
	if err != nil {
		return nil, err
	}

	var details []model.QuotationDetail
	for _, quotation := range quotations {
		details = append(details, quotation.Details...)
	}

	return details, nil
}

func (c *AcceptanceTestController) GetCategoriesByIDs(ids []uint) ([]model.Category, error) {
	var categories []model.Category
	if len(ids) == 0 {
		return categories, nil
	}

	err := c.db.Where("id IN ?", ids).Find(&categories).Error
	return categories, err
}
