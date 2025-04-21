package model

type InstrumentRequestStatus string

const (
	StatusDraft    InstrumentRequestStatus = "draft"
	StatusPending  InstrumentRequestStatus = "pending"  
	StatusApproved InstrumentRequestStatus = "approved" 
	StatusRejected InstrumentRequestStatus = "rejected"
)

func ValidItemRequestStatus() []InstrumentRequestStatus {
	return []InstrumentRequestStatus{
		StatusDraft,
		StatusPending,
		StatusApproved,
		StatusRejected,
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
