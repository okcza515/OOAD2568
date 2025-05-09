// MEP-1014
package controller

import (
	model "ModEd/asset/model"
	"fmt"

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

func (c *AcceptanceTestController) GetCategoriesByIDs(ids []uint) ([]model.Category, error) {
	var categories []model.Category
	if len(ids) == 0 {
		return categories, nil
	}

	err := c.db.Where("id IN ?", ids).Find(&categories).Error
	return categories, err
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

func (c *AcceptanceTestController) PrintQuotationDetailsByTOR(torID uint) {
	details, err := c.GetQuotationDetailsByTOR(torID)
	if err != nil {
		fmt.Println("Error retrieving quotation details:", err)
		return
	}

	if len(details) == 0 {
		fmt.Println("No quotation details found for TOR ID:", torID)
		return
	}

	fmt.Println("Quotation Details for TOR ID:", torID)
	for _, detail := range details {
		fmt.Printf("QuotationDetailID: %d\n", detail.QuotationDetailID)
		fmt.Printf("InstrumentLabel: %s\n", detail.InstrumentLabel)
		if detail.Description != nil {
			fmt.Printf("Description: %s\n", *detail.Description)
		} else {
			fmt.Println("Description: (none)")
		}
		fmt.Printf("CategoryID: %d\n", detail.CategoryID)
		fmt.Printf("Quantity: %d\n", detail.Quantity)
		fmt.Printf("Offered Price: %.2f\n", detail.OfferedPrice)
		fmt.Println("------")
	}
}

func (c *AcceptanceTestController) PrintCategoriesByIDs(ids []uint) {
	categories, err := c.GetCategoriesByIDs(ids)
	if err != nil {
		fmt.Println("Error retrieving categories:", err)
		return
	}

	if len(categories) == 0 {
		fmt.Println("No categories found for the provided IDs.")
		return
	}

	fmt.Println("Categories:")
	for _, category := range categories {
		fmt.Printf("ID: %d\n", category.ID)
		fmt.Printf("Category Name: %s\n", category.CategoryName)
		if category.Description != nil {
			fmt.Printf("Description: %s\n", *category.Description)
		} else {
			fmt.Println("Description: (none)")
		}
		fmt.Println("------")
	}
}
