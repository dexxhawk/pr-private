package reviewer

import "github.com/dexxhawk/pr-private/internal/domain"

type Reviewer struct {
	PRID   string `db:"pr_id"`
	UserID string `db:"user_id"`
}

func (Reviewer) Model(domain domain.Reviewer) Reviewer {
	return Reviewer{
		PRID:   domain.PRID,
		UserID: domain.UserID,
	}
}

func (model Reviewer) Domain() domain.Reviewer {
	return domain.Reviewer{
		PRID:   model.PRID,
		UserID: model.UserID,
	}
}

func (Reviewer) Domains(models []Reviewer) []domain.Reviewer {
	domains := make([]domain.Reviewer, 0, len(models))
	for _, m := range models {
		domains = append(domains, m.Domain())
	}
	return domains
}
