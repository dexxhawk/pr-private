package pr

import (
	"github.com/dexxhawk/pr-private/pkg/query_runner"

	"github.com/Masterminds/squirrel"
)

type Repo struct {
	queryBuilder squirrel.StatementBuilderType
	runner       query_runner.Runner
}

func New(
	queryBuilder squirrel.StatementBuilderType,
	runner query_runner.Runner,
) Repo {
	return Repo{
		queryBuilder: queryBuilder,
		runner:       runner,
	}
}
