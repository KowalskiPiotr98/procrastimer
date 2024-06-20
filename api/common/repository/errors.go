package repository

import (
	"errors"
	"github.com/lib/pq"
)

var (
	DataNotFoundErr = errors.New("requested data was not found in the database")
)

func IsDataNotFoundErr(err error) bool {
	if err.Error() == "sql: no rows in result set" {
		return true
	}
	var pgErr *pq.Error
	if errors.As(err, &pgErr) && pgErr.Code == "23503" {
		return true
	}
	return false
}
