package handler

import (
    "ModEd/core/cli"
    "ModEd/curriculum/controller"
    "ModEd/curriculum/model"
    "ModEd/curriculum/utils"
    "fmt"
    "strconv"
    "time"
)

type InternShipModuleMenuStateHandler struct {
    menuManager *cli.CLIMenuStateManager
    wrapper     *controller.InternshipModuleWrapper
}

func NewInternShipModuleMenuStateHandler(manager *cli.CLIMenuStateManager, wrapper *controller.InternshipModuleWrapper) *InternShipModuleMenuStateHandler {
    return &InternShipModuleMenuStateHandler{
        menuManager: manager,
        wrapper:     wrapper,
    }
}

func (handler *InternShipModuleMenuStateHandler) Render() {
    fmt.Println("\n==== Internship Application System ====")
    fmt.Println("1. Create Internship Application")
    fmt.Println("2. Evaluate Student Performance")
    fmt.Println("3. Evaluate Student Report")
    fmt.Println("4. Update Approval Status")
    fmt.Println("Type 'exit' to quit")
    fmt.Print("Enter your choice: ")
}

func (handler *InternShipModuleMenuStateHandler) HandleUserInput(input string) error {
    switch input {
    case "1":
        return handler.handleCreateInternshipApplication()
    case "2":
        return handler.handleEvaluateStudentPerformance()
    case "3":
        return handler.handleEvaluateStudentReport()
    case "4":
        return handler.handleUpdateApprovalStatus()
    case "exit":
        fmt.Println("Exiting...")
        return nil
    default:
        fmt.Println("Invalid input")
        return nil
    }
}

func (handler *InternShipModuleMenuStateHandler) handleCreateInternshipApplication() error {
    studentCode := utils.GetUserInput("Enter Student Code: ")

    companyName := utils.GetUserInput("Enter Company Name: ")

    company, err := handler.wrapper.Company.GetCompanyByName(companyName)
    if err != nil {
        fmt.Printf("Error finding company with name '%s': %v\n", companyName, err)
        return err
    }

    advisorCodeStr := utils.GetUserInput("Enter Advisor Code: ")
    advisorCode, err := strconv.Atoi(advisorCodeStr)
    if err != nil {
        fmt.Println("Invalid Advisor Code.")
        return err
    }

    application := &model.InternshipApplication{
        TurninDate:            time.Now(),
        ApprovalAdvisorStatus: model.WAIT,
        ApprovalCompanyStatus: model.WAIT,
        AdvisorCode:           uint(advisorCode),
        CompanyId:             company.ID,
        StudentCode:           studentCode,
    }

    err = handler.wrapper.InternshipApplication.RegisterInternshipApplications([]*model.InternshipApplication{application})
    if err != nil {
        fmt.Println("Error creating internship application:", err)
        return err
    }

    fmt.Println("Internship application created successfully!")
    return nil
}

func (handler *InternShipModuleMenuStateHandler) handleUpdateApprovalStatus() error {
    studentCode := utils.GetUserInput("Enter Student Code: ")

    advisorStatusStr := utils.GetUserInput("Enter Advisor Approval Status (APPROVED/REJECT): ")
    advisorStatus := model.ApprovedStatus(advisorStatusStr)

    companyStatusStr := utils.GetUserInput("Enter Company Approval Status (APPROVED/REJECT): ")
    companyStatus := model.ApprovedStatus(companyStatusStr)

    err := handler.wrapper.Approved.UpdateApprovalStatuses(studentCode, advisorStatus, companyStatus)
    if err != nil {
        fmt.Println("Error updating approval statuses:", err)
        return err
    }

    fmt.Println("Approval statuses updated successfully!")
    return nil
}

func (handler *InternShipModuleMenuStateHandler) handleEvaluateStudentPerformance() error {
    studentCode := utils.GetUserInput("Enter Student Code: ")

    reportScoreStr := utils.GetUserInput("Enter Report Score: ")
    reportScore, err := strconv.Atoi(reportScoreStr)
    if err != nil {
        fmt.Println("Invalid Report Score.")
        return err
    }

    err = handler.wrapper.Report.UpdateReportScore(studentCode, reportScore)
    if err != nil {
        fmt.Println("Error updating report score:", err)
        return err
    }

    fmt.Println("Student performance evaluated successfully!")
    return nil
}

func (handler *InternShipModuleMenuStateHandler) handleEvaluateStudentReport() error {
    studentCode := utils.GetUserInput("Enter Student Code: ")

    supervisorScoreStr := utils.GetUserInput("Enter Supervisor Score: ")
    supervisorScore, err := strconv.Atoi(supervisorScoreStr)
    if err != nil {
        fmt.Println("Invalid Supervisor Score.")
        return err
    }

    mentorScoreStr := utils.GetUserInput("Enter Mentor Score: ")
    mentorScore, err := strconv.Atoi(mentorScoreStr)
    if err != nil {
        fmt.Println("Invalid Mentor Score.")
        return err
    }

    err = handler.wrapper.Review.UpdateReviewScore(studentCode, supervisorScore, mentorScore)
    if err != nil {
        fmt.Println("Error updating review score:", err)
        return err
    }

    fmt.Println("Student report evaluated successfully!")
    return nil
}