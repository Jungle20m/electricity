package storage

import (
	"context"
	"gorm.io/gorm"
)

const TransactionContextKey = "TransactionContextKey"

type Storage struct {
	db *gorm.DB
}

func NewStorage(db *gorm.DB) *Storage {
	return &Storage{db: db}
}

func (s *Storage) GetConnection(ctx context.Context) *gorm.DB {
	tx, isTxFromCtx := s.GetTransactionFromCtx(ctx)
	if !isTxFromCtx {
		return s.db
	} else {
		return tx
	}
}

func (s *Storage) WithTx(ctx context.Context, fn func(txContext context.Context) error) error {
	tx := s.db.Begin()
	defer tx.Rollback()

	c := context.WithValue(ctx, TransactionContextKey, tx)
	err := fn(c)

	if err != nil {
		return err
	}
	tx.Commit()
	return nil
}

func (s *Storage) GetTransactionFromCtx(ctx context.Context) (*gorm.DB, bool) {
	tx, ok := ctx.Value(TransactionContextKey).(*gorm.DB)
	return tx, ok
}
