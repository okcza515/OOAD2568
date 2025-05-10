// MEP-1014
package model

type AcceptanceStatus string

const (
	AcceptanceStatusPending  AcceptanceStatus = "pending"
	AcceptanceStatusApproved AcceptanceStatus = "approved"
	AcceptanceStatusRejected AcceptanceStatus = "rejected"
)

func (s AcceptanceStatus) IsValid() bool {
	switch s {
	case AcceptanceStatusPending, AcceptanceStatusApproved, AcceptanceStatusRejected:
		return true
	default:
		return false
	}
}
