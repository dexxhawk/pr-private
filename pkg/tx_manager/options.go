package tx_manager

import "database/sql"

type beginTxOptions = sql.TxOptions

type BeginTxOption func(options *beginTxOptions)

func WithIsolationLevel(level sql.IsolationLevel) BeginTxOption {
	return BeginTxOption(
		func(options *beginTxOptions) {
			options.Isolation = level
		},
	)
}
