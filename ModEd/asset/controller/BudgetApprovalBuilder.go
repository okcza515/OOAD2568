package controller

import (
	model "ModEd/asset/model"
)

type BudgetApprovalBuilder struct {
	approval *model.BudgetApproval
}

func NewBudgetApprovalBuilder() *BudgetApprovalBuilder {
	return &BudgetApprovalBuilder{
		approval: &model.BudgetApproval{},
	}
}

func (b *BudgetApprovalBuilder) WithInstrumentRequestID(id uint) *BudgetApprovalBuilder {
	b.approval.InstrumentRequestID = id
	return b
}

func (b *BudgetApprovalBuilder) WithApproverID(approverID uint) *BudgetApprovalBuilder {
	b.approval.ApproverID = &approverID
	return b
}

func (b *BudgetApprovalBuilder) WithStatus(status model.BudgetApprovalStatus) *BudgetApprovalBuilder {
	b.approval.Status = status
	return b
}

func (b *BudgetApprovalBuilder) Build() *model.BudgetApproval {
	return b.approval
}
