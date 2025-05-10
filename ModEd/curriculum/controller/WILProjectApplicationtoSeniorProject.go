package controller

import (
	"ModEd/curriculum/model"
	"ModEd/project/controller"
	"fmt"

	"gorm.io/gorm"
)

type WILProjectApplicationToSeniorProjectController struct {
    Connector               *gorm.DB
    SeniorProjectController *controller.SeniorProjectController
}

func NewWILToSeniorProjectController(Connector *gorm.DB) *WILProjectApplicationToSeniorProjectController {
    Connector.AutoMigrate(&model.WILProjectApplication{})
    return &WILProjectApplicationToSeniorProjectController{
        Connector:               Connector,
        SeniorProjectController: controller.NewSeniorProjectController(Connector), 
    }
}

func (controller *WILProjectApplicationToSeniorProjectController) NewSeniorProjectbyWILProjectApplication(wilProject *model.WILProjectApplication) (uint, error) {
    adapter := &model.WILProjectApplicationToSeniorProjectAdapter{WILProject: wilProject}
    fmt.Println("adapter.WILProject:", adapter.WILProject)
    seniorProject, err := adapter.ToSeniorProject()
    if err != nil {
        return 0, err
    }
    
    err = controller.SeniorProjectController.Insert(seniorProject)

	if err != nil {
		return 0, err
	}

    return seniorProject.ID, nil
}