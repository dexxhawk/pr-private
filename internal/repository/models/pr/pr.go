package pr

import (
	"database/sql"
	"time"

	"github.com/dexxhawk/pr-private/internal/domain"
)

type PR struct {
	ID        string       `db:"id"`
	Name      string       `db:"name"`
	AuthorID  string       `db:"author_id"`
	Status    int16        `db:"status"`
	CreatedAt time.Time    `db:"created_at"`
	MergedAt  sql.NullTime `db:"merged_at"`
}

func (PR) Model(domain domain.PR) PR {
	model := PR{
		ID:        domain.ID,
		Name:      domain.Name,
		AuthorID:  domain.AuthorID,
		Status:    domain.Status,
		CreatedAt: domain.CreatedAt,
		MergedAt:  sql.NullTime{},
	}

	if domain.MergedAt != nil {
		model.MergedAt.Time = *domain.MergedAt
		model.MergedAt.Valid = true
	}

	return model
}

func (model PR) Domain() domain.PR {
	domain := domain.PR{
		ID:        model.ID,
		Name:      model.Name,
		AuthorID:  model.AuthorID,
		Status:    model.Status,
		CreatedAt: model.CreatedAt,
	}

	if model.MergedAt.Valid {
		domain.MergedAt = &model.MergedAt.Time
	}

	return domain
}

func (PR) Domains(models []PR) []domain.PR {
	domains := make([]domain.PR, 0, len(models))
	for _, m := range models {
		domains = append(domains, m.Domain())
	}
	return domains
}
