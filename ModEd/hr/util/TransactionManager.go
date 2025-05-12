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
	return tm.ExecuteContext(context.Background(), txFunc)
}

func (tm *TransactionManager) ExecuteContext(ctx context.Context, txFunc func(tx *gorm.DB) error) error {
	startTime := time.Now()

	err := tm.DB.WithContext(ctx).Transaction(txFunc)

	duration := time.Since(startTime)

	if ctxErr := ctx.Err(); ctxErr != nil {
		log.Printf("Transaction context finished after %v: %v", duration, ctxErr)
		return ctxErr
	}

	if err != nil {
		log.Printf("Transaction rolled back after %v. Error: %v", duration, err)
	}
	return err
}
