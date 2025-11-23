package tx_context

import (
	"context"

	"github.com/jmoiron/sqlx"
)

type TxContext struct{}

func (TxContext) SetTxToContext(ctx context.Context, tx sqlx.Tx) context.Context {
	return context.WithValue(ctx, contextKey, tx)
}

func (TxContext) GetTxFromContext(ctx context.Context) *sqlx.Tx {
	val := ctx.Value(contextKey)
	if val == nil {
		return nil
	}

	tx := val.(sqlx.Tx)
	return &tx
}
