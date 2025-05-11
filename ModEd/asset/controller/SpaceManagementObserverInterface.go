// MEP-1013
package controller

import (
	"ModEd/core"
)

type SpaceManagementObserverInterface[T core.RecordInterface] interface {
	HandleEvent(eventType string, dataContext T)
	GetObserverID() string
}
