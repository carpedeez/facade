package database

import (
	"database/sql"
	"fmt"

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
		"username", "title", "descr", "photourl",
	).Vals(
		goqu.Vals{username, title, description, ""},
	).Returning("id").ToSQL()
	if err != nil {
		return 0, fmt.Errorf("failed to create sql query from parameters: %w", err)
	}
	var id int64
	err = q.DB.QueryRow(sql, params...).Scan(&id)
	fmt.Println(id)
	if id == 0 {
		return 0, fmt.Errorf("failed to execute database query: %w", err)
	}
	return id, nil
}

func (q Querier) CreateUser(username, fname, lname string) error {
	sql, params, err := goqu.Insert("users").Cols(
		"username", "fname", "lname", "photourl", "sociallinks",
	).Vals(
		goqu.Vals{username, fname, lname, "", "{}"},
	).ToSQL()
	if err != nil {
		return fmt.Errorf("failed to create sql query from parameters: %w", err)
	}
	_, err = q.DB.Exec(sql, params...)
	if err != nil {
		return fmt.Errorf("failed to execute database query: %w", err)
	}
	return nil
}
