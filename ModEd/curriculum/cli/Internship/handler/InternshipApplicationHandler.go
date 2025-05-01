package handler

import (
    "ModEd/core/cli"
    "ModEd/curriculum/controller"
    "ModEd/curriculum/model"
    "fmt"
    "time"
)

type InternshipApplicationHandler struct {
    manager *cli.CLIMenuStateManager
    wrapper *controller.InternshipModuleWrapper
}

func NewInternshipApplicationHandler(manager *cli.CLIMenuStateManager, wrapper *controller.InternshipModuleWrapper) *InternshipApplicationHandler {
    return &InternshipApplicationHandler{
        manager: manager,
        wrapper: wrapper,
    }
}

func (handler *InternshipApplicationHandler) RegisterApplication(studentCode string, companyId uint, advisorCode uint) error {
    application := &model.InternshipApplication{
        TurninDate:            time.Now(),
        ApprovalAdvisorStatus: model.WAIT,
        ApprovalCompanyStatus: model.WAIT,
        AdvisorCode:           advisorCode,
        CompanyId:             companyId,
        StudentCode:           studentCode,
    }

    err := handler.wrapper.InternshipApplication.RegisterInternshipApplications([]*model.InternshipApplication{application})
    if err != nil {
        return fmt.Errorf("failed to register internship application: %w", err)
    }

    fmt.Println("Internship application registered successfully!")
    return nil
}

func (handler *InternshipApplicationHandler) ListApplications() error {
    applications, err := handler.wrapper.InternshipApplication.GetAllInternshipApplications()
    if err != nil {
        return fmt.Errorf("failed to retrieve internship applications: %w", err)
    }

    fmt.Println("Internship Applications:")
    for _, app := range applications {
        fmt.Printf("ID: %d, StudentCode: %s, AdvisorStatus: %s, CompanyStatus: %s\n",
            app.ID, app.StudentCode, app.ApprovalAdvisorStatus, app.ApprovalCompanyStatus)
    }
    return nil
}

func (handler *InternshipApplicationHandler) GetApplicationStatus(id uint) error {
    status, err := handler.wrapper.InternshipApplication.GetApplicationStatusByID(id)
    if err != nil {
        return fmt.Errorf("failed to retrieve application status: %w", err)
    }

    fmt.Printf("Application Status for ID %d: %s\n", id, status)
    return nil
}

func (handler *InternshipApplicationHandler) DeleteApplication(id uint) error {
    err := handler.wrapper.InternshipApplication.DeleteApplicationByID(id)
    if err != nil {
        return fmt.Errorf("failed to delete internship application: %w", err)
    }

    fmt.Printf("Internship application with ID %d deleted successfully!\n", id)
    return nil
}