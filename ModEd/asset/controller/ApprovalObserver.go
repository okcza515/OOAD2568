// MEP-1014
package controller

type ApprovalObserver interface {
	OnApproved(id uint)
	OnRejected(id uint)
}
