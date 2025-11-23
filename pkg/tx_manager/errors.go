package tx_manager

import "errors"

var ErrNoTx = errors.New("ctx doesnt has sqlx.Tx")
