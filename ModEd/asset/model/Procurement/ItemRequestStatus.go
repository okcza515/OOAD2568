// MEP-1014
package model

type ItemRequestStatus string

const (
	StatusDraft    ItemRequestStatus = "draft"
	StatusPending  ItemRequestStatus = "pending"
	StatusApproved ItemRequestStatus = "approved"
	StatusRejected ItemRequestStatus = "rejected"
)

func ValidItemRequestStatus() []ItemRequestStatus {
	return []ItemRequestStatus{StatusDraft, StatusPending, StatusApproved, StatusRejected}
}

func (s ItemRequestStatus) IsValid() bool {
	for _, valid := range ValidItemRequestStatus() {
		if s == valid {
			return true
		}
	}
	return false
}
