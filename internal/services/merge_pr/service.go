package merge_pr

type Service struct {
	prRepo PRRepo
	reviewerRepo ReviewerRepo
}

func New(
	prRepo PRRepo,
	reviewerRepo ReviewerRepo,
) Service {
	return Service{
		prRepo: prRepo,
		reviewerRepo: reviewerRepo,
	}
}
