// MEP-1014
package model

type AcceptanceStatus string

const (
	AcceptanceStatusPending  AcceptanceStatus = "pending"
	AcceptanceStatusApproved AcceptanceStatus = "approved"
	AcceptanceStatusRejected AcceptanceStatus = "rejected"
	AcceptanceStatusImported AcceptanceStatus = "imported"
)

func (s AcceptanceStatus) IsValid() bool {
	switch s {
	case AcceptanceStatusPending, AcceptanceStatusApproved, AcceptanceStatusRejected, AcceptanceStatusImported:
		return true
	default:
		return false
	}
}
