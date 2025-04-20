package controller

import (
	"ModEd/core"
	"ModEd/recruit/model"
	"ModEd/recruit/util"
	"fmt"

	"gorm.io/gorm"
)

type AdminController struct {
	Base *core.BaseController
	DB   *gorm.DB
}

// สร้าง AdminController
func CreateAdminController(db *gorm.DB) *AdminController {
	return &AdminController{
		Base: core.NewBaseController("Admin", db),
		DB:   db,
	}
}

// Insert Admin ธรรมดา
func (controller *AdminController) CreateAdmin(admin *model.Admin) error {
	return controller.Base.Insert(admin)
}

// ดึง Admin ทั้งหมด
func (c *AdminController) GetAllAdmins() ([]*model.Admin, error) {
	var admins []*model.Admin

	if err := c.DB.Find(&admins).Error; err != nil {
		return nil, fmt.Errorf("failed to query admins: %w", err)
	}

	return admins, nil
}

// อ่าน CSV แล้ว Insert
func (c *AdminController) ReadAdminsFromCSV(filePath string) error {
	// Clear table ก่อน
	if err := c.DB.Exec("DELETE FROM admins").Error; err != nil {
		fmt.Println("Error clearing table:", err)
		return err
	}

	// Insert ข้อมูลจาก CSV
	admins, err := util.InsertFromCSVOrJSON[model.Admin](filePath, c.DB)
	if err != nil {
		return err
	}

	fmt.Printf("Inserted %d admins from file.\n", len(admins))
	return nil
}
