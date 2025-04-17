package controller

import (
	"ModEd/recruit/controller/SQL"
	"ModEd/recruit/model"
	"ModEd/recruit/util"
	"fmt"

	"gorm.io/gorm"
)

type ApplicationRoundController struct {
	sqlCtrl SQL.SQLController[model.ApplicationRound]
}

func CreateApplicationRoundController(db *gorm.DB) *ApplicationRoundController {
	return &ApplicationRoundController{
		sqlCtrl: SQL.NewGormSQLController[model.ApplicationRound](db),
	}
}

func (controller *ApplicationRoundController) CreateApplicationRound(round *model.ApplicationRound) error {
	return controller.sqlCtrl.Create(round)
}

func (c *ApplicationRoundController) ReadApplicationRoundsFromCSV(filePath string) error {
	gormDB := c.sqlCtrl.(*SQL.GormSQLController[model.ApplicationRound]).GetDB()

	if err := c.sqlCtrl.ClearTable("application_rounds"); err != nil {
		fmt.Println("Error clearing table:", err)
		return err
	}

	rounds, err := util.InsertFromCSVOrJSON[model.ApplicationRound](filePath, gormDB)
	if err != nil {
		return err
	}

	fmt.Printf("Inserted %d application rounds from file.\n", len(rounds))
	return nil
}

func (c *ApplicationRoundController) GetAllRounds() ([]*model.ApplicationRound, error) {
	models, err := c.sqlCtrl.GetAll()
	if err != nil {
		return nil, err
	}

	var rounds []*model.ApplicationRound
	for _, model := range models {
		rounds = append(rounds, &model)
	}

	return rounds, nil
}
