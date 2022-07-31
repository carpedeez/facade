package database

import (
	"database/sql"

	"github.com/doug-martin/goqu/v9"
)

type Querier struct {
	DB *sql.DB
}

func (q Querier) DeleteDisplay(id uint64) error {
	sql, params, err := goqu.Delete("displays").Where(goqu.Ex{"id": id}).ToSQL()
	if err != nil {
		return err
	}
	_, err = q.DB.Exec(sql, params...)
	if err != nil {
		return err
	}
	return nil
}

func (q Querier) CreateDisplay(username, title, description string) (int64, error) {
	sql, params, err := goqu.Insert("displays").Cols(
		"username", "title", "descr",
	).Vals(
		goqu.Vals{username, title, description},
	).ToSQL()
	if err != nil {
		return 0, err
	}
	res, err := q.DB.Exec(sql, params...)
	if err != nil {
		return 0, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}
	return id, nil
}
