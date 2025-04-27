// MEP-1014
package model

type ProcurementStatus string

const (
	ProcurementStatusPending  ProcurementStatus = "pending"
	ProcurementStatusApproved ProcurementStatus = "approved"
	ProcurementStatusRejected ProcurementStatus = "rejected"
)

func (s ProcurementStatus) IsValid() bool {
	switch s {
	case ProcurementStatusPending, ProcurementStatusApproved, ProcurementStatusRejected:
		return true
	default:
		return false
	}
}
