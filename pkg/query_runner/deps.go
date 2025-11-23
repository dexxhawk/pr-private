package query_runner

import (
	"context"

	"github.com/jmoiron/sqlx"
)

type txContextGetter interface {
	GetTxFromContext(ctx context.Context) *sqlx.Tx
}
