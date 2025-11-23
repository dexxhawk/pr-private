// оборачивает логику выбора tx или db

package query_runner

import (
	"context"
	"database/sql"
	"fmt"
)

type sqlRunner interface {
	GetContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error
	ExecContext(ctx context.Context, query string, args ...any) (sql.Result, error)
	SelectContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error
}

func (r *Runner) GetContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error {
	var runner sqlRunner = r.db

	tx := r.txGetter.GetTxFromContext(ctx)
	if tx != nil {
		runner = tx
	}

	err := runner.GetContext(ctx, dest, query, args...)
	if err != nil {
		err = pgErrEnricher(err)
		return fmt.Errorf("run sqlx getx: %w", err)
	}
	return nil
}

func (r *Runner) ExecContext(ctx context.Context, query string, args ...any) (sql.Result, error) {
	var runner sqlRunner = r.db

	tx := r.txGetter.GetTxFromContext(ctx)
	if tx != nil {
		runner = tx
	}

	result, err := runner.ExecContext(ctx, query, args...)
	if err != nil {
		err = pgErrEnricher(err)
		return nil, fmt.Errorf("run sql exec: %w", err)
	}

	return result, nil
}

func (r *Runner) SelectContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error {
	var runner sqlRunner = r.db

	tx := r.txGetter.GetTxFromContext(ctx)
	if tx != nil {
		runner = tx
	}

	err := runner.SelectContext(ctx, dest, query, args...)
	if err != nil {
		err = pgErrEnricher(err)
		return fmt.Errorf("run sqlx getx: %w", err)
	}
	return nil
}
