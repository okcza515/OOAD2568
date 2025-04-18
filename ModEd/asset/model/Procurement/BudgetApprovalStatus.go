// MEP-1014
package model

type BudgetApprovalStatus string

const (
	BudgetStatusPending  BudgetApprovalStatus = "pending"
	BudgetStatusApproved BudgetApprovalStatus = "approved"
	BudgetStatusRejected BudgetApprovalStatus = "rejected"
)

// Optional: Validation helper
func (s BudgetApprovalStatus) IsValid() bool {
	switch s {
	case BudgetStatusPending, BudgetStatusApproved, BudgetStatusRejected:
		return true
	default:
		return false
	}
}
