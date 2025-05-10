package model

import (
	"fmt"
	"strings"
)

type Action int

const (
	ActionApprove Action = iota
	ActionReject
)

var actionMap = map[string]Action{
	"approve": ActionApprove,
	"reject":  ActionReject,
}

func ParseAction(actionStr string) (Action, error) {
	action, ok := actionMap[strings.ToLower(actionStr)]
	if !ok {
		return -1, fmt.Errorf("invalid action: %s", actionStr)
	}
	return action, nil
}
