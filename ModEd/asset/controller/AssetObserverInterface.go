package controller

import (
	"ModEd/core"
)

type AssetObserver[T core.RecordInterface] interface {
	HandleEvent(eventType string, dataContext T)
	GetObserverID() string
}
