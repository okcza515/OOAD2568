// MEP-1003 Student Recruitment
package controller

import (
	"ModEd/core"
	"ModEd/recruit/model"
	"ModEd/recruit/util"
	"fmt"

	"gorm.io/gorm"
)

type ApplicationRoundController struct {
	Base *core.BaseController[*model.ApplicationRound]
	DB   *gorm.DB
}

func NewApplicationRoundController(db *gorm.DB) *ApplicationRoundController {
	return &ApplicationRoundController{
		Base: core.NewBaseController[*model.ApplicationRound](db),
		DB:   db,
	}
}

func (controller *ApplicationRoundController) CreateApplicationRound(round *model.ApplicationRound) error {
	return controller.Base.Insert(round)
}

func (c *ApplicationRoundController) GetAllRounds() ([]*model.ApplicationRound, error) {
	return c.Base.List(nil)
}

func (c *ApplicationRoundController) ReadApplicationRoundsFromCSV(filePath string) error {
	if err := c.DB.Exec("DELETE FROM application_rounds").Error; err != nil {
		fmt.Println("Error clearing table:", err)
		return err
	}

	rounds, err := util.InsertFromCSVOrJSON[model.ApplicationRound](filePath, c.DB)
	if err != nil {
		return err
	}

	fmt.Printf("Inserted %d application rounds from file.\n", len(rounds))
	return nil
}
