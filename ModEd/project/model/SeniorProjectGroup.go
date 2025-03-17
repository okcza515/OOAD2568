package model

import (
	"ModEd/common/model"

	"gorm.io/gorm"
)

type SeniorProjectGroup struct {
	gorm.Model
	GroupMember  []model.Student           `json:"group_member"`
	Advisor      Advisor                   `json:"advisor"`
	Committee    []Committee               `json:"committee"`
	Schedule     []SeniorProjectSchedule   `json:"schedule"`
	Assignment   []SeniroProjectAssignment `json:"assignment"`
	Presentation []Presentation            `json:"presentation"`
	Report       []Report                  `json:"report"`
	Assessment   Assessment                `json:"assessment"`
}
