package util

import (
	"context"
	"log"
	"time"

	"gorm.io/gorm"
)

// TransactionManager encapsulates transaction operations.
type TransactionManager struct {
	DB *gorm.DB
}

func (tm *TransactionManager) Execute(txFunc func(tx *gorm.DB) error) error {
	// Call the context-aware version with a background context
	return tm.ExecuteContext(context.Background(), txFunc)
}

func (tm *TransactionManager) ExecuteContext(ctx context.Context, txFunc func(tx *gorm.DB) error) error {
	startTime := time.Now()

	// Use WithContext to ensure the transaction respects the passed context
	err := tm.DB.WithContext(ctx).Transaction(txFunc)

	duration := time.Since(startTime)

	// Check context error first, as it might be the primary reason for failure
	if ctxErr := ctx.Err(); ctxErr != nil {
		log.Printf("Transaction context finished after %v: %v", duration, ctxErr)
		// Return the context error, potentially masking the DB error if both occurred,
		// but the context cancellation is often the root cause to report.
		return ctxErr
	}

	if err != nil {
		log.Printf("Transaction rolled back after %v. Error: %v", duration, err)
	}
	return err
}
