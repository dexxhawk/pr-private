package domain

import "time"

type PR struct {
	ID        string
	Name      string
	AuthorID  string
	Status    int16
	CreatedAt time.Time
	MergedAt  *time.Time
}
