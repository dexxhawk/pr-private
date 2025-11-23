package tx_manager

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/jmoiron/sqlx"
)

func (m *Manager) GetTx(ctx context.Context) *sqlx.Tx {
	return m.txContext.GetTxFromContext(ctx)
}

func (m *Manager) Begin(ctx context.Context, options ...BeginTxOption) (context.Context, error) {
	opts := (*sql.TxOptions)(nil)

	if len(options) != 0 {
		opts = &sql.TxOptions{}
		for _, o := range options {
			o(opts)
		}
	}

	tx, err := m.db.BeginTxx(ctx, opts)
	if err != nil {
		return context.Background(), err
	}

	return m.txContext.SetTxToContext(ctx, *tx), nil
}

func (m *Manager) Commit(ctx context.Context) error {
	tx := m.txContext.GetTxFromContext(ctx)
	if tx == nil {
		return ErrNoTx
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("tx commit: %w", err)
	}

	return nil
}

func (m *Manager) Rollback(ctx context.Context) error {
	tx := m.txContext.GetTxFromContext(ctx)
	if tx == nil {
		return ErrNoTx
	}

	if err := tx.Rollback(); err != nil {
		return fmt.Errorf("tx rollback: %w", err)
	}
	return nil
}

func (m *Manager) Do(ctx context.Context, callback func(ctx context.Context) error, options ...BeginTxOption) error {
	hasExternalTx := m.txContext.GetTxFromContext(ctx) != nil

	var err error
	if !hasExternalTx {
		ctx, err = m.Begin(ctx, options...)
		if err != nil {
			return fmt.Errorf("begin tx: %w", err)
		}
	}

	err = callback(ctx)
	// блок ошибки при callback'е
	if err != nil {
		// при наличии внешней транзакции
		// RollBack НЕ делаем
		if hasExternalTx {
			return fmt.Errorf("do callback in tx: %w", err)
		}
		// внешней транзакции НЕТ
		// делаем rollback
		rollbackErr := m.Rollback(ctx)
		if rollbackErr != nil {
			rollbackErr = fmt.Errorf("rollback: %w", rollbackErr)
		}
		return errors.Join(err, rollbackErr)
	}
	// ошибки при callback нет

	// при наличии внешней транзакции
	// commit НЕ делаем
	if hasExternalTx {
		return nil
	}
	// внешней транзакции НЕТ
	// делаем commit
	if commitErr := m.Commit(ctx); commitErr != nil {
		rollbackErr := m.Rollback(ctx)
		if rollbackErr != nil {
			rollbackErr = fmt.Errorf("rollback: %w", rollbackErr)
		}
		return errors.Join(commitErr, rollbackErr)
	}

	return nil
}
