package tx_manager

import "github.com/jmoiron/sqlx"

type Manager struct {
	db        *sqlx.DB
	txContext txContext
}

func New(
	db *sqlx.DB,
	txContext txContext,
) Manager {
	return Manager{
		db:        db,
		txContext: txContext,
	}
}
