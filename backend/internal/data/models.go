package data

import (
	"database/sql"
	"errors"
)

var (
	ErrRecordNotFound = errors.New("sql: no rows in result set")
)

type Models struct {
	Results ResultModel
}

func NewModels(db *sql.DB) Models {
	return Models{
		Results: ResultModel{DB: db},
	}
}
