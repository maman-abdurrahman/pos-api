package utils

import (
	"context"

	"gorm.io/gorm"
)

// WithTransaction runs fn inside a DB transaction.
// It commits if fn returns nil, otherwise rollbacks.
func WithTransaction(
	ctx context.Context,
	db *gorm.DB,
	fn func(tx *gorm.DB) error,
) error {
	return db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		return fn(tx)
	})
}
