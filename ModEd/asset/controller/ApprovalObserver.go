// MEP-1014
package controller

type ApprovalObserver interface {
	OnApproved(id uint, approverID uint) error
	OnRejected(id uint, approverID uint) error
}

