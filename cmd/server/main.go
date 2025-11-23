package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
	_ "github.com/joho/godotenv/autoload"
	_ "github.com/lib/pq"

	"github.com/dexxhawk/pr-private/internal/generated/api"
	server "github.com/dexxhawk/pr-private/internal/http"
	rpr "github.com/dexxhawk/pr-private/internal/repository/pr"
	rreviewer "github.com/dexxhawk/pr-private/internal/repository/reviewer"
	rteam "github.com/dexxhawk/pr-private/internal/repository/team"
	ruser "github.com/dexxhawk/pr-private/internal/repository/user"
	"github.com/dexxhawk/pr-private/internal/services/add_team"
	"github.com/dexxhawk/pr-private/internal/services/create_pr"
	"github.com/dexxhawk/pr-private/internal/services/get_review"
	"github.com/dexxhawk/pr-private/internal/services/get_team"
	"github.com/dexxhawk/pr-private/internal/services/merge_pr"
	"github.com/dexxhawk/pr-private/internal/services/reassign_pr"
	"github.com/dexxhawk/pr-private/internal/services/set_isactive"
	"github.com/dexxhawk/pr-private/pkg/query_runner"
	"github.com/dexxhawk/pr-private/pkg/tx_context"
	"github.com/dexxhawk/pr-private/pkg/tx_manager"
)

func main() {
	db, err := sqlx.Open("postgres", os.Getenv("SERVICE_DB_URL"))
	if err != nil {
		panic(fmt.Errorf("get db conn: %w", err))
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(fmt.Errorf("ping db : %w", err))
	}

	sqlRunner := query_runner.New(db, tx_context.TxContext{})
	teamRepo := rteam.New(
		squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar),
		sqlRunner,
	)
	userRepo := ruser.New(
		squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar),
		sqlRunner,
	)
	reviewerRepo := rreviewer.New(
		squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar),
		sqlRunner,
	)
	prRepo := rpr.New(
		squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar),
		sqlRunner,
	)

	txManager := tx_manager.New(db, tx_context.TxContext{})

	addTeamSrv := add_team.New(&txManager, &teamRepo, &userRepo)
	getTeamSrv := get_team.New(&userRepo, &teamRepo)
	setIsActiveSrv := set_isactive.New(&userRepo)
	getUserReviewPRsSrv := get_review.New(&reviewerRepo)
	createPRSrv := create_pr.New(&txManager, &teamRepo, &userRepo, &prRepo, &reviewerRepo)
	mergePRSrv := merge_pr.New(&prRepo, &reviewerRepo)
	reassignPRSrv := reassign_pr.New(&prRepo, &reviewerRepo, &userRepo)

	httpHandlers := server.New(&addTeamSrv, &getTeamSrv, &setIsActiveSrv, &getUserReviewPRsSrv, &createPRSrv, &mergePRSrv, &reassignPRSrv)

	server, err := api.NewServer(&httpHandlers)
	if err != nil {
		panic(fmt.Errorf("init http server: %w", err))
	}

	http.ListenAndServe(os.Getenv("SERVICE_URL"), server)
}
