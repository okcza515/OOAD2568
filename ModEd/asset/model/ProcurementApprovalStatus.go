// MEP-1014
package model

type ProcurementApprovalStatus string

const (
	ProcurementApprovalStatusPending  ProcurementApprovalStatus = "pending"
	ProcurementApprovalStatusApproved ProcurementApprovalStatus = "approved"
	ProcurementApprovalStatusRejected ProcurementApprovalStatus = "rejected"
)

func (s ProcurementApprovalStatus) IsValid() bool {
	switch s {
	case ProcurementApprovalStatusPending, ProcurementApprovalStatusApproved, ProcurementApprovalStatusRejected:
		return true
	default:
		return false
	}
}
