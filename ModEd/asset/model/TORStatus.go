// MEP-1014
package model

type TORStatus string

const (
	TORStatusAnnounced TORStatus = "announced"
	TORStatusSelected  TORStatus = "quotationselected"
)

func (s TORStatus) IsValid() bool {
	switch s {
	case TORStatusAnnounced, TORStatusSelected:
		return true
	default:
		return false
	}
}
