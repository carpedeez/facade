package database

import (
	"database/sql"
	"fmt"

	"github.com/doug-martin/goqu/v9"
)

type Querier interface {
	DeleteDisplay(int64) error
	CreateDisplay(int64, string, string) (int64, error)
	CreateUser(string, string, string) error
}

type querierImpl struct {
	DB *sql.DB
}

func (q querierImpl) DeleteDisplay(id int64) error {
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

func (q querierImpl) CreateDisplay(userID int64, title, description string) (int64, error) {
	sql, params, err := goqu.Insert("displays").Cols(
		"user_id", "title", "descr", "photo_url",
	).Vals(
		goqu.Vals{userID, title, description, ""},
	).Returning("id").ToSQL()
	if err != nil {
		return 0, fmt.Errorf("failed to create sql query from parameters: %w", err)
	}
	var id int64
	err = q.DB.QueryRow(sql, params...).Scan(&id)
	if id == 0 {
		return 0, fmt.Errorf("failed to execute database query: %w", err)
	}
	return id, nil
}

func (q querierImpl) CreateUser(username, fname, lname string) error {
	sql, params, err := goqu.Insert("users").Cols(
		"username", "email", "first_name", "last_name", "photo_url", "social_links",
	).Vals(
		goqu.Vals{username, "", fname, lname, "", "{}"},
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
