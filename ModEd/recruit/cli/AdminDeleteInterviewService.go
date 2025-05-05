package cli

import (
	"ModEd/recruit/controller"
)

type AdminInterviewService interface {
	DeleteInterview(interviewID uint) error
}

type adminInterviewService struct {
	interviewCtrl *controller.InterviewController
}

func NewAdminInterviewService(ctrl *controller.InterviewController) AdminInterviewService {
	return &adminInterviewService{
		interviewCtrl: ctrl,
	}
}

func (s *adminInterviewService) DeleteInterview(interviewID uint) error {
	return s.interviewCtrl.DeleteInterview(interviewID)
}
