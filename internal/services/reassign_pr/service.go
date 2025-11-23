package reassign_pr

type Service struct {
	prRepo PRRepo
	reviewerRepo ReviewerRepo
	userRepo UserRepo
}

func New(
	prRepo PRRepo,
	reviewerRepo ReviewerRepo,
	userRepo UserRepo,
) Service {
	return Service{
		prRepo: prRepo,
		reviewerRepo: reviewerRepo,
		userRepo:userRepo,
	}
}
