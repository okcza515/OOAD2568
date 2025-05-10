package model

type RequestStatus interface {
	SetStatus(status string)
	SetReason(reason string)
}
