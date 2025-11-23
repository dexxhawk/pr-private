package get_review

type Service struct {
	reviewerRepo ReviewerRepo
}

func New(
	reviewerRepo ReviewerRepo,
) Service {
	return Service{
		reviewerRepo: reviewerRepo,
	}
}
