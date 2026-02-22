package repository

import (
	"errors"

	"github.com/lib/pq"
)

// postgres unique violation code
const dbUniqueViolation = "23505"

func isUniqueConstraintError(err error) bool {
	var pqErr *pq.Error
	if errors.As(err, &pqErr) {
		return pqErr.Code == dbUniqueViolation
	}
	return false
}
