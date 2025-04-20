// MEP-1003 Student Recruitment
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

func CreateAdminController(db *gorm.DB) *AdminController {
	return &AdminController{
		Base: core.NewBaseController("Admin", db),
		DB:   db,
	}
}

func (controller *AdminController) CreateAdmin(admin *model.Admin) error {
	return controller.Base.Insert(admin)
}

func (c *AdminController) GetAllAdmins() ([]*model.Admin, error) {
	var admins []*model.Admin

	if err := c.DB.Find(&admins).Error; err != nil {
		return nil, fmt.Errorf("failed to query admins: %w", err)
	}

	return admins, nil
}

func (c *AdminController) ReadAdminsFromCSV(filePath string) error {
	if err := c.DB.Exec("DELETE FROM admins").Error; err != nil {
		fmt.Println("Error clearing table:", err)
		return err
	}

	admins, err := util.InsertFromCSVOrJSON[model.Admin](filePath, c.DB)
	if err != nil {
		return err
	}

	fmt.Printf("Inserted %d admins from file.\n", len(admins))
	return nil
}
