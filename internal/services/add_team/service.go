package add_team

import "github.com/dexxhawk/pr-private/internal/domain"

type Service struct {
	txManager domain.TxManager
	teamRepo  TeamRepo
	userRepo  UserRepo
}

func New(
	txManager domain.TxManager,
	teamRepo TeamRepo,
	userRepo UserRepo,
) Service {
	return Service{
		txManager: txManager,
		teamRepo:  teamRepo,
		userRepo:  userRepo,
	}
}
