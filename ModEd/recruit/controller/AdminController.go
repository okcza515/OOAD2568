package controller

import (
	"ModEd/recruit/controller/SQL"
	"ModEd/recruit/model"
	"ModEd/recruit/util"
	"fmt"

	"gorm.io/gorm"
)

type AdminController struct {
	sqlCtrl SQL.SQLController[model.Admin]
}

func CreateAdminController(db *gorm.DB) *AdminController {
	return &AdminController{
		sqlCtrl: SQL.NewGormSQLController[model.Admin](db),
	}
}

func (controller *AdminController) CreateAdmin(admin *model.Admin) error {
	return controller.sqlCtrl.Create(admin)
}

func (c *AdminController) ReadAdminsFromCSV(filePath string) error {
	gormDB := c.sqlCtrl.(*SQL.GormSQLController[model.Admin]).GetDB()

	if err := c.sqlCtrl.ClearTable("admins"); err != nil {
		fmt.Println("Error clearing table:", err)
		return err
	}

	admins, err := util.InsertFromCSVOrJSON[model.Admin](filePath, gormDB)
	if err != nil {
		return err
	}

	fmt.Printf("Inserted %d admins from file.\n", len(admins))
	return nil
}
