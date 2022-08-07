package database

import (
	"database/sql"
	"fmt"

	"github.com/doug-martin/goqu/v9"
)

type Querier interface {
	CreateDisplay(userID string, title, description string) (int64, error)
	GetDisplay(id int64) (Display, error)
	UpdateDisplay(id int64, title, description *string) (Display, error)
	DeleteDisplay(id int64) error

	CreateItem(userID string, displayID int64, externalLink string) (Item, error)
	GetItem(id int64) (Item, error)
	GetItems(displayID int64) ([]Item, error)
	UpdateItem(id int64, externalLink, socialPostLink, photoURL *string) (Item, error)
	DeleteItem(id int64) error

	IsItemOwner(userID string, id int64) (bool, error)
	IsDisplayOwner(userID string, id int64) (bool, error)
}

type querierImpl struct {
	DB *sql.DB
}

func (q querierImpl) CreateDisplay(userID string, title, description string) (int64, error) {
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

func (q querierImpl) GetDisplay(id int64) (Display, error) {
	d := Display{}
	sql, params, err := goqu.Select("*").From("displays").Where(goqu.Ex{"id": id}).ToSQL()
	if err != nil {
		return d, fmt.Errorf("failed to create sql query from parameters: %w", err)
	}
	err = q.DB.QueryRow(sql, params...).Scan(&d)
	if err != nil {
		return d, fmt.Errorf("failed to execute database query: %w", err)
	}
	return d, nil
}

func (q querierImpl) UpdateDisplay(id int64, title, description *string) (Display, error) {
	d := Display{}
	u := goqu.Update("displays")
	r := goqu.Record{}
	if title != nil {
		r["title"] = *title
	}
	if description != nil {
		r["descr"] = *description
	}
	sql, params, err := u.Set(r).Returning("*").ToSQL()
	if err != nil {
		return d, fmt.Errorf("failed to create sql query from parameters: %w", err)
	}
	err = q.DB.QueryRow(sql, params...).Scan(&d)
	if err != nil {
		return d, fmt.Errorf("failed to execute database query: %w", err)
	}
	return d, nil
}

func (q querierImpl) DeleteDisplay(id int64) error {
	sql, params, err := goqu.Delete("displays").Where(goqu.Ex{"id": id}).ToSQL()
	if err != nil {
		return fmt.Errorf("failed to create sql query from parameters: %w", err)
	}
	_, err = q.DB.Exec(sql, params...)
	if err != nil {
		return fmt.Errorf("failed to execute database query: %w", err)
	}
	return nil
}

func (q querierImpl) CreateItem(userID string, displayID int64, externalLink string) (Item, error) {
	item := Item{}
	sql, params, err := goqu.Insert("items").Cols(
		"user_id", "display_id", "external_link", "social_post_link", "photo_url",
	).Vals(
		goqu.Vals{userID, displayID, externalLink, "", ""},
	).Returning("*").ToSQL()
	if err != nil {
		return item, fmt.Errorf("failed to create sql query from parameters: %w", err)
	}
	err = q.DB.QueryRow(sql, params...).Scan(&item)
	if err != nil {
		return item, fmt.Errorf("failed to execute database query: %w", err)
	}
	return item, nil
}

func (q querierImpl) GetItem(id int64) (Item, error) {
	item := Item{}
	sql, params, err := goqu.Select("*").From("items").Where(goqu.Ex{"id": id}).ToSQL()
	if err != nil {
		return item, fmt.Errorf("failed to create sql query from parameters: %w", err)
	}
	err = q.DB.QueryRow(sql, params...).Scan(&item)
	if err != nil {
		return item, fmt.Errorf("failed to execute database query: %w", err)
	}
	return item, nil
}

func (q querierImpl) GetItems(displayID int64) ([]Item, error) {
	items := []Item{}
	sql, params, err := goqu.Select("*").From("items").Where(goqu.Ex{"display_id": displayID}).ToSQL()
	if err != nil {
		return items, fmt.Errorf("failed to create sql query from parameters: %w", err)
	}
	r, err := q.DB.Query(sql, params...)
	if err != nil {
		return items, fmt.Errorf("failed to execute database query: %w", err)
	}
	r.Scan(&items)
	if err != nil {
		return items, fmt.Errorf("failed to scan results: %w", err)
	}
	return items, nil
}

func (q querierImpl) UpdateItem(id int64, externalLink, socialPostLink, photoURL *string) (Item, error) {
	i := Item{}
	u := goqu.Update("items")
	r := goqu.Record{}
	if externalLink != nil {
		r["external_link"] = *externalLink
	}
	if socialPostLink != nil {
		r["social_post_link"] = *socialPostLink
	}
	if photoURL != nil {
		r["photo_url"] = *photoURL
	}
	sql, params, err := u.Set(r).Returning("*").ToSQL()
	if err != nil {
		return i, fmt.Errorf("failed to create sql query from parameters: %w", err)
	}
	err = q.DB.QueryRow(sql, params...).Scan(&i)
	if err != nil {
		return i, fmt.Errorf("failed to execute database query: %w", err)
	}
	return i, nil
}

func (q querierImpl) DeleteItem(id int64) error {
	sql, params, err := goqu.Delete("items").Where(goqu.Ex{"id": id}).ToSQL()
	if err != nil {
		return fmt.Errorf("failed to create sql query from parameters: %w", err)
	}
	_, err = q.DB.Exec(sql, params...)
	if err != nil {
		return fmt.Errorf("failed to execute database query: %w", err)
	}
	return nil
}

func (q querierImpl) IsItemOwner(userID string, id int64) (bool, error) {
	sql, params, err := goqu.Select("user_id").From("items").Where(goqu.Ex{"id": id}).ToSQL()
	if err != nil {
		return false, fmt.Errorf("failed to create sql query from parameters: %w", err)
	}
	var uid string
	err = q.DB.QueryRow(sql, params...).Scan(&uid)
	if err != nil {
		return false, fmt.Errorf("failed to execute database query: %w", err)
	}
	if uid == userID {
		return true, nil
	}
	return false, nil
}

func (q querierImpl) IsDisplayOwner(userID string, id int64) (bool, error) {
	sql, params, err := goqu.Select("user_id").From("displays").Where(goqu.Ex{"id": id}).ToSQL()
	if err != nil {
		return false, fmt.Errorf("failed to create sql query from parameters: %w", err)
	}
	var uid string
	err = q.DB.QueryRow(sql, params...).Scan(&uid)
	if err != nil {
		return false, fmt.Errorf("failed to execute database query: %w", err)
	}
	if uid == userID {
		return true, nil
	}
	return false, nil
}
