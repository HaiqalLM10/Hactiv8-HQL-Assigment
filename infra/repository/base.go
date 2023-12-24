package infraRepository

import (
	"errors"

	"github.com/jmoiron/sqlx"
)

var (
	ErrorRepositoryNotFound = errors.New("repository not found")
)

type repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) repository {
	return repository{
		db: db,
	}
}
