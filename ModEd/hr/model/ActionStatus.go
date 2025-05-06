package model

import "fmt"

type Action int

const (
	ActionApprove Action = iota
	ActionReject
)

func ParseAction(actionStr string) (Action, error) {
	switch actionStr {
	case "approve", "APPROVE":
		return ActionApprove, nil
	case "reject", "REJECT":
		return ActionReject, nil
	default:
		return -1, fmt.Errorf("invalid action: %s", actionStr)
	}
}
