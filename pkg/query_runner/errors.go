package query_runner

import (
	"errors"

	"github.com/lib/pq"
)

var ErrUnique = errors.New("unique key")

func pgErrEnricher(err error) error {
	var pqError *pq.Error
	if errors.As(err, &pqError) {
		switch pqError.Code {
		case "23505":
			err = errors.Join(ErrUnique, err)
		}
	}
	return err
}
