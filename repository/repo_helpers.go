package repository

import (
	"errors"

	"github.com/lib/pq"
)

// postgreSQL unique violation code
const dbUniqueViolation = "23505"

func isUniqueConstraintError(err error) bool {
	// pq package PostgreSQL error type
	var pqErr *pq.Error
	if errors.As(err, &pqErr) {
		return pqErr.Code == dbUniqueViolation
	}
	return false
}
