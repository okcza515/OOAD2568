// MEP-1014
package model

type QuotationStatus string

const (
	QuotationStatusPending  QuotationStatus = "pending"
	QuotationStatusApproved QuotationStatus = "approved"
	QuotationStatusRejected QuotationStatus = "rejected"
)

func (s QuotationStatus) IsValid() bool {
	switch s {
	case QuotationStatusPending, QuotationStatusApproved, QuotationStatusRejected:
		return true
	default:
		return false
	}
}
