package create_pr

import "github.com/dexxhawk/pr-private/internal/domain"

type Service struct {
	txManager domain.TxManager
	teamRepo TeamRepo
	userRepo  UserRepo
	prRepo  PRRepo
	reviewerRepo ReviewerRepo
	
}

func New(
	txManager domain.TxManager,
	teamRepo TeamRepo,
	userRepo UserRepo,
	prRepo PRRepo,
	reviewerRepo ReviewerRepo,
) Service {
	return Service{
		txManager: txManager,
		teamRepo: teamRepo,
		userRepo:  userRepo,
		prRepo:  prRepo,
		reviewerRepo: reviewerRepo,
	}
}
