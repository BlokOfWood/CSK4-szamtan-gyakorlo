package data

import (
	"database/sql"
	"time"
)

type Result struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	Score     int       `json:"score"`
	CreatedAt time.Time `json:"-"`
}

type ResultModel struct {
	DB *sql.DB
}

func (m ResultModel) Insert(result *Result) error {
	stmt := `INSERT INTO results (name, score, created_at) VALUES (?, ?, ?)
	RETURNING id, created_at`

	return m.DB.QueryRow(stmt, result.Name, result.Score, time.Now()).Scan(&result.ID, &result.CreatedAt)
}

func (m ResultModel) Get(id int64) (*Result, error) {
	if id < 1 {
		return nil, ErrRecordNotFound
	}

	stmt := `
		SELECT id, name, score, created_at 
		FROM results 
		WHERE id = ?`

	row := m.DB.QueryRow(stmt, id)

	var result Result

	err := row.Scan(&result.ID, &result.Name, &result.Score, &result.CreatedAt)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (m ResultModel) GetAll() ([]*Result, error) {
	stmt := `
		SELECT id, name, score, created_at 
		FROM results 
		ORDER BY score DESC`

	rows, err := m.DB.Query(stmt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []*Result

	for rows.Next() {
		var result Result

		err := rows.Scan(&result.ID, &result.Name, &result.Score, &result.CreatedAt)
		if err != nil {
			return nil, err
		}

		results = append(results, &result)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return results, nil
}
