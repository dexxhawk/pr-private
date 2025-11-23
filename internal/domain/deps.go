package domain

import (
	"context"

	"github.com/jmoiron/sqlx"

	"github.com/dexxhawk/pr-private/pkg/tx_manager"
)

type TxManager interface {
	GetTx(ctx context.Context) *sqlx.Tx
	Begin(ctx context.Context, options ...tx_manager.BeginTxOption) (context.Context, error)
	Commit(ctx context.Context) error
	Rollback(ctx context.Context) error

	Do(ctx context.Context, callback func(ctx context.Context) error, options ...tx_manager.BeginTxOption) error
}
