package model

import "fmt"

// commonActionHandlers maps actions to their handler functions for any RequestStatus.
var commonActionHandlers = map[Action]func(RequestStatus, string){
	ActionApprove: func(r RequestStatus, _ string) {
		r.SetStatus("approve")
	},
	ActionReject: func(r RequestStatus, reason string) {
		r.SetStatus("reject")
		r.SetReason(reason)
	},
}

// ApplyStatus updates the status on any RequestStatus using the common map.
func ApplyStatus(r RequestStatus, action Action, reason string) error {
	if handler, ok := commonActionHandlers[action]; ok {
		handler(r, reason)
		return nil
	}
	return fmt.Errorf("invalid action: %v", action)
}
