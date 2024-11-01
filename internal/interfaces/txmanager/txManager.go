package txmanager

import (
	"context"

	"github.com/jmoiron/sqlx"
)

type TxManager interface {
	Begin(ctx context.Context) (*sqlx.Tx, error)
	Commit(ctx context.Context, dbtx *sqlx.Tx) error
	Rollback(ctx context.Context, dbtx *sqlx.Tx) error
}

type txManager struct {
	db *sqlx.DB
}

type Opts struct {
	DB *sqlx.DB
}

func NewTxManager(opt *Opts) TxManager {
	return &txManager{db: opt.DB}
}

func (tx *txManager) Begin(ctx context.Context) (*sqlx.Tx, error) {
	dbTx, err := tx.db.BeginTxx(ctx, nil)
	if err != nil {
		return nil, err
	}
	return dbTx, nil
}

func (tx *txManager) Commit(ctx context.Context, dbtx *sqlx.Tx) error {
	if err := dbtx.Commit(); err != nil {
		return err
	}
	return nil
}

func (tx *txManager) Rollback(ctx context.Context, dbtx *sqlx.Tx) error {
	if err := dbtx.Rollback(); err != nil {
		return err
	}
	return nil
}
