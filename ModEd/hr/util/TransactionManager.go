package util

import (
	"gorm.io/gorm"
)

// TransactionManager encapsulates transaction operations.
type TransactionManager struct {
	DB *gorm.DB
}

// Execute runs the provided function within a transaction.
func (tm *TransactionManager) Execute(txFunc func(tx *gorm.DB) error) error {
	return tm.DB.Transaction(txFunc)
}
