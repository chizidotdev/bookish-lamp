package sqlx

import (
	"context"
	"database/sql"
	"fmt"
)

// Store provides all functions to execute db queries and transactions
type Store struct {
	*Queries
	ConnDB *sql.DB
}

// NewStore creates a new Store
func NewStore(db *sql.DB) *Store {
	return &Store{
		Queries: New(db),
		ConnDB:  db,
	}
}

// ExecTx executes a function within a transaction.
func (s *Store) ExecTx(ctx context.Context, fn func(*Queries) error) error {
	tx, err := s.ConnDB.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	qtx := s.WithTx(tx)

	err = fn(qtx)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx error: %v, rb error: %v", err, rbErr)
		}
		return err
	}
	return tx.Commit()
}
