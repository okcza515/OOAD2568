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
	Base *core.BaseController[*model.Admin]
	DB   *gorm.DB
}

func NewAdminController(db *gorm.DB) *AdminController {
	return &AdminController{
		Base: core.NewBaseController[*model.Admin](db),
		DB:   db,
	}
}

func (controller *AdminController) CreateAdmin(admin *model.Admin) error {
	return controller.Base.Insert(admin)
}

func (c *AdminController) GetAllAdmins() ([]*model.Admin, error) {
	return c.Base.List(nil)
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
