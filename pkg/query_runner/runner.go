package query_runner

// если в ctx есть транзакция --> запрос от sqlx.Tx
// если транзакции НЕТ, запрос от sqlx.DB
//
// Также пакет оборачивает проверку ошибки
// при билде запроса от squirell

import "github.com/jmoiron/sqlx"

type Runner struct {
	db       *sqlx.DB
	txGetter txContextGetter
}

func New(
	db *sqlx.DB,
	txGetter txContextGetter,
) Runner {
	return Runner{
		db:       db,
		txGetter: txGetter,
	}
}
