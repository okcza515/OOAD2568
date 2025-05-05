// MEP-1014
package model

type InstrumentRequestStatus string

const (
	InstrumentRequestStatusPending  InstrumentRequestStatus = "pending"
	InstrumentRequestStatusApproved InstrumentRequestStatus = "approved"
	InstrumentRequestStatusRejected InstrumentRequestStatus = "rejected"
)

func ValidItemRequestStatus() []InstrumentRequestStatus {
	return []InstrumentRequestStatus{
		InstrumentRequestStatusPending,
		InstrumentRequestStatusApproved,
		InstrumentRequestStatusRejected,
	}
}

func (s InstrumentRequestStatus) IsValid() bool {
	for _, valid := range ValidItemRequestStatus() {
		if s == valid {
			return true
		}
	}
	return false
}
