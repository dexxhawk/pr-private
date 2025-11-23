package tx_manager

import (
	"context"

	"github.com/jmoiron/sqlx"
)

type txContext interface {
	SetTxToContext(ctx context.Context, tx sqlx.Tx) context.Context
	GetTxFromContext(ctx context.Context) *sqlx.Tx
}
