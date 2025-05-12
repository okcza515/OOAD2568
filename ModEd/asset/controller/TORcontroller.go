// MEP-1014
package controller

import (
	model "ModEd/asset/model"
	"fmt"

	"gorm.io/gorm"
)

type TORController struct {
	db *gorm.DB
}

func CreateTORController(db *gorm.DB) *TORController {
	return &TORController{db: db}
}

func (c *TORController) CreateTOR(tor *model.TOR) error {
	return c.db.Create(tor).Error
}

func (c *TORController) GetAllTORs() ([]model.TOR, error) {
	var tors []model.TOR
	err := c.db.
		Preload("InstrumentRequest.Instruments.Category").
		Find(&tors).Error
	return tors, err
}

func (c *TORController) GetTORByID(id uint) (*model.TOR, error) {
	var tor model.TOR
	err := c.db.
		Preload("InstrumentRequest.Instruments.Category").
		First(&tor, "tor_id = ?", id).Error
	return &tor, err
}

func (c *TORController) DeleteTOR(id uint) error {
	return c.db.Delete(&model.TOR{}, id).Error
}

func (c *TORController) UpdateTotalPrice(torID uint) error {
	var tor model.TOR
	err := c.db.Preload("Quotations.Details").First(&tor, torID).Error
	if err != nil {
		fmt.Printf("Failed to find TOR with ID %d: %v\n", torID, err)
		return err
	}

	fmt.Printf("Loaded TOR ID: %d with %d quotations\n", tor.TORID, len(tor.Quotations))

	totalPrice := 0.0
	for _, quotation := range tor.Quotations {
		fmt.Printf("Quotation ID: %d has %d details\n", quotation.QuotationID, len(quotation.Details))
		for _, detail := range quotation.Details {
			totalPrice += detail.OfferedPrice * float64(detail.Quantity)
			fmt.Printf("Adding Price: %.2f * %d = %.2f\n", detail.OfferedPrice, detail.Quantity, detail.OfferedPrice*float64(detail.Quantity))
		}
	}

	fmt.Printf("Total price calculated: %.2f\n", totalPrice)

	err = c.db.Model(&model.TOR{}).
		Where("tor_id = ?", torID).
		Update("total_price", totalPrice).Error

	if err != nil {
		fmt.Printf("Failed to update TotalPrice for TOR %d: %v\n", torID, err)
		return err
	}

	fmt.Printf("TotalPrice for TOR ID %d updated to %.2f\n", torID, totalPrice)
	return nil
}
