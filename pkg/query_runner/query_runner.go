// оборачивает обработку построения запроса от squirell

package query_runner

import (
	"context"
	"database/sql"
	"fmt"
)

type Builder interface {
	ToSql() (string, []interface{}, error)
}

func (r *Runner) Getx(ctx context.Context, dest interface{}, query Builder) error {
	queryText, args, err := query.ToSql()
	if err != nil {
		return fmt.Errorf("get query and args from builder: %w", err)
	}

	return r.GetContext(ctx, dest, queryText, args...)
}

func (r *Runner) Execx(ctx context.Context, query Builder) (sql.Result, error) {
	queryText, args, err := query.ToSql()
	if err != nil {
		return nil, fmt.Errorf("get query and args from builder: %w", err)
	}

	return r.ExecContext(ctx, queryText, args...)
}

func (r *Runner) Selectx(ctx context.Context, dest interface{}, query Builder) error {
	queryText, args, err := query.ToSql()
	if err != nil {
		return fmt.Errorf("get query and args from builder: %w", err)
	}

	return r.SelectContext(ctx, dest, queryText, args...)
}
