package model

type InstrumentRequestStatus string

const (
	InstrumentRequestStatusDraft    InstrumentRequestStatus = "draft"
	InstrumentRequestStatusPending  InstrumentRequestStatus = "pending"
	InstrumentRequestStatusApproved InstrumentRequestStatus = "approved"
	InstrumentRequestStatusRejected InstrumentRequestStatus = "rejected"
)

func ValidItemRequestStatus() []InstrumentRequestStatus {
	return []InstrumentRequestStatus{
		InstrumentRequestStatusDraft,
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
