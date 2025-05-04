package controller

import (
	"fmt"
	"strconv"
)

// Reviewable is any request that can be approved or rejected.
type Reviewable interface {
	ApplyStatus(action, reason string) error
}

type fetcher func(id uint) (Reviewable, error)
type saver func(Reviewable) error

// ReviewRequest does the common parsing / fetching / status logic.
func ReviewRequest(
	requestID, action, reason string,
	getByID fetcher,
	save saver,
) error {
	id, err := strconv.ParseUint(requestID, 10, 32)
	if err != nil {
		return fmt.Errorf("invalid request ID: %v", err)
	}
	req, err := getByID(uint(id))
	if err != nil {
		return fmt.Errorf("failed to fetch request: %v", err)
	}
	if err := req.ApplyStatus(action, reason); err != nil {
		return err
	}
	if err := save(req); err != nil {
		return fmt.Errorf("failed to save request: %v", err)
	}
	return nil
}
